# Sub2API Hot Reload Development Setup

## 🚀 Tính năng
- **Backend Hot Reload**: Code Go changes tự động reload mà không cần rebuild
- **Frontend Hot Reload**: Vue.js changes tự động reload với Vite dev server
- **Volume Mounts**: Code changes được sync ngay lập tức vào container
- **Fast Development**: Không cần build Docker image mỗi khi code thay đổi

## 📋 Cách sử dụng

### 1. Stop containers hiện tại (nếu đang chạy)
```bash
cd deploy
docker-compose down
```

### 2. Start với Hot Reload
```bash
cd deploy
docker compose -f docker-compose.hot-reload.yml up --build
```

### 3. Truy cập ứng dụng
- **Backend API**: http://localhost:8080
- **Frontend Dev**: http://localhost:5173

## 🔧 Cách hoạt động

### Backend (Go)
- Sử dụng **air** tool để watch file changes
- Khi file `.go` thay đổi → auto rebuild & restart server
- Config file: `.air.toml` (tự động tạo)

### Frontend (Vue.js)
- Sử dụng **Vite dev server** với HMR (Hot Module Replacement)
- Khi file `.vue`, `.ts`, `.css` thay đổi → auto reload browser
- Port: 5173

## 📝 Workflow Development

1. **Code changes** trên local machine
2. **Tự động sync** vào container qua volume mounts
3. **Auto reload**:
   - Backend: air detect changes → rebuild → restart (takes ~2-5s)
   - Frontend: Vite HMR → instant browser update (takes ~100ms)

## 🛑 Stop Services

```bash
cd deploy
docker compose -f docker-compose.hot-reload.yml down
```

## 🔄 So sánh với Production Mode

| Feature | Hot Reload Dev | Production |
|---------|----------------|------------|
| Build time | Lần đầu ~5-10 phút | Mỗi lần ~5-10 phút |
| Code changes | Instant (2-5s) | Cần rebuild |
| Performance | Chậm hơn (debug mode) | Tối ưu (release mode) |
| Use case | Development | Production |

## ⚠️ Lưu ý

1. **Lần đầu chạy** sẽ lâu hơn vì cần:
   - Build Docker image dev
   - Install dependencies (node_modules, go modules)

2. **Performance**: Debug mode chậm hơn release mode
   - Chỉ dùng cho development
   - Production dùng `docker-compose.yml`

3. **Port conflicts**: 
   - Frontend dev server chạy port 5173
   - Đảm bảo port này không bị占用

4. **File permissions**:
   - Một số file có thể được tạo bởi container với permissions khác
   - Nếu gặp lỗi permissions, chạy: `sudo chown -R $USER:$USER .`

## 🐛 Troubleshooting

### Backend không reload
```bash
# Check air process
docker exec sub2api-hot-reload ps aux | grep air

# Restart backend manually
docker exec sub2api-hot-reload pkill -f air
docker exec sub2api-hot-reload air &
```

### Frontend không reload
```bash
# Restart Vite dev server
docker exec sub2api-hot-reload pkill -f vite
docker exec sub2api-hot-reload sh -c "cd /app/frontend && pnpm run dev --host 0.0.0.0 --port 5173" &
```

### Volume mount không work
```bash
# Check volume mounts
docker inspect sub2api-hot-reload | grep -A 10 Mounts

# Rebuild container
docker compose -f docker-compose.hot-reload.yml up --build --force-recreate
```

## 📚 Tài liệu tham khảo

- [Air - Go live reload](https://github.com/cosmtrek/air)
- [Vite - Hot Module Replacement](https://vitejs.dev/guide/features.html#hot-module-replacement)
- [Docker Compose Override](https://docs.docker.com/compose/extends/)