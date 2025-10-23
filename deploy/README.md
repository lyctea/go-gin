# 部署脚本使用说明

本目录包含用于在服务器上部署和管理 Go Gin 应用的脚本。

## 📦 脚本列表

| 脚本文件 | 功能说明 |
|---------|---------|
| `start.sh` | 启动应用 |
| `stop.sh` | 停止应用 |
| `restart.sh` | 重启应用 |
| `status.sh` | 查看应用状态 |

## 🚀 快速部署

### 1. 上传文件到服务器

将编译后的文件和脚本上传到服务器：

```bash
# 在本地执行，上传到服务器
scp go-gin your-server:/opt/go-gin/
scp deploy/*.sh your-server:/opt/go-gin/
```

或者使用 rsync：

```bash
rsync -avz go-gin deploy/*.sh your-server:/opt/go-gin/
```

### 2. 在服务器上设置权限

```bash
# SSH 登录到服务器
ssh your-server

# 进入应用目录
cd /opt/go-gin

# 添加执行权限
chmod +x go-gin
chmod +x *.sh
```

### 3. 启动应用

```bash
./start.sh
```

## 📝 脚本详细说明

### start.sh - 启动脚本

**功能：**
- 检查应用是否已经运行
- 创建必要的日志目录
- 后台启动应用
- 记录进程 PID
- 执行健康检查

**使用方法：**
```bash
./start.sh
```

**脚本会输出：**
- 启动状态
- 进程 PID
- 日志文件路径
- 健康检查结果

**注意事项：**
- 确保可执行文件存在：`/opt/go-gin/go-gin`
- 如果应用已在运行，脚本会提示并退出
- 日志默认保存在：`/var/log/go-gin/`

---

### stop.sh - 停止脚本

**功能：**
- 优雅停止应用进程
- 清理 PID 文件
- 超时后强制终止进程

**使用方法：**
```bash
./stop.sh
```

**停止流程：**
1. 发送 SIGTERM 信号（优雅停止）
2. 等待最多 15 秒
3. 如果未停止，发送 SIGKILL 信号（强制终止）

---

### restart.sh - 重启脚本

**功能：**
- 先停止应用
- 等待 2 秒
- 重新启动应用

**使用方法：**
```bash
./restart.sh
```

**适用场景：**
- 更新应用后重启
- 配置修改后重启
- 应用异常需要重启

---

### status.sh - 状态检查脚本

**功能：**
- 显示应用运行状态
- 显示进程信息（PID、运行时间、内存、CPU）
- 显示网络信息（端口监听）
- 执行健康检查
- 提供管理命令提示

**使用方法：**
```bash
./status.sh
```

**输出信息：**
- 应用状态：运行中 / 未运行
- PID 和资源使用情况
- 监听端口
- 健康检查结果
- 管理命令提示

---

## ⚙️ 配置说明

所有脚本中的关键配置变量：

```bash
APP_NAME="go-gin"                        # 应用名称
APP_DIR="/opt/go-gin"                    # 应用目录
APP_BIN="${APP_DIR}/${APP_NAME}"         # 可执行文件路径
PID_FILE="${APP_DIR}/${APP_NAME}.pid"    # PID 文件路径
LOG_DIR="/var/log/${APP_NAME}"           # 日志目录
```

### 修改应用目录

如果你的应用不在 `/opt/go-gin` 目录，需要修改脚本中的 `APP_DIR` 变量：

```bash
# 编辑脚本
vim start.sh

# 修改这一行
APP_DIR="/your/custom/path"
```

对所有脚本（start.sh、stop.sh、restart.sh、status.sh）都进行相同的修改。

---

## 📋 常用命令

```bash
# 启动应用
./start.sh

# 停止应用
./stop.sh

# 重启应用
./restart.sh

# 查看状态
./status.sh

# 查看实时日志
tail -f /var/log/go-gin/app.log

# 查看错误日志
tail -f /var/log/go-gin/error.log

# 查看最近 50 行日志
tail -50 /var/log/go-gin/app.log

# 检查端口占用
netstat -tlnp | grep 8080
# 或
ss -tlnp | grep 8080

# 手动健康检查
curl http://localhost:8080/health

# 查看进程详情
ps aux | grep go-gin
```

---

## 🔧 故障排查

### 应用无法启动

1. **检查可执行文件是否存在：**
   ```bash
   ls -lh /opt/go-gin/go-gin
   ```

2. **检查执行权限：**
   ```bash
   chmod +x /opt/go-gin/go-gin
   ```

3. **查看错误日志：**
   ```bash
   tail -50 /var/log/go-gin/error.log
   ```

4. **检查端口是否被占用：**
   ```bash
   netstat -tlnp | grep 8080
   ```

5. **检查数据库连接：**
   确保 MySQL 数据库可访问，配置正确。

### 应用频繁崩溃

1. **查看日志找出原因：**
   ```bash
   tail -100 /var/log/go-gin/error.log
   ```

2. **检查系统资源：**
   ```bash
   free -h          # 内存
   df -h            # 磁盘
   top              # CPU
   ```

3. **查看进程状态：**
   ```bash
   ./status.sh
   ```

### 端口被占用

```bash
# 查找占用 8080 端口的进程
netstat -tlnp | grep 8080

# 或
lsof -i :8080

# 终止占用进程
kill <PID>
```

---

## 🔐 安全建议

1. **不要使用 root 用户运行应用**
   ```bash
   # 创建专用用户
   sudo useradd -r -s /bin/false go-gin
   sudo chown -R go-gin:go-gin /opt/go-gin
   sudo chown -R go-gin:go-gin /var/log/go-gin
   
   # 使用专用用户运行
   sudo -u go-gin ./start.sh
   ```

2. **限制文件权限**
   ```bash
   chmod 750 /opt/go-gin
   chmod 750 /opt/go-gin/*.sh
   chmod 755 /opt/go-gin/go-gin
   ```

3. **定期检查日志**
   ```bash
   # 查看最近的错误
   grep -i error /var/log/go-gin/*.log
   ```

---

## 📊 日志管理

### 日志位置

- 应用日志：`/var/log/go-gin/app.log`
- 错误日志：`/var/log/go-gin/error.log`

### 日志轮转

创建日志轮转配置，避免日志文件过大：

```bash
sudo vim /etc/logrotate.d/go-gin
```

添加以下内容：

```
/var/log/go-gin/*.log {
    daily
    rotate 14
    compress
    delaycompress
    missingok
    notifempty
    create 0640 go-gin go-gin
    sharedscripts
    postrotate
        /opt/go-gin/restart.sh > /dev/null 2>&1 || true
    endscript
}
```

---

## 🔄 更新应用

当需要更新应用时：

```bash
# 1. 在本地重新编译
make build

# 2. 上传新版本到服务器
scp go-gin your-server:/tmp/

# 3. 在服务器上替换并重启
ssh your-server
cd /opt/go-gin
./stop.sh
cp /tmp/go-gin /opt/go-gin/go-gin
chmod +x go-gin
./start.sh

# 4. 验证
./status.sh
curl http://localhost:8080/health
```

---

## 🌐 配合 Nginx 使用

如果使用 Nginx 作为反向代理，建议配置：

```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }

    location /health {
        proxy_pass http://localhost:8080/health;
        access_log off;
    }
}
```

---

## 📞 技术支持

如有问题，请：

1. 查看应用日志
2. 查看错误日志
3. 检查网络和端口
4. 验证数据库连接
5. 参考主项目 README 文档

---

## 📜 许可证

本项目脚本遵循 MIT 许可证。

