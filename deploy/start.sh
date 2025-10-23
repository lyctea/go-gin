#!/bin/bash

# ====================================================
# Go Gin 应用启动脚本
# ====================================================

# 配置变量
APP_NAME="go-gin"
APP_DIR="/opt/go-gin"                    # 应用目录，根据实际情况修改
APP_BIN="${APP_DIR}/${APP_NAME}"         # 可执行文件路径
PID_FILE="${APP_DIR}/${APP_NAME}.pid"    # PID 文件
LOG_DIR="/var/log/${APP_NAME}"           # 日志目录
LOG_FILE="${LOG_DIR}/app.log"            # 应用日志
ERROR_LOG="${LOG_DIR}/error.log"         # 错误日志

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

# 检查命令是否存在
check_command() {
    if ! command -v $1 &> /dev/null; then
        log_error "命令 $1 未找到，请先安装"
        return 1
    fi
    return 0
}

# 创建必要的目录
create_dirs() {
    if [ ! -d "$LOG_DIR" ]; then
        log_info "创建日志目录: $LOG_DIR"
        mkdir -p "$LOG_DIR"
    fi
}

# 检查应用是否正在运行
check_running() {
    if [ -f "$PID_FILE" ]; then
        PID=$(cat "$PID_FILE")
        if ps -p $PID > /dev/null 2>&1; then
            return 0  # 正在运行
        else
            # PID 文件存在但进程不存在，删除过期的 PID 文件
            rm -f "$PID_FILE"
            return 1  # 未运行
        fi
    fi
    return 1  # 未运行
}

# 获取进程 PID
get_pid() {
    if [ -f "$PID_FILE" ]; then
        cat "$PID_FILE"
    else
        echo ""
    fi
}

# 启动应用
start() {
    log_info "正在启动 $APP_NAME..."

    # 检查可执行文件是否存在
    if [ ! -f "$APP_BIN" ]; then
        log_error "可执行文件不存在: $APP_BIN"
        log_error "请确保已将编译后的文件上传到服务器并放置在正确的目录"
        exit 1
    fi

    # 检查可执行权限
    if [ ! -x "$APP_BIN" ]; then
        log_warn "文件没有执行权限，正在添加..."
        chmod +x "$APP_BIN"
    fi

    # 检查是否已经在运行
    if check_running; then
        PID=$(get_pid)
        log_warn "$APP_NAME 已经在运行中 (PID: $PID)"
        log_info "如需重启，请先执行 stop.sh 停止服务"
        exit 0
    fi

    # 创建必要的目录
    create_dirs

    # 设置环境变量（可选）
    # export GIN_MODE=release
    # export DB_HOST=localhost
    # export DB_PORT=3306

    # 启动应用（后台运行）
    log_info "启动命令: $APP_BIN"
    nohup "$APP_BIN" >> "$LOG_FILE" 2>> "$ERROR_LOG" &
    
    # 获取进程 PID
    PID=$!
    
    # 保存 PID 到文件
    echo $PID > "$PID_FILE"
    
    # 等待一下，检查是否启动成功
    sleep 2
    
    if ps -p $PID > /dev/null 2>&1; then
        log_info "✓ $APP_NAME 启动成功！"
        log_info "  PID: $PID"
        log_info "  日志文件: $LOG_FILE"
        log_info "  错误日志: $ERROR_LOG"
        log_info ""
        log_info "查看日志: tail -f $LOG_FILE"
        log_info "查看状态: ps -p $PID"
        
        # 等待服务启动
        log_info ""
        log_info "等待服务启动..."
        sleep 3
        
        # 健康检查
        if check_command curl; then
            log_info "执行健康检查..."
            if curl -s http://localhost:8080/health > /dev/null; then
                log_info "✓ 健康检查通过！服务已成功启动"
                log_info "访问地址: http://localhost:8080"
            else
                log_warn "健康检查失败，请查看日志: tail -f $ERROR_LOG"
            fi
        else
            log_info "curl 命令未安装，跳过健康检查"
            log_info "可以手动检查: curl http://localhost:8080/health"
        fi
    else
        log_error "✗ $APP_NAME 启动失败"
        log_error "请查看错误日志: $ERROR_LOG"
        rm -f "$PID_FILE"
        exit 1
    fi
}

# 主函数
main() {
    log_info "================================"
    log_info " $APP_NAME 启动脚本"
    log_info "================================"
    log_info ""
    
    start
}

# 执行主函数
main

