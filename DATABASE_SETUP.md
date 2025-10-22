# 数据库设置指南

## 准备工作

### 1. 安装 MySQL
确保你的系统已经安装了 MySQL 数据库服务器。

### 2. 创建数据库
登录到 MySQL 并创建数据库：

```sql
CREATE DATABASE go_gin_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 3. 配置数据库连接
修改 `config/database.go` 文件中的数据库配置：

```go
dbConfig := DatabaseConfig{
    Host:     "localhost",
    Port:     "3306",
    User:     "root",
    Password: "your_password", // 修改为你的MySQL密码
    DBName:   "go_gin_db",
    Charset:  "utf8mb4",
}
```

## 数据库表结构

应用启动时会自动创建以下表：

### users 表
- `id` (主键，自增)
- `name` (varchar(100), 非空)
- `age` (int, 非空)
- `email` (varchar(100), 唯一索引)
- `created_at` (时间戳)
- `updated_at` (时间戳)
- `deleted_at` (时间戳, 软删除)

### articles 表
- `id` (主键，自增)
- `title` (varchar(200), 非空)
- `content` (text)
- `user_id` (外键，关联 users 表)
- `created_at` (时间戳)
- `updated_at` (时间戳)
- `deleted_at` (时间戳, 软删除)

## API 接口

### 用户接口
- `GET /users` - 获取所有用户
- `GET /users/:id` - 根据ID获取用户
- `POST /users` - 创建用户
- `PUT /users/:id` - 更新用户
- `DELETE /users/:id` - 删除用户（软删除）

### 文章接口
- `GET /articles` - 获取所有文章
- `GET /articles/:id` - 根据ID获取文章
- `POST /articles` - 创建文章
- `PUT /articles/:id` - 更新文章
- `DELETE /articles/:id` - 删除文章（软删除）

## 示例请求

### 创建用户
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三",
    "age": 25,
    "email": "zhangsan@example.com"
  }'
```

### 创建文章
```bash
curl -X POST http://localhost:8080/articles \
  -H "Content-Type: application/json" \
  -d '{
    "title": "我的第一篇文章",
    "content": "这是文章的内容",
    "user_id": 1
  }'
```

### 获取用户列表
```bash
curl http://localhost:8080/users
```

### 获取文章列表
```bash
curl http://localhost:8080/articles
```

## 技术特性

- ✅ GORM ORM 框架
- ✅ MySQL 数据库
- ✅ 自动迁移表结构
- ✅ 软删除支持
- ✅ 外键关联
- ✅ 预加载关联数据
- ✅ 连接池配置
- ✅ RESTful API 设计
- ✅ 错误处理
- ✅ 统一响应格式

## 最佳实践

1. **连接池配置**: 已在 `config/database.go` 中配置了合理的连接池参数
2. **软删除**: 所有删除操作都使用软删除，数据不会真正删除
3. **外键关联**: Article 和 User 之间建立了外键关系
4. **预加载**: 查询文章时自动预加载关联的用户信息，避免 N+1 查询问题
5. **错误处理**: 所有数据库操作都有完善的错误处理
6. **日志**: GORM 日志级别设置为 Info，方便调试

## 运行项目

```bash
# 安装依赖
go mod tidy

# 运行项目
go run main.go
```

项目启动后，会自动：
1. 连接到 MySQL 数据库
2. 执行数据库表结构迁移（自动创建表）
3. 启动 HTTP 服务器在 8080 端口

