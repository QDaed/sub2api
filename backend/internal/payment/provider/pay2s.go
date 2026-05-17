package provider

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Wei-Shaw/sub2api/internal/payment"
)

// Pay2S API endpoints
const (
	pay2sProductionAPI = "https://payment.pay2s.vn"
	pay2sSandboxAPI    = "https://sandbox-payment.pay2s.vn"
	pay2sCreateOrder   = "/v1/gateway/api/create"
)

// Pay2S response codes
const (
	pay2sResultCodeSuccess             = 0
	pay2sResultCodeAuthSuccess         = 9000
	pay2sResultCodeCancelled           = 2
	pay2sResultCodeInvalidSign         = 1002
	pay2sResultCodeTransactionNotFound = 1003
)

// Pay2S implements payment.Provider for Pay2S Vietnamese payment gateway
type Pay2S struct {
	instanceID string
	config     map[string]string
	httpClient *http.Client
}

// Pay2S configuration keys
const (
	pay2sPartnerCode  = "partnerCode"
	pay2sAccessKey    = "accessKey"
	pay2sSecretKey    = "secretKey"
	pay2sEnvironment  = "environment" // "production" or "sandbox"
	pay2sBankAccounts = "bankAccounts"
	pay2sBankID       = "bankId"
	pay2sPartnerName = "partnerName"
)

// Pay2S bank account entry
type pay2sBankAccount struct {
	AccountNumber string `json:"account_number"`
	BankID        string `json:"bank_id"`
}

// Pay2S API request/response structures
type pay2sCreateOrderRequest struct {
	AccessKey    string              `json:"accessKey"`
	PartnerCode  string              `json:"partnerCode"`
	PartnerName  string              `json:"partnerName,omitempty"`
	RequestID    string              `json:"requestId"`
	Amount       int64               `json:"amount"`
	OrderID      string              `json:"orderId"`
	OrderInfo    string              `json:"orderInfo"`
	OrderType    string              `json:"orderType,omitempty"`
	BankAccounts []pay2sBankAccount  `json:"bankAccounts,omitempty"`
	RedirectURL  string              `json:"redirectUrl"`
	IpnURL       string              `json:"ipnUrl"`
	RequestType  string              `json:"requestType"`
	Signature    string              `json:"signature"`
}

type pay2sCreateOrderResponse struct {
	PartnerCode string                `json:"partnerCode"`
	RequestID   string                `json:"requestId"`
	OrderID     string                `json:"orderId"`
	Amount      int64                 `json:"amount"`
	Lang        string                `json:"lang,omitempty"`
	Message     string                `json:"message"`
	ResultCode  int                   `json:"resultCode"`
	QRList      []pay2sQRItem         `json:"qrList,omitempty"`
	PayURL      string                `json:"payUrl,omitempty"`
}

type pay2sQRItem struct {
	BankID         string `json:"bank_id"`
	AccountNumber  string `json:"account_number"`
	AccountName    string `json:"account_name"`
	QRCode         string `json:"qrCode"`
	QRURL          string `json:"qrUrl"`
}

type pay2sIPNNotification struct {
	PartnerCode  string `json:"partnerCode"`
	OrderID      string `json:"orderId"`
	RequestID    string `json:"requestId"`
	Amount       int64  `json:"amount"`
	OrderInfo    string `json:"orderInfo"`
	OrderType    string `json:"orderType"`
	TransID      int64  `json:"transId"`
	ResultCode   int    `json:"resultCode"`
	Message      string `json:"message"`
	PayType      string `json:"payType"`
	ResponseTime string `json:"responseTime"`
	ExtraData    string `json:"extraData"`
	M2Signature  string `json:"m2signature"`
}

// NewPay2S creates a new Pay2S provider instance
func NewPay2S(instanceID string, config map[string]string) (*Pay2S, error) {
	required := []string{pay2sPartnerCode, pay2sAccessKey, pay2sSecretKey}
	for _, k := range required {
		if config[k] == "" {
			return nil, fmt.Errorf("pay2s config missing required key: %s", k)
		}
	}

	// Default to production if not specified
	if config[pay2sEnvironment] == "" {
		config[pay2sEnvironment] = "production"
	}

	return &Pay2S{
		instanceID: instanceID,
		config:     config,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}, nil
}

