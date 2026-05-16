#!/bin/sh
set -e

echo "=========================================="
echo "Starting Sub2API in Hot Reload Mode"
echo "=========================================="

# Check if HOT_RELOAD is enabled
if [ "$HOT_RELOAD" != "true" ]; then
    echo "HOT_RELOAD not enabled, starting normal server..."
    exec /app/sub2api
fi

echo "Hot reload mode enabled"
echo "Backend changes will auto-reload with air"
echo "Frontend changes will auto-reload with Vite dev server"

# Function to cleanup background processes
cleanup() {
    echo "Shutting down services..."
    if [ -n "$BACKEND_PID" ]; then
        kill $BACKEND_PID 2>/dev/null || true
    fi
    if [ -n "$FRONTEND_PID" ]; then
        kill $FRONTEND_PID 2>/dev/null || true
    fi
    exit 0
}

trap cleanup SIGTERM SIGINT

# Start backend with air hot reload
echo "Starting backend with air hot reload..."
cd /app/backend

# Create .air.toml if not exists
if [ ! -f .air.toml ]; then
    cat > .air.toml << 'EOF'
root = "."
tmp_dir = "tmp"

[build]
cmd = "go build -o ./tmp/main ./cmd/server"
bin = "tmp/main"
include_ext = ["go", "tpl", "tmpl", "html"]
exclude_ext = ["_test.go"]
include_dir = []
exclude_dir = ["tmp", "vendor", "testdata"]
exclude_file = []
exclude_regex = ["_test\\.go"]
delay = 1000
stop_on_error = true
send_interrupt = false
kill_delay = 0

[log]
time = true
main_only = false

[color]
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"

[misc]
clean_on_exit = false
EOF
fi

# Start air in background
air &
BACKEND_PID=$!

# Start frontend with Vite dev server
echo "Starting frontend with Vite dev server..."
cd /app/frontend

# Check if node_modules exists, if not install
if [ ! -d "node_modules" ]; then
    echo "Installing frontend dependencies..."
    pnpm install --ignore-scripts
fi

# Start Vite dev server in background
pnpm run dev --host 0.0.0.0 --port 5173 &
FRONTEND_PID=$!

echo "=========================================="
echo "Services started!"
echo "Backend API: http://localhost:8080"
echo "Frontend Dev: http://localhost:5173"
echo "=========================================="
echo "Press Ctrl+C to stop all services"

# Wait for processes
wait