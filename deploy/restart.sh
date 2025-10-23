#!/bin/bash

# ====================================================
# Go Gin 应用重启脚本
# ====================================================

# 配置变量
APP_NAME="go-gin"
APP_DIR="/opt/go-gin"                    # 应用目录，根据实际情况修改
SCRIPT_DIR=$(cd "$(dirname "$0")" && pwd)

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 日志函数
log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

# 主函数
main() {
    log_info "================================"
    log_info " $APP_NAME 重启脚本"
    log_info "================================"
    log_info ""
    
    # 停止服务
    if [ -f "${SCRIPT_DIR}/stop.sh" ]; then
        bash "${SCRIPT_DIR}/stop.sh"
    else
        log_info "stop.sh 脚本不存在，跳过停止步骤"
    fi
    
    # 等待一下
    sleep 2
    
    # 启动服务
    if [ -f "${SCRIPT_DIR}/start.sh" ]; then
        bash "${SCRIPT_DIR}/start.sh"
    else
        log_info "start.sh 脚本不存在"
        exit 1
    fi
}

# 执行主函数
main

