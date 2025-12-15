# 渠道管理 API 文档

## 概述

渠道管理系统支持舆情监控平台的渠道管理功能。系统预设了8个主流社交媒体平台渠道：小红书、微博、快手、抖音、西瓜视频、百度贴吧、知乎、虎扑。

## API 接口

### 1. 获取渠道列表

**接口地址：** `GET /api/v1/channels`

**认证要求：** 需要登录（Bearer Token）

**查询参数：**
- `status` (可选): 状态筛选，`active` 或 `1` 表示只获取启用的渠道

**示例请求：**
```bash
# 获取所有渠道
curl -X GET http://localhost:8080/api/v1/channels \
  -H "Authorization: Bearer <token>"

# 只获取启用的渠道
curl -X GET "http://localhost:8080/api/v1/channels?status=active" \
  -H "Authorization: Bearer <token>"
```

**响应示例：**
```json
{
  "data": [
    {
      "id": 1,
      "name": "小红书",
      "code": "xiaohongshu",
      "description": "小红书平台",
      "icon": "xiaohongshu",
      "sort": 1,
      "status": 1,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    },
    {
      "id": 2,
      "name": "微博",
      "code": "weibo",
      "description": "微博平台",
      "icon": "weibo",
      "sort": 2,
      "status": 1,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

### 2. 获取渠道详情

**接口地址：** `GET /api/v1/channels/:id`

**认证要求：** 需要登录（Bearer Token）

**路径参数：**
- `id`: 渠道ID

**示例请求：**
```bash
curl -X GET http://localhost:8080/api/v1/channels/1 \
  -H "Authorization: Bearer <token>"
```

**响应示例：**
```json
{
  "data": {
    "id": 1,
    "name": "小红书",
    "code": "xiaohongshu",
    "description": "小红书平台",
    "icon": "xiaohongshu",
    "sort": 1,
    "status": 1,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 3. 创建渠道（需要 admin 角色）

**接口地址：** `POST /api/v1/channels`

**认证要求：** 需要登录且具有 admin 角色

**请求体：**
```json
{
  "name": "新渠道",
  "code": "new_channel",
  "description": "新渠道描述",
  "icon": "new_channel_icon",
  "sort": 10
}
```

**字段说明：**
- `name` (必填): 渠道名称，最大50字符
- `code` (必填): 渠道代码，最大50字符，必须唯一
- `description` (可选): 渠道描述，最大255字符
- `icon` (可选): 图标URL或图标标识，最大255字符
- `sort` (可选): 排序值，默认0

**示例请求：**
```bash
curl -X POST http://localhost:8080/api/v1/channels \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "B站",
    "code": "bilibili",
    "description": "B站平台",
    "icon": "bilibili",
    "sort": 9
  }'
```

**响应示例：**
```json
{
  "message": "创建成功",
  "data": {
    "id": 9,
    "name": "B站",
    "code": "bilibili",
    "description": "B站平台",
    "icon": "bilibili",
    "sort": 9,
    "status": 1,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 4. 更新渠道（需要 admin 角色）

**接口地址：** `PUT /api/v1/channels/:id`

**认证要求：** 需要登录且具有 admin 角色

**路径参数：**
- `id`: 渠道ID

**请求体：**
```json
{
  "name": "更新后的渠道名",
  "description": "更新后的描述",
  "icon": "updated_icon",
  "sort": 5,
  "status": 1
}
```

**字段说明：**
- `name` (可选): 渠道名称
- `description` (可选): 渠道描述
- `icon` (可选): 图标URL或图标标识
- `sort` (可选): 排序值
- `status` (可选): 状态，1-正常，2-禁用

**示例请求：**
```bash
curl -X PUT http://localhost:8080/api/v1/channels/1 \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "小红书",
    "description": "小红书社交电商平台",
    "sort": 1,
    "status": 1
  }'
```

**响应示例：**
```json
{
  "message": "更新成功"
}
```

### 5. 删除渠道（需要 admin 角色）

**接口地址：** `DELETE /api/v1/channels/:id`

**认证要求：** 需要登录且具有 admin 角色

**路径参数：**
- `id`: 渠道ID

**示例请求：**
```bash
curl -X DELETE http://localhost:8080/api/v1/channels/9 \
  -H "Authorization: Bearer <token>"
```

**响应示例：**
```json
{
  "message": "删除成功"
}
```

## 预设渠道数据

系统初始化时会自动创建以下8个渠道：

| ID | 名称 | 代码 | 排序 |
|----|------|------|------|
| 1 | 小红书 | xiaohongshu | 1 |
| 2 | 微博 | weibo | 2 |
| 3 | 快手 | kuaishou | 3 |
| 4 | 抖音 | douyin | 4 |
| 5 | 西瓜视频 | xigua | 5 |
| 6 | 百度贴吧 | tieba | 6 |
| 7 | 知乎 | zhihu | 7 |
| 8 | 虎扑 | hupu | 8 |

## 权限说明

- **查看渠道**：需要登录认证
- **创建/更新/删除渠道**：需要 admin 角色

## 错误码说明

- `400` - 请求参数错误
- `401` - 未认证或 token 无效
- `403` - 无权限访问（需要 admin 角色）
- `404` - 渠道不存在
- `500` - 服务器内部错误

## 使用示例

### 完整流程示例

```bash
# 1. 登录获取 token
TOKEN=$(curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}' \
  | jq -r '.data.token')

# 2. 获取所有启用的渠道
curl -X GET "http://localhost:8080/api/v1/channels?status=active" \
  -H "Authorization: Bearer $TOKEN"

# 3. 创建新渠道（需要 admin 角色）
curl -X POST http://localhost:8080/api/v1/channels \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "B站",
    "code": "bilibili",
    "description": "B站平台",
    "icon": "bilibili",
    "sort": 9
  }'

# 4. 更新渠道
curl -X PUT http://localhost:8080/api/v1/channels/9 \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "sort": 1,
    "status": 1
  }'

# 5. 删除渠道
curl -X DELETE http://localhost:8080/api/v1/channels/9 \
  -H "Authorization: Bearer $TOKEN"
```

## 渠道代码说明

渠道代码采用小写英文，便于程序识别和处理：

- `xiaohongshu` - 小红书
- `weibo` - 微博
- `kuaishou` - 快手
- `douyin` - 抖音
- `xigua` - 西瓜视频
- `tieba` - 百度贴吧
- `zhihu` - 知乎
- `hupu` - 虎扑

这些代码可以用于：
- 舆情数据关联
- 数据统计和筛选
- 前端展示和筛选

