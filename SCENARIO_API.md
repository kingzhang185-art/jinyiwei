# 场景和监测组管理 API 文档

## 概述

场景和监测组管理系统支持创建监测场景，并为每个场景配置多个监测组。每个监测组可以绑定多个渠道、关键词和排除词，实现精准的舆情监测。

## 数据模型关系

- **场景（Scenario）**：一个场景可以包含多个监测组
- **监测组（MonitoringGroup）**：属于某个场景，可以绑定多个渠道、关键词和排除词
- **渠道（Channel）**：监测组可以绑定多个渠道（多对多关系）
- **关键词（Keyword）**：监测组可以绑定多个关键词（一对多关系）
- **排除词（ExclusionWord）**：监测组可以绑定多个排除词（一对多关系）

## 场景管理 API

### 1. 创建场景

**接口地址：** `POST /api/v1/scenarios`

**认证要求：** 需要登录且具有 admin 角色

**请求体：**
```json
{
  "name": "企业品牌监测",
  "tag_id": 1
}
```

**字段说明：**
- `name` (必填): 场景名称，最大100字符
- `tag_id` (必填): 场景标签ID（关联到tags表）

**示例请求：**
```bash
curl -X POST http://localhost:8080/api/v1/scenarios \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "企业品牌监测",
    "tag_id": 1
  }'
```

**响应示例：**
```json
{
  "message": "创建成功",
  "data": {
    "id": 1,
    "name": "企业品牌监测",
    "tag_id": 1,
    "status": 1,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 2. 获取场景列表

**接口地址：** `GET /api/v1/scenarios`

**认证要求：** 需要登录

**查询参数：**
- `status` (可选): 状态筛选，`active` 或 `1` 表示只获取启用的场景

**示例请求：**
```bash
curl -X GET "http://localhost:8080/api/v1/scenarios?status=active" \
  -H "Authorization: Bearer <token>"
```

### 3. 获取场景详情

**接口地址：** `GET /api/v1/scenarios/:id`

**认证要求：** 需要登录

### 4. 获取场景及其监测组

**接口地址：** `GET /api/v1/scenarios/:id/groups`

**认证要求：** 需要登录

**响应示例：**
```json
{
  "data": {
    "id": 1,
    "name": "企业品牌监测",
    "tag_id": 1,
    "tag": {
      "id": 1,
      "name": "企业",
      "code": "enterprise"
    },
    "groups": [
      {
        "id": 1,
        "name": "第1组",
        "channels": [...],
        "keywords": [...],
        "exclusion_words": [...]
      }
    ]
  }
}
```

### 5. 更新场景

**接口地址：** `PUT /api/v1/scenarios/:id`

**认证要求：** 需要登录且具有 admin 角色

**请求体：**
```json
{
  "name": "更新后的场景名",
  "tag_id": 2,
  "status": 1
}
```

### 6. 删除场景

**接口地址：** `DELETE /api/v1/scenarios/:id`

**认证要求：** 需要登录且具有 admin 角色

**注意：** 删除场景会级联删除所有关联的监测组及其配置

## 监测组管理 API

### 1. 创建监测组

**接口地址：** `POST /api/v1/monitoring-groups`

**认证要求：** 需要登录且具有 admin 角色

**请求体：**
```json
{
  "scenario_id": 1,
  "name": "第1组",
  "sort": 1
}
```

**字段说明：**
- `scenario_id` (必填): 所属场景ID
- `name` (必填): 监测组名称，最大100字符
- `sort` (可选): 排序值，默认0

**示例请求：**
```bash
curl -X POST http://localhost:8080/api/v1/monitoring-groups \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "scenario_id": 1,
    "name": "第1组",
    "sort": 1
  }'
