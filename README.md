# Ticket System

学习go的第二天
简易工单管理服务（Go + Gin + GORM + MySQL）。

## 项目概述
- 后端使用 Gin 构建 HTTP API。
- 使用 GORM 连接 MySQL 并通过 `db.AutoMigrate(&model.Ticket{})` 自动迁移模型。

## 前提条件
- Go (本项目 `go.mod` 指定 `go 1.23.0`)。
- MySQL 可用（本项目 `main.go` 使用的 DSN 示例为 `root:123456@tcp(127.0.0.1:3306)/ticket`）。

建议不要在生产中硬编码密码，启动前可修改 `main.go` 中的 DSN，或改造为从环境变量读取。

## 快速启动（使用 Docker 启动 MySQL）
在 Windows PowerShell 中运行：

```powershell
docker run --name ticket-mysql -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=ticket -p 3306:3306 -d mysql:8.0
```

然后在项目根目录运行：

```powershell
cd d:\JetBrains\ticket-system
go mod tidy
go run .
```

服务默认监听端口 `:8080`。

也可以构建可执行文件：

```powershell
go build -o ticket-system.exe
.\ticket-system.exe
```

## 配置
- 默认 DSN 在 `main.go` 文件中：

- 示例： `root:123456@tcp(127.0.0.1:3306)/ticket?charset=utf8mb4&parseTime=True&loc=Local`

如果想使用自定义凭据或不同数据库名，请修改 `main.go` 的 `dsn` 字符串或改为从环境变量读取。

## API 端点
路由定义在 `internal/router/router.go`：

- `POST /tickets/` — 创建 Ticket
  - 请求 JSON: `{"title":"...","description":"..."}`
  - 返回: 新创建的 Ticket 对象（JSON）

- `GET /tickets/` — 列表所有 Ticket
  - 返回: Ticket 数组（JSON）

- `GET /tickets/:id` — 根据 ID 获取 Ticket
  - 返回: 单个 Ticket（JSON）或 404

- `PUT /tickets/:id/status` — 更新 Ticket 状态
  - 请求 JSON: `{"status":"new_status"}`
  - 返回: 操作结果消息

示例 curl 请求：

```bash
# Create
curl -X POST http://localhost:8080/tickets/ -H "Content-Type: application/json" -d '{"title":"Bug A","description":"描述"}'

# List
curl http://localhost:8080/tickets/

# Get
curl http://localhost:8080/tickets/1

# Update status
curl -X PUT http://localhost:8080/tickets/1/status -H "Content-Type: application/json" -d '{"status":"closed"}'
```

## 数据库迁移
程序启动时会调用 `db.AutoMigrate(&model.Ticket{})` 自动创建/更新表结构；无需手动迁移。