func (p *Pay2S) Name() string        { return "Pay2S" }
func (p *Pay2S) ProviderKey() string { return payment.TypePay2S }
func (p *Pay2S) SupportedTypes() []payment.PaymentType {
	return []payment.PaymentType{payment.TypePay2S}
}

func (p *Pay2S) MerchantIdentityMetadata() map[string]string {
	if p == nil {
		return nil
	}
	partnerCode := strings.TrimSpace(p.config[pay2sPartnerCode])
	if partnerCode == "" {
		return nil
	}
	return map[string]string{
		"partnerCode": partnerCode,
	}
}

func (p *Pay2S) getAPIBaseURL() string {
	if p.config[pay2sEnvironment] == "sandbox" {
		return pay2sSandboxAPI
	}
	return pay2sProductionAPI
}

// generateSignature creates HMAC-SHA256 signature for Pay2S API.
// Empty values are included as "key=" so the hash matches Pay2S's fixed formula.
func (p *Pay2S) generateSignature(data map[string]string) string {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var parts []string
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%s", k, data[k]))
	}
	rawHash := strings.Join(parts, "&")

	h := hmac.New(sha256.New, []byte(p.config[pay2sSecretKey]))
	h.Write([]byte(rawHash))
	return hex.EncodeToString(h.Sum(nil))
}

// verifySignature verifies the signature from Pay2S
func (p *Pay2S) verifySignature(data map[string]string, signature string) bool {
	expectedSignature := p.generateSignature(data)
	return expectedSignature == signature
}

// parseBankAccounts parses the bankAccounts config into Pay2S bank account entries.
// Supports "account_number|bank_id" format or separate bankId config.
func (p *Pay2S) parseBankAccounts() []pay2sBankAccount {
	accountsRaw := strings.TrimSpace(p.config[pay2sBankAccounts])
	if accountsRaw == "" {
		return nil
	}

	// Check if using pipe-delimited format: "99999999|ACB"
	if idx := strings.Index(accountsRaw, "|"); idx > 0 {
		return []pay2sBankAccount{
			{
				AccountNumber: strings.TrimSpace(accountsRaw[:idx]),
				BankID:        strings.TrimSpace(accountsRaw[idx+1:]),
			},
		}
	}

	// Fall back to separate bankId config
	bankID := strings.TrimSpace(p.config[pay2sBankID])
	if bankID != "" {
		return []pay2sBankAccount{
			{
				AccountNumber: accountsRaw,
				BankID:        bankID,
			},
		}
	}

	return nil
}

