# Go Gin 框架项目

这是一个使用Go Gin框架和MySQL数据库构建的RESTful Web API项目。

## 项目结构

```
go-gin/
├── config/                # 配置文件
│   ├── database.go       # 数据库配置和连接
│   └── migrate.go        # 数据库迁移
├── handlers/             # 请求处理器
│   ├── article_handler.go
│   ├── health_handler.go
│   └── user_handler.go
├── models/               # 数据模型
│   ├── article.go
│   ├── response.go
│   └── user.go
├── routes/               # 路由配置
│   └── router.go
├── go.mod               # Go模块文件
├── main.go              # 主程序文件
├── init.sql             # 数据库初始化脚本
├── DATABASE_SETUP.md    # 数据库设置详细指南
└── README.md            # 项目说明文档
```

## 功能特性

- ✅ RESTful API 设计
- ✅ MySQL 数据库集成（使用 GORM）
- ✅ 用户管理 CRUD 接口
- ✅ 文章管理 CRUD 接口
- ✅ 外键关联（文章-用户）
- ✅ 软删除支持
- ✅ 自动数据库迁移
- ✅ 连接池配置
- ✅ 统一的响应格式
- ✅ 日志记录和错误恢复中间件
- ✅ 健康检查接口

## API接口

### 健康检查
- **GET** `/health` - 返回服务状态信息

### 用户管理
- **GET** `/users` - 获取所有用户
- **GET** `/users/:id` - 根据ID获取用户
- **POST** `/users` - 创建新用户
- **PUT** `/users/:id` - 更新用户信息
- **DELETE** `/users/:id` - 删除用户（软删除）

### 文章管理
- **GET** `/articles` - 获取所有文章
- **GET** `/articles/:id` - 根据ID获取文章
- **POST** `/articles` - 创建新文章
- **PUT** `/articles/:id` - 更新文章信息
- **DELETE** `/articles/:id` - 删除文章（软删除）

## 快速开始

### 1. 准备 MySQL 数据库

首先，确保你已经安装了 MySQL 数据库，然后创建数据库：

```bash
mysql -u root -p
```

执行以下 SQL 命令：
```sql
CREATE DATABASE go_gin_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

或者直接使用提供的脚本：
```bash
mysql -u root -p < init.sql
```

### 2. 配置数据库连接

编辑 `config/database.go` 文件，修改数据库连接信息：

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

### 3. 安装依赖
```bash
go mod tidy
```

### 4. 运行项目
```bash
go run main.go
```

项目启动后会自动：
- 连接到 MySQL 数据库
- 自动创建所需的数据表（users、articles）
- 启动 HTTP 服务器在 8080 端口

### 5. 访问接口
服务器将在 `http://localhost:8080` 启动

## 测试接口

### 健康检查
```bash
curl http://localhost:8080/health
```

### 用户相关

**创建用户**
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三",
    "age": 25,
    "email": "zhangsan@example.com"
  }'
```

**获取用户列表**
```bash
curl http://localhost:8080/users
```

**获取单个用户**
```bash
curl http://localhost:8080/users/1
```

**更新用户**
```bash
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三（已更新）",
    "age": 26,
    "email": "zhangsan@example.com"
  }'
```

**删除用户**
```bash
curl -X DELETE http://localhost:8080/users/1
```

### 文章相关

**创建文章**
```bash
curl -X POST http://localhost:8080/articles \
  -H "Content-Type: application/json" \
  -d '{
    "title": "我的第一篇文章",
    "content": "这是文章内容...",
    "user_id": 1
  }'
```

**获取文章列表**
```bash
curl http://localhost:8080/articles
```

**获取单个文章**
```bash
curl http://localhost:8080/articles/1
```

**更新文章**
```bash
curl -X PUT http://localhost:8080/articles/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "更新后的标题",
    "content": "更新后的内容...",
    "user_id": 1
  }'
```

**删除文章**
```bash
curl -X DELETE http://localhost:8080/articles/1
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

- **Go** 1.21+
- **Gin** - Web框架
- **GORM** - ORM框架
- **MySQL** - 关系型数据库
- **JSON** - 数据格式

## 数据库设计

### Users 表
| 字段 | 类型 | 说明 |
|------|------|------|
| id | uint | 主键，自增 |
| name | varchar(100) | 用户名 |
| age | int | 年龄 |
| email | varchar(100) | 邮箱（唯一） |
| created_at | timestamp | 创建时间 |
| updated_at | timestamp | 更新时间 |
| deleted_at | timestamp | 删除时间（软删除） |

### Articles 表
| 字段 | 类型 | 说明 |
|------|------|------|
| id | uint | 主键，自增 |
| title | varchar(200) | 文章标题 |
| content | text | 文章内容 |
| user_id | uint | 用户ID（外键） |
| created_at | timestamp | 创建时间 |
| updated_at | timestamp | 更新时间 |
| deleted_at | timestamp | 删除时间（软删除） |

## 最佳实践

本项目遵循以下 Gin + GORM 最佳实践：

1. **目录结构清晰** - 按照功能模块划分目录（config、models、handlers、routes）
2. **数据库连接池** - 配置合理的连接池参数，提高性能
3. **自动迁移** - 使用 GORM AutoMigrate 自动管理表结构
4. **软删除** - 使用 `gorm.DeletedAt` 实现软删除，数据可恢复
5. **外键关联** - 正确定义模型间的关联关系
6. **预加载** - 使用 `Preload` 避免 N+1 查询问题
7. **错误处理** - 完善的错误处理和日志记录
8. **统一响应** - 使用统一的 Response 结构返回数据
9. **RESTful 设计** - 遵循 RESTful API 设计规范
10. **代码注释** - 清晰的中文注释，便于维护

## 更多信息

详细的数据库配置和使用说明，请查看 [DATABASE_SETUP.md](./DATABASE_SETUP.md)
