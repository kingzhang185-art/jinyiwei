你现在是一个资深 Golang 开发工程师，请基于以下要求帮我生成一个完整的 Golang 项目脚手架：

🎯 项目名称

sentinel-opinion-monitor

🏗 技术栈要求

Golang >= 1.20

Gin (HTTP API)

MySQL

Redis

GORM v2

Viper 配置管理

Zap 日志

Wire 依赖注入（可选，但推荐）

📁 项目结构要求（必须严格按照此结构生成）
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
│         ├── mysql/
│         └── redis/
│
│── config/
│    ├── config.yaml
│
│── go.mod
│── README.md

⚙ 配置文件 config.yaml 示例（必须生成）
server:
  port: 8080

mysql:
  host: 127.0.0.1
  port: 3306
  user: root
  password: root
  database: opinion_db
  max_idle_conn: 10
  max_open_conn: 30

redis:
  addr: 127.0.0.1:6379
  password: ""
  db: 0

log:
  level: info

🚀 两个入口必须包含以下逻辑
1. Web 服务入口 cmd/web/main.go

初始化配置

初始化日志

初始化 MySQL

初始化 Redis

注册路由

启动 Gin Server

优雅退出（graceful shutdown）

2. 脚本任务入口 cmd/job/main.go

用于执行例如“舆情抓取”、“关键词扫描”、“定时监控”的任务。

必须包含：

初始化配置

初始化日志

初始化 MySQL

初始化 Redis

示例任务：ScanOpinionJob()

支持 CLI 参数，例如：

go run cmd/job/main.go --task=scan

📌 必须生成的示例业务（用于校验项目可运行）
数据模型：Opinion（舆情）
id          bigint
content     text
source      varchar
created_at  datetime

生成示例接口：

GET /ping
GET /opinion/:id

生成示例任务：

ScanOpinionJob()：打印 “scanning opinion...”

🧪 保证项目可以：

go run cmd/web/main.go 能成功启动

go run cmd/job/main.go --task=scan 能运行脚本

MySQL/Redis 能成功初始化

目录结构清晰，可扩展

📝 最后要求

请生成：

完整项目结构

每个文件的代码内容

config.yaml

go.mod

必要的注释

README.md

格式必须严格可被 Cursor 识别并逐文件创建。