// CreatePayment initiates a Pay2S payment.
// The amount in req.Amount is in USD; it is converted to VND at 27,000 VND = 1 USD
// before being sent to the Pay2S gateway.
func (p *Pay2S) CreatePayment(ctx context.Context, req payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	amount, err := strconv.ParseFloat(req.Amount, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid amount: %w", err)
	}
	// Convert USD → VND (27,000 VND per 1 USD)
	vndAmount := int64(amount * 27000)
	if vndAmount <= 0 {
		return nil, fmt.Errorf("invalid amount after VND conversion: %f USD -> %d VND", amount, vndAmount)
	}

	// Pay2S orderInfo: AISHOPACC + 6 random digits
	orderInfo := "AISHOPACC" + generateRandomString(6)

	requestType := payment.TypePay2S
	bankAccounts := p.parseBankAccounts()

	// Derive notify URL from return URL origin when not provided.
	notifyURL := req.NotifyURL
	if notifyURL == "" && req.ReturnURL != "" {
		if parsed, err := url.Parse(req.ReturnURL); err == nil && parsed.Scheme != "" && parsed.Host != "" {
			notifyURL = parsed.Scheme + "://" + parsed.Host + "/api/v1/payment/webhook/pay2s"
		}
	}

	// Build signature data per Pay2S docs:
	// accessKey=$accessKey&amount=$amount&bankAccounts=Array&ipnUrl=$ipnUrl&orderId=$orderId&orderInfo=$orderInfo&partnerCode=$partnerCode&redirectUrl=$redirectUrl&requestId=$requestId&requestType=$requestType
	sigData := map[string]string{
		"accessKey":    p.config[pay2sAccessKey],
		"amount":       strconv.FormatInt(vndAmount, 10),
		"bankAccounts": "Array",
		"ipnUrl":       notifyURL,
		"orderId":      req.OrderID,
		"orderInfo":    orderInfo,
		"partnerCode":  p.config[pay2sPartnerCode],
		"redirectUrl":  req.ReturnURL,
		"requestId":    req.OrderID,
		"requestType":  requestType,
	}
	signature := p.generateSignature(sigData)

	apiRequest := pay2sCreateOrderRequest{
		AccessKey:    p.config[pay2sAccessKey],
		PartnerCode:  p.config[pay2sPartnerCode],
		RequestID:    req.OrderID,
		Amount:       vndAmount,
		OrderID:      req.OrderID,
		OrderInfo:    orderInfo,
		BankAccounts: bankAccounts,
		RedirectURL:  req.ReturnURL,
		IpnURL:       notifyURL,
		RequestType:  requestType,
		Signature:    signature,
	}

	jsonData, err := json.Marshal(apiRequest)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	apiURL := p.getAPIBaseURL() + pay2sCreateOrder
	httpReq, err := http.NewRequestWithContext(ctx, "POST", apiURL, strings.NewReader(string(jsonData)))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json; charset=UTF-8")

	resp, err := p.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("pay2s api error: status %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse response
	var apiResponse pay2sCreateOrderResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	// Handle error responses that don't match the success format
	if apiResponse.PartnerCode == "" {
		var genericResp struct {
			Status  bool   `json:"status"`
			Message string `json:"message"`
		}
		if err := json.Unmarshal(body, &genericResp); err == nil && !genericResp.Status {
			return nil, fmt.Errorf("pay2s rejected request: %s", genericResp.Message)
		}
		return nil, fmt.Errorf("pay2s unexpected response: %s", string(body))
	}

	// Check result code
	if apiResponse.ResultCode != pay2sResultCodeSuccess && apiResponse.ResultCode != pay2sResultCodeAuthSuccess {
		return nil, fmt.Errorf("pay2s error: %s (code: %d)", apiResponse.Message, apiResponse.ResultCode)
	}

	// Extract QR code from first item if available
	var qrCode string
	if len(apiResponse.QRList) > 0 {
		qrCode = apiResponse.QRList[0].QRCode
	}

	return &payment.CreatePaymentResponse{
		PayURL:     apiResponse.PayURL,
		QRCode:     qrCode,
		Currency:   "USD",
		ResultType: payment.CreatePaymentResultOrderCreated,
	}, nil
}

// QueryOrder queries the payment status from Pay2S
func (p *Pay2S) QueryOrder(ctx context.Context, tradeNo string) (*payment.QueryOrderResponse, error) {
	return &payment.QueryOrderResponse{
		TradeNo: tradeNo,
		Status:  payment.ProviderStatusPending,
		Metadata: map[string]string{
			"note": "Pay2S query API not implemented - rely on IPN notifications",
		},
	}, nil
}

// VerifyNotification parses and verifies Pay2S IPN webhook or return-redirect params.
// Pay2S uses TWO different signature formulas:
//   IPN:    accessKey&amount&extraData&message&orderId&orderInfo&orderType&partnerCode&payType&requestId&responseTime&resultCode&transId
//   Return: accessKey&amount&message&orderId&orderInfo&orderType&partnerCode&payType&requestId&responseTime&resultCode
// We try both and accept whichever matches.
func (p *Pay2S) VerifyNotification(ctx context.Context, rawBody string, headers map[string]string) (*payment.PaymentNotification, error) {
	var notification pay2sIPNNotification
	if err := json.Unmarshal([]byte(rawBody), &notification); err != nil {
		return nil, fmt.Errorf("unmarshal ipn: %w", err)
	}

	// Verify partner code
	if notification.PartnerCode != p.config[pay2sPartnerCode] {
		return nil, fmt.Errorf("invalid partner code")
	}

	// Try IPN signature formula (includes extraData + transId)
	extraData := notification.ExtraData
	if extraData == "" {
		extraData = ""
	}
	rawHashIPN := fmt.Sprintf("accessKey=%s&amount=%d&extraData=%s&message=%s&orderId=%s&orderInfo=%s&orderType=%s&partnerCode=%s&payType=%s&requestId=%s&responseTime=%s&resultCode=%d&transId=%d",
		p.config[pay2sAccessKey],
		notification.Amount,
		extraData,
		notification.Message,
		notification.OrderID,
		notification.OrderInfo,
		notification.OrderType,
		notification.PartnerCode,
		notification.PayType,
		notification.RequestID,
		notification.ResponseTime,
		notification.ResultCode,
		notification.TransID,
	)

	h := hmac.New(sha256.New, []byte(p.config[pay2sSecretKey]))
	h.Write([]byte(rawHashIPN))
	expectedSigIPN := hex.EncodeToString(h.Sum(nil))

	// Try return-redirect signature formula (no extraData, no transId)
	rawHashReturn := fmt.Sprintf("accessKey=%s&amount=%d&message=%s&orderId=%s&orderInfo=%s&orderType=%s&partnerCode=%s&payType=%s&requestId=%s&responseTime=%s&resultCode=%d",
		p.config[pay2sAccessKey],
		notification.Amount,
		notification.Message,
		notification.OrderID,
		notification.OrderInfo,
		notification.OrderType,
		notification.PartnerCode,
		notification.PayType,
		notification.RequestID,
		notification.ResponseTime,
		notification.ResultCode,
	)

	h2 := hmac.New(sha256.New, []byte(p.config[pay2sSecretKey]))
	h2.Write([]byte(rawHashReturn))
	expectedSigReturn := hex.EncodeToString(h2.Sum(nil))

	if expectedSigIPN != notification.M2Signature && expectedSigReturn != notification.M2Signature {
		return nil, fmt.Errorf("invalid pay2s signature")
	}

	// Map Pay2S result codes to Sub2API status
	var status string
	switch notification.ResultCode {
	case pay2sResultCodeSuccess, pay2sResultCodeAuthSuccess:
		status = payment.NotificationStatusSuccess
	default:
		status = payment.NotificationStatusPaid
	}

	// Convert VND amount back to USD (27,000 VND = 1 USD)
	amount := float64(notification.Amount) / 27000.0
	tradeNo := strconv.FormatInt(notification.TransID, 10)
	if tradeNo == "0" || tradeNo == "" {
		// For return-redirect notifications, transId may not be present.
		// Use requestId as a fallback trade identifier.
		tradeNo = notification.RequestID
	}

	return &payment.PaymentNotification{
		TradeNo:  tradeNo,
		OrderID:  notification.OrderID,
		Amount:   amount,
		Status:   status,
		RawData:  rawBody,
		Metadata: map[string]string{
			"resultCode": strconv.Itoa(notification.ResultCode),
			"message":    notification.Message,
			"payType":    notification.PayType,
		},
	}, nil
}

// Refund processes a refund request (placeholder - Pay2S refund API not documented)
func (p *Pay2S) Refund(ctx context.Context, req payment.RefundRequest) (*payment.RefundResponse, error) {
	return nil, fmt.Errorf("pay2s refund not implemented - API not documented")
}

// Cancel cancels a payment (placeholder - Pay2S cancel API not documented)
func (p *Pay2S) Cancel(ctx context.Context, tradeNo string) error {
	return fmt.Errorf("pay2s cancel not implemented - API not documented")
}

func generateRandomString(n int) string {
	const charset = "0123456789"
	b := make([]byte, n)
	for i := range b {
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[idx.Int64()]
	}
	return string(b)
}
