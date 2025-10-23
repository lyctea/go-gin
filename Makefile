.PHONY: build run test clean tidy install help dev

# 变量定义
BINARY_NAME=go-gin
GO=go
GOFLAGS=-v

# 默认目标
.DEFAULT_GOAL := help

# 编译项目
build:
	@echo "正在编译项目..."
	$(GO) build $(GOFLAGS) -o $(BINARY_NAME) main.go
	@echo "编译完成: $(BINARY_NAME)"

# 运行项目
run:
	@echo "正在运行项目..."
	$(GO) run main.go

# 开发模式（使用 air 或直接运行）
dev: run

# 运行测试
test:
	@echo "正在运行测试..."
	$(GO) test -v ./...

# 清理编译文件
clean:
	@echo "正在清理编译文件..."
	@rm -f $(BINARY_NAME)
	@rm -rf ./tmp
	@echo "清理完成"

# 整理依赖
tidy:
	@echo "正在整理依赖..."
	$(GO) mod tidy
	@echo "依赖整理完成"

# 安装依赖
install:
	@echo "正在安装依赖..."
	$(GO) mod download
	@echo "依赖安装完成"

# 格式化代码
fmt:
	@echo "正在格式化代码..."
	$(GO) fmt ./...
	@echo "格式化完成"

# 代码检查
vet:
	@echo "正在进行代码检查..."
	$(GO) vet ./...
	@echo "检查完成"

# 构建并运行
build-run: build
	@echo "正在运行编译后的程序..."
	./$(BINARY_NAME)

# 完整检查（格式化、检查、测试）
check: fmt vet test

# 显示帮助信息
help:
	@echo "可用的 make 命令："
	@echo "  make build      - 编译项目"
	@echo "  make run        - 运行项目（开发模式）"
	@echo "  make dev        - 开发模式（同 run）"
	@echo "  make test       - 运行测试"
	@echo "  make clean      - 清理编译文件"
	@echo "  make tidy       - 整理依赖"
	@echo "  make install    - 安装依赖"
	@echo "  make fmt        - 格式化代码"
	@echo "  make vet        - 代码检查"
	@echo "  make build-run  - 编译并运行"
	@echo "  make check      - 完整检查（格式化、检查、测试）"
	@echo "  make help       - 显示此帮助信息"

