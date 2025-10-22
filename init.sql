-- 创建数据库
CREATE DATABASE IF NOT EXISTS go_gin_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 使用数据库
USE go_gin_db;

-- 注意：表结构会由 GORM 自动迁移创建
-- 以下是一些测试数据，在应用启动并自动创建表后可以手动执行

-- 插入测试用户数据
-- INSERT INTO users (name, age, email, created_at, updated_at) VALUES
-- ('张三', 25, 'zhangsan@example.com', NOW(), NOW()),
-- ('李四', 30, 'lisi@example.com', NOW(), NOW()),
-- ('王五', 28, 'wangwu@example.com', NOW(), NOW());

-- 插入测试文章数据（需要先创建用户）
-- INSERT INTO articles (title, content, user_id, created_at, updated_at) VALUES
-- ('Go语言入门', '这是一篇关于Go语言的入门教程...', 1, NOW(), NOW()),
-- ('Gin框架实战', '这是一篇关于Gin框架的实战文章...', 1, NOW(), NOW()),
-- ('MySQL数据库优化', '这是一篇关于MySQL数据库优化的文章...', 2, NOW(), NOW());

-- 查看表结构
-- SHOW TABLES;
-- DESCRIBE users;
-- DESCRIBE articles;

