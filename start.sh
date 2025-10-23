#!/bin/bash

# 颜色定义
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 可执行文件名
BINARY_NAME="go-gin"

# 日志文件
LOG_FILE="app.log"

# PID文件
PID_FILE="app.pid"

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}     Go-Gin 应用启动脚本${NC}"
echo -e "${GREEN}========================================${NC}"

# 检查可执行文件是否存在
if [ ! -f "./${BINARY_NAME}" ]; then
    echo -e "${RED}错误: 找不到可执行文件 '${BINARY_NAME}'${NC}"
    echo -e "${YELLOW}请先运行 'make build' 编译项目${NC}"
    exit 1
fi

# 检查是否已经在运行
if [ -f "${PID_FILE}" ]; then
    OLD_PID=$(cat ${PID_FILE})
    if ps -p ${OLD_PID} > /dev/null 2>&1; then
        echo -e "${YELLOW}检测到应用正在运行 (PID: ${OLD_PID})${NC}"
        echo -e "${YELLOW}正在停止旧进程...${NC}"
        kill ${OLD_PID}
        
        # 等待进程结束
        WAIT_TIME=0
        MAX_WAIT=10
        while ps -p ${OLD_PID} > /dev/null 2>&1; do
            sleep 1
            WAIT_TIME=$((WAIT_TIME + 1))
            if [ ${WAIT_TIME} -ge ${MAX_WAIT} ]; then
                echo -e "${RED}进程未能正常停止，强制终止...${NC}"
                kill -9 ${OLD_PID}
                sleep 1
                break
            fi
        done
        
        echo -e "${GREEN}旧进程已停止${NC}"
        rm -f ${PID_FILE}
    else
        echo -e "${YELLOW}清理旧的PID文件...${NC}"
        rm -f ${PID_FILE}
    fi
fi

# 检查端口是否被占用
PORT=8080
if lsof -Pi :${PORT} -sTCP:LISTEN -t >/dev/null 2>&1; then
    echo -e "${RED}错误: 端口 ${PORT} 已被占用${NC}"
    echo "占用端口的进程："
    lsof -i :${PORT}
    exit 1
fi

# 启动应用
echo -e "${GREEN}正在启动 ${BINARY_NAME}...${NC}"

# 判断是否后台运行
if [ "$1" == "-d" ] || [ "$1" == "--daemon" ]; then
    # 后台运行
    nohup ./${BINARY_NAME} > ${LOG_FILE} 2>&1 &
    PID=$!
    echo ${PID} > ${PID_FILE}
    echo -e "${GREEN}应用已在后台启动 (PID: ${PID})${NC}"
    echo -e "日志文件: ${LOG_FILE}"
    echo -e "查看日志: tail -f ${LOG_FILE}"
    echo -e "停止应用: kill ${PID}"
else
    # 前台运行
    ./${BINARY_NAME}
fi

