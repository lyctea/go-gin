# Go Gin 框架项目

这是一个使用Go Gin框架构建的简单Web API项目。

## 项目结构

```
go-gin/
├── go.mod          # Go模块文件
├── main.go         # 主程序文件
└── README.md       # 项目说明文档
```

## 功能特性

- 健康检查接口
- 用户管理API（增删改查）
- 统一的响应格式
- 日志记录和错误恢复中间件

## API接口

### 1. 健康检查
- **GET** `/health`
- 返回服务状态信息

### 2. 获取用户列表
- **GET** `/users`
- 返回所有用户信息

### 3. 根据ID获取用户
- **GET** `/users/:id`
- 返回指定ID的用户信息

### 4. 创建用户
- **POST** `/users`
- 创建新用户

## 快速开始

### 1. 安装依赖
```bash
go mod tidy
```

### 2. 运行项目
```bash
go run main.go
```

### 3. 访问接口
服务器将在 `http://localhost:8080` 启动

## 测试接口

### 健康检查
```bash
curl http://localhost:8080/health
```

### 获取用户列表
```bash
curl http://localhost:8080/users
```

### 创建用户
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"新用户","age":25}'
```

## 响应格式

所有API接口都使用统一的响应格式：

```json
{
  "code": 200,
  "message": "操作成功",
  "data": {
    // 具体数据
  }
}
```

## 技术栈

- Go 1.21+
- Gin Web框架
- JSON数据格式
