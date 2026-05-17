package provider

import (
	"testing"

	"github.com/Wei-Shaw/sub2api/internal/payment"
	"github.com/stretchr/testify/assert"
)

func TestNewPay2S(t *testing.T) {
	tests := []struct {
		name        string
		instanceID  string
		config      map[string]string
		wantErr     bool
		errContains string
	}{
		{
			name:       "valid config",
			instanceID: "test-instance",
			config: map[string]string{
				"partnerCode": "TEST_PARTNER",
				"accessKey":   "test_access_key",
				"secretKey":   "test_secret_key",
			},
			wantErr: false,
		},
		{
			name:       "valid config with environment",
			instanceID: "test-instance",
			config: map[string]string{
				"partnerCode": "TEST_PARTNER",
				"accessKey":   "test_access_key",
				"secretKey":   "test_secret_key",
				"environment": "sandbox",
			},
			wantErr: false,
		},
		{
			name:       "missing partnerCode",
			instanceID: "test-instance",
			config: map[string]string{
				"accessKey": "test_access_key",
				"secretKey": "test_secret_key",
			},
			wantErr:     true,
			errContains: "partnerCode",
		},
		{
			name:       "missing accessKey",
			instanceID: "test-instance",
			config: map[string]string{
				"partnerCode": "TEST_PARTNER",
				"secretKey":   "test_secret_key",
			},
			wantErr:     true,
			errContains: "accessKey",
		},
		{
			name:       "missing secretKey",
			instanceID: "test-instance",
			config: map[string]string{
				"partnerCode": "TEST_PARTNER",
				"accessKey":   "test_access_key",
			},
			wantErr:     true,
			errContains: "secretKey",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			provider, err := NewPay2S(tt.instanceID, tt.config)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
				assert.Nil(t, provider)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, provider)
				assert.Equal(t, tt.instanceID, provider.instanceID)
			}
		})
	}
}

func TestPay2S_InterfaceMethods(t *testing.T) {
	config := map[string]string{
		"partnerCode": "TEST_PARTNER",
		"accessKey":   "test_access_key",
		"secretKey":   "test_secret_key",
	}

	provider, err := NewPay2S("test-instance", config)
	assert.NoError(t, err)
	assert.NotNil(t, provider)

	assert.Equal(t, "Pay2S", provider.Name())
	assert.Equal(t, payment.TypePay2S, provider.ProviderKey())
	assert.Equal(t, []payment.PaymentType{payment.TypePay2S}, provider.SupportedTypes())
}

func TestPay2S_MerchantIdentityMetadata(t *testing.T) {
	tests := []struct {
		name     string
		config   map[string]string
		expected map[string]string
	}{
		{
			name: "with partnerCode",
			config: map[string]string{
				"partnerCode": "TEST_PARTNER",
				"accessKey":   "test_access_key",
				"secretKey":   "test_secret_key",
			},
			expected: map[string]string{
				"partnerCode": "TEST_PARTNER",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			provider, err := NewPay2S("test-instance", tt.config)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, provider.MerchantIdentityMetadata())
		})
	}
}

func TestPay2S_GenerateSignature(t *testing.T) {
	config := map[string]string{
		"partnerCode": "TEST_PARTNER",
		"accessKey":   "test_access_key",
		"secretKey":   "test_secret_key",
	}

	provider, err := NewPay2S("test-instance", config)
	assert.NoError(t, err)

	data := map[string]string{
		"accessKey": "ABC123",
		"amount":    "100000",
		"orderId":   "INV123456",
		"partnerCode": "PAY2S",
	}

	signature := provider.generateSignature(data)
	assert.NotEmpty(t, signature)
	assert.Len(t, signature, 64) // HMAC-SHA256 produces 64-character hex string
}

func TestPay2S_VerifySignature(t *testing.T) {
	config := map[string]string{
		"partnerCode": "TEST_PARTNER",
		"accessKey":   "test_access_key",
		"secretKey":   "test_secret_key",
	}

	provider, err := NewPay2S("test-instance", config)
	assert.NoError(t, err)

	data := map[string]string{
		"accessKey":   "ABC123",
		"amount":      "100000",
		"orderId":     "INV123456",
		"partnerCode": "PAY2S",
	}

	signature := provider.generateSignature(data)
	assert.True(t, provider.verifySignature(data, signature))

	// Test with invalid signature
	assert.False(t, provider.verifySignature(data, "invalid_signature"))
}

func TestPay2S_GetAPIBaseURL(t *testing.T) {
	tests := []struct {
		name       string
		environment string
		expected    string
	}{
		{
			name:       "production",
			environment: "production",
			expected:    pay2sProductionAPI,
		},
		{
			name:       "sandbox",
			environment: "sandbox",
			expected:    pay2sSandboxAPI,
		},
		{
			name:       "default (empty)",
			environment: "",
			expected:    pay2sProductionAPI,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := map[string]string{
				"partnerCode": "TEST_PARTNER",
				"accessKey":   "test_access_key",
				"secretKey":   "test_secret_key",
				"environment": tt.environment,
			}

			provider, err := NewPay2S("test-instance", config)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, provider.getAPIBaseURL())
		})
	}
}