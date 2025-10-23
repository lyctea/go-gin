#!/bin/bash

# ====================================================
# Go Gin 应用停止脚本
# ====================================================

# 配置变量
APP_NAME="go-gin"
APP_DIR="/opt/go-gin"                    # 应用目录，根据实际情况修改
PID_FILE="${APP_DIR}/${APP_NAME}.pid"    # PID 文件

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 日志函数
log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 获取进程 PID
get_pid() {
    if [ -f "$PID_FILE" ]; then
        cat "$PID_FILE"
    else
        echo ""
    fi
}

# 检查应用是否正在运行
check_running() {
    if [ -f "$PID_FILE" ]; then
        PID=$(cat "$PID_FILE")
        if ps -p $PID > /dev/null 2>&1; then
            return 0  # 正在运行
        else
            return 1  # 未运行
        fi
    fi
    return 1  # 未运行
}

# 停止应用
stop() {
    log_info "正在停止 $APP_NAME..."

    # 检查是否在运行
    if ! check_running; then
        log_warn "$APP_NAME 未运行"
        # 清理可能存在的过期 PID 文件
        if [ -f "$PID_FILE" ]; then
            rm -f "$PID_FILE"
            log_info "已清理过期的 PID 文件"
        fi
        exit 0
    fi

    # 获取 PID
    PID=$(get_pid)
    log_info "找到进程 PID: $PID"

    # 尝试优雅停止 (SIGTERM)
    log_info "发送 SIGTERM 信号..."
    kill $PID

    # 等待进程结束
    WAIT_TIME=0
    MAX_WAIT=15
    while ps -p $PID > /dev/null 2>&1; do
        sleep 1
        WAIT_TIME=$((WAIT_TIME + 1))
        if [ $WAIT_TIME -ge $MAX_WAIT ]; then
            log_warn "进程未在 ${MAX_WAIT} 秒内停止，强制终止..."
            kill -9 $PID
            sleep 1
            break
        fi
    done

    # 检查是否已停止
    if ps -p $PID > /dev/null 2>&1; then
        log_error "✗ 无法停止 $APP_NAME (PID: $PID)"
        exit 1
    else
        log_info "✓ $APP_NAME 已停止"
        rm -f "$PID_FILE"
    fi
}

# 主函数
main() {
    log_info "================================"
    log_info " $APP_NAME 停止脚本"
    log_info "================================"
    log_info ""
    
    stop
}

# 执行主函数
main