```

### 2. 根据场景ID获取监测组列表

**接口地址：** `GET /api/v1/monitoring-groups/scenario/:scenario_id`

**认证要求：** 需要登录

### 3. 获取监测组详情

**接口地址：** `GET /api/v1/monitoring-groups/:id`

**认证要求：** 需要登录

**响应示例：**
```json
{
  "data": {
    "id": 1,
    "scenario_id": 1,
    "name": "第1组",
    "channels": [
      {
        "id": 1,
        "name": "小红书",
        "code": "xiaohongshu"
      },
      {
        "id": 2,
        "name": "微博",
        "code": "weibo"
      }
    ],
    "keywords": [
      {
        "id": 1,
        "keyword": "品牌名"
      }
    ],
    "exclusion_words": [
      {
        "id": 1,
        "word": "广告"
      }
    ]
  }
}
```

### 4. 更新监测组

**接口地址：** `PUT /api/v1/monitoring-groups/:id`

**认证要求：** 需要登录且具有 admin 角色

**请求体：**
```json
{
  "name": "更新后的组名",
  "sort": 2,
  "status": 1
}
```

### 5. 删除监测组

**接口地址：** `DELETE /api/v1/monitoring-groups/:id`

**认证要求：** 需要登录且具有 admin 角色

**注意：** 删除监测组会级联删除所有关联的渠道、关键词和排除词

### 6. 分配渠道

**接口地址：** `POST /api/v1/monitoring-groups/:id/channels`

**认证要求：** 需要登录且具有 admin 角色

**请求体：**
```json
{
  "channel_ids": [1, 2, 3]
}
```

**示例请求：**
```bash
curl -X POST http://localhost:8080/api/v1/monitoring-groups/1/channels \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "channel_ids": [1, 2, 3]
  }'
```

**注意：** 此操作会替换现有的渠道绑定，不是追加

### 7. 添加关键词

**接口地址：** `POST /api/v1/monitoring-groups/:id/keywords`

**认证要求：** 需要登录且具有 admin 角色

**请求体：**
```json
{
  "keyword": "品牌名"
}
```

**示例请求：**
```bash
curl -X POST http://localhost:8080/api/v1/monitoring-groups/1/keywords \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "keyword": "品牌名"
  }'
```

### 8. 删除关键词

**接口地址：** `DELETE /api/v1/monitoring-groups/:id/keywords/:keyword_id`

**认证要求：** 需要登录且具有 admin 角色

### 9. 获取关键词列表

**接口地址：** `GET /api/v1/monitoring-groups/:id/keywords`

**认证要求：** 需要登录

**响应示例：**
```json
{
  "data": [
    {
      "id": 1,
      "group_id": 1,
      "keyword": "品牌名",
      "created_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

### 10. 添加排除词

**接口地址：** `POST /api/v1/monitoring-groups/:id/exclusion-words`

**认证要求：** 需要登录且具有 admin 角色

**请求体：**
```json
{
  "word": "广告"
}
```

**示例请求：**
```bash
curl -X POST http://localhost:8080/api/v1/monitoring-groups/1/exclusion-words \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "word": "广告"
  }'
```

### 11. 删除排除词

**接口地址：** `DELETE /api/v1/monitoring-groups/:id/exclusion-words/:word_id`

**认证要求：** 需要登录且具有 admin 角色

### 12. 获取排除词列表

**接口地址：** `GET /api/v1/monitoring-groups/:id/exclusion-words`

**认证要求：** 需要登录

## 完整使用流程示例

### 1. 创建场景

```bash
# 创建场景
SCENARIO=$(curl -X POST http://localhost:8080/api/v1/scenarios \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "企业品牌监测",
    "tag_id": 1
  }' | jq -r '.data.id')
```

### 2. 创建监测组

```bash
# 创建监测组
GROUP=$(curl -X POST http://localhost:8080/api/v1/monitoring-groups \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d "{
    \"scenario_id\": $SCENARIO,
    \"name\": \"第1组\",
    \"sort\": 1
  }" | jq -r '.data.id')
```

### 3. 配置监测组

```bash
# 分配渠道
curl -X POST http://localhost:8080/api/v1/monitoring-groups/$GROUP/channels \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "channel_ids": [1, 2, 3]
  }'

# 添加关键词
curl -X POST http://localhost:8080/api/v1/monitoring-groups/$GROUP/keywords \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"keyword": "品牌名"}'

curl -X POST http://localhost:8080/api/v1/monitoring-groups/$GROUP/keywords \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"keyword": "产品名"}'

# 添加排除词
curl -X POST http://localhost:8080/api/v1/monitoring-groups/$GROUP/exclusion-words \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"word": "广告"}'

curl -X POST http://localhost:8080/api/v1/monitoring-groups/$GROUP/exclusion-words \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"word": "推广"}'
```

### 4. 查看完整配置

```bash
# 获取场景及其所有监测组配置
curl -X GET http://localhost:8080/api/v1/scenarios/$SCENARIO/groups \
  -H "Authorization: Bearer $TOKEN"
```

## 权限说明

- **查看场景和监测组**：需要登录认证
- **创建/更新/删除场景和监测组**：需要 admin 角色
- **配置监测组（渠道、关键词、排除词）**：需要 admin 角色

## 错误码说明

- `400` - 请求参数错误
- `401` - 未认证或 token 无效
- `403` - 无权限访问（需要 admin 角色）
- `404` - 资源不存在
- `500` - 服务器内部错误

