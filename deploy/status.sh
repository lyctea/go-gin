#!/bin/bash

# ====================================================
# Go Gin 应用状态检查脚本
# ====================================================

# 配置变量
APP_NAME="go-gin"
APP_DIR="/opt/go-gin"                    # 应用目录，根据实际情况修改
PID_FILE="${APP_DIR}/${APP_NAME}.pid"    # PID 文件
LOG_DIR="/var/log/${APP_NAME}"           # 日志目录

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
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

log_status() {
    echo -e "${BLUE}[STATUS]${NC} $1"
}

# 检查命令是否存在
check_command() {
    if ! command -v $1 &> /dev/null; then
        return 1
    fi
    return 0
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
        fi
    fi
    return 1  # 未运行
}

# 获取进程运行时间
get_uptime() {
    local pid=$1
    if [ -z "$pid" ]; then
        echo "N/A"
        return
    fi
    
    # 获取进程启动时间
    if check_command ps; then
        ps -p $pid -o etime= | tr -d ' '
    else
        echo "N/A"
    fi
}

# 获取内存使用
get_memory() {
    local pid=$1
    if [ -z "$pid" ]; then
        echo "N/A"
        return
    fi
    
    if check_command ps; then
        ps -p $pid -o rss= | awk '{printf "%.2f MB", $1/1024}'
    else
        echo "N/A"
    fi
}

# 获取 CPU 使用率
get_cpu() {
    local pid=$1
    if [ -z "$pid" ]; then
        echo "N/A"
        return
    fi
    
    if check_command ps; then
        ps -p $pid -o %cpu= | tr -d ' '
    else
        echo "N/A"
    fi
}

# 健康检查
health_check() {
    if check_command curl; then
        if curl -s http://localhost:8080/health > /dev/null 2>&1; then
            echo -e "${GREEN}✓ 健康${NC}"
        else
            echo -e "${RED}✗ 异常${NC}"
        fi
    else
        echo "N/A (curl 未安装)"
    fi
}

# 检查端口
check_port() {
    if check_command netstat; then
        netstat -tlnp 2>/dev/null | grep :8080 | awk '{print $4}'
    elif check_command ss; then
        ss -tlnp 2>/dev/null | grep :8080 | awk '{print $4}'
    else
        echo "N/A"
    fi
}

# 显示状态
show_status() {
    echo "================================"
    echo " $APP_NAME 状态信息"
    echo "================================"
    echo ""

    # 检查是否运行
    if check_running; then
        PID=$(get_pid)
        log_status "应用状态: ${GREEN}运行中${NC}"
        echo ""
        echo "基本信息:"
        echo "  PID:          $PID"
        echo "  运行时间:     $(get_uptime $PID)"
        echo "  内存使用:     $(get_memory $PID)"
        echo "  CPU 使用:     $(get_cpu $PID)%"
        echo ""
        echo "网络信息:"
        echo "  监听端口:     $(check_port)"
        echo "  健康状态:     $(health_check)"
        echo ""
        echo "文件位置:"
        echo "  PID 文件:     $PID_FILE"
        echo "  日志目录:     $LOG_DIR"
        echo ""
        echo "管理命令:"
        echo "  查看日志:     tail -f $LOG_DIR/app.log"
        echo "  查看错误:     tail -f $LOG_DIR/error.log"
        echo "  停止服务:     ./stop.sh"
        echo "  重启服务:     ./restart.sh"
        
    else
        log_status "应用状态: ${RED}未运行${NC}"
        echo ""
        echo "管理命令:"
        echo "  启动服务:     ./start.sh"
        
        # 检查日志文件是否存在
        if [ -d "$LOG_DIR" ]; then
            echo ""
            echo "最近日志:"
            if [ -f "$LOG_DIR/error.log" ]; then
                echo "  错误日志:     tail -20 $LOG_DIR/error.log"
            fi
            if [ -f "$LOG_DIR/app.log" ]; then
                echo "  应用日志:     tail -20 $LOG_DIR/app.log"
            fi
        fi
    fi
    
    echo ""
    echo "================================"
}

# 主函数
main() {
    show_status
}

# 执行主函数
main

