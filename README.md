# Sentinel Opinion Monitor

舆情监控系统 - 基于 Golang + Gin 构建的舆情监控平台

## 🎯 项目简介

Sentinel Opinion Monitor 是一个用于监控、分析和管理的舆情监控系统，支持舆情数据的抓取、存储、查询和分析。

## 🏗 技术栈

- **Golang** >= 1.20
- **Gin** - HTTP Web 框架
- **MySQL** - 关系型数据库
- **Redis** - 缓存数据库
- **GORM v2** - ORM 框架
- **Viper** - 配置管理
- **Zap** - 高性能日志库

## 📁 项目结构

```
sentinel-opinion-monitor/
│── cmd/
│    ├── web/           # Web 服务入口
│    │    └── main.go
│    └── job/           # 定时任务/脚本运行入口
│         └── main.go
│
│── internal/
│    ├── config/        # Viper 配置对象 + 加载逻辑
│    ├── server/        # Gin HTTP Server
│    ├── router/        # 路由
│    ├── handler/       # HTTP Handler
│    ├── service/       # 业务逻辑层
│    ├── repository/    # MySQL/Redis 数据访问层
│    ├── model/         # 数据库模型
│    ├── job/           # 脚本/定时任务逻辑
│    └── pkg/
│         ├── logger/   # Zap 日志
│         ├── mysql/    # MySQL 连接管理
│         └── redis/    # Redis 连接管理
│
│── config/
│    └── config.yaml    # 配置文件
│
│── docker/
│    └── mysql/
│         └── init.sql  # MySQL 初始化脚本
│
│── docker-compose.yml  # Docker Compose 配置
│── go.mod
│── README.md
```

## 🚀 快速开始

### 前置要求

- Go >= 1.20
- Docker & Docker Compose（推荐）或 MySQL >= 5.7 和 Redis >= 6.0

### 使用 Docker 启动 MySQL 和 Redis（推荐）

项目已包含 `docker-compose.yml` 文件，可以一键启动 MySQL 和 Redis：

```bash
# 启动 MySQL 和 Redis 服务
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f

# 停止服务
docker-compose down

# 停止并删除数据卷（注意：会删除所有数据）
docker-compose down -v
```

启动后，MySQL 和 Redis 会自动配置好，数据库 `opinion_db` 会自动创建，并且会自动执行初始化 SQL 脚本创建表结构。

### 手动安装数据库

如果不使用 Docker，需要手动安装：

1. 创建 MySQL 数据库：

```sql
CREATE DATABASE opinion_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

2. 修改 `config/config.yaml` 中的数据库配置：

```yaml
mysql:
  host: 127.0.0.1
  port: 3306
  user: root
  password: your_password
  database: opinion_db
```

3. 确保 Redis 服务正在运行

### 安装依赖

```bash
go mod download
```

### 运行 Web 服务

```bash
go run cmd/web/main.go
```

服务将在 `http://localhost:8080` 启动

### 运行任务脚本

```bash
go run cmd/job/main.go --task=scan
```

## 📌 API 接口

### 健康检查

```
GET /ping
```

响应：
```json
{
  "message": "pong",
  "status": "ok"
}
```

### 获取舆情详情

```
GET /opinion/:id
GET /api/v1/opinions/:id
```

### 获取所有舆情

```
GET /api/v1/opinions
```

### 创建舆情

```
POST /api/v1/opinions
Content-Type: application/json

{
  "content": "舆情内容",
  "source": "来源"
}
```

## ⚙️ 配置说明

配置文件位于 `config/config.yaml`：

```yaml
server:
  port: 8080          # 服务端口

mysql:
  host: 127.0.0.1     # MySQL 主机
  port: 3306          # MySQL 端口
  user: root          # MySQL 用户名
  password: root      # MySQL 密码
  database: opinion_db # 数据库名
  max_idle_conn: 10   # 最大空闲连接数
  max_open_conn: 30   # 最大打开连接数

redis:
  addr: 127.0.0.1:6379 # Redis 地址
  password: ""         # Redis 密码（空字符串表示无密码）
  db: 0               # Redis 数据库编号

log:
  level: info         # 日志级别 (debug/info/warn/error)
```

## 🧪 开发指南

### 添加新的 API 接口

1. 在 `internal/model/` 中定义数据模型
2. 在 `internal/repository/` 中实现数据访问层
3. 在 `internal/service/` 中实现业务逻辑
4. 在 `internal/handler/` 中实现 HTTP 处理器
5. 在 `internal/router/router.go` 中注册路由

### 添加新的任务

1. 在 `internal/job/` 中实现任务逻辑
2. 在 `cmd/job/main.go` 中添加任务调度

## 📝 数据库模型

### Opinion（舆情）

| 字段 | 类型 | 说明 |
|------|------|------|
| id | bigint | 主键，自增 |
| content | text | 舆情内容 |
| source | varchar(255) | 来源 |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

## 🔧 构建

```bash
# 构建 Web 服务
go build -o bin/web cmd/web/main.go

# 构建任务脚本
go build -o bin/job cmd/job/main.go
```

## 📄 License

MIT

