# 标签管理 API 文档

## 概述

标签管理系统支持场景标签的增删改查功能。系统预设了6个场景标签：企业、品牌、商品、产品、人物、其他。

## API 接口

### 1. 获取标签列表

**接口地址：** `GET /api/v1/tags`

**认证要求：** 需要登录（Bearer Token）

**查询参数：**
- `type` (可选): 标签类型，如 `scene`
- `status` (可选): 状态筛选，`active` 或 `1` 表示只获取启用的标签

**示例请求：**
```bash
# 获取所有标签
curl -X GET http://localhost:8080/api/v1/tags \
  -H "Authorization: Bearer <token>"

# 获取场景标签
curl -X GET "http://localhost:8080/api/v1/tags?type=scene" \
  -H "Authorization: Bearer <token>"

# 只获取启用的标签
curl -X GET "http://localhost:8080/api/v1/tags?status=active" \
  -H "Authorization: Bearer <token>"
```

**响应示例：**
```json
{
  "data": [
    {
      "id": 1,
      "name": "企业",
      "code": "enterprise",
      "description": "企业相关场景标签",
      "type": "scene",
      "sort": 1,
      "status": 1,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    },
    {
      "id": 2,
      "name": "品牌",
      "code": "brand",
      "description": "品牌相关场景标签",
      "type": "scene",
      "sort": 2,
      "status": 1,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

### 2. 获取标签详情

**接口地址：** `GET /api/v1/tags/:id`

**认证要求：** 需要登录（Bearer Token）

**路径参数：**
- `id`: 标签ID

**示例请求：**
```bash
curl -X GET http://localhost:8080/api/v1/tags/1 \
  -H "Authorization: Bearer <token>"
```

**响应示例：**
```json
{
  "data": {
    "id": 1,
    "name": "企业",
    "code": "enterprise",
    "description": "企业相关场景标签",
    "type": "scene",
    "sort": 1,
    "status": 1,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 3. 创建标签（需要 admin 角色）

**接口地址：** `POST /api/v1/tags`

**认证要求：** 需要登录且具有 admin 角色

**请求体：**
```json
{
  "name": "新标签",
  "code": "new_tag",
  "description": "新标签描述",
  "type": "scene",
  "sort": 10
}
```

**字段说明：**
- `name` (必填): 标签名称，最大50字符
- `code` (必填): 标签代码，最大50字符，必须唯一
- `description` (可选): 标签描述，最大255字符
- `type` (可选): 标签类型，默认 `scene`
- `sort` (可选): 排序值，默认0

**示例请求：**
```bash
curl -X POST http://localhost:8080/api/v1/tags \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "新标签",
    "code": "new_tag",
    "description": "新标签描述",
    "type": "scene",
    "sort": 10
  }'
```

**响应示例：**
```json
{
  "message": "创建成功",
  "data": {
    "id": 7,
    "name": "新标签",
    "code": "new_tag",
    "description": "新标签描述",
    "type": "scene",
    "sort": 10,
    "status": 1,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 4. 更新标签（需要 admin 角色）

**接口地址：** `PUT /api/v1/tags/:id`

**认证要求：** 需要登录且具有 admin 角色

**路径参数：**
- `id`: 标签ID

**请求体：**
```json
{
  "name": "更新后的标签名",
  "description": "更新后的描述",
  "sort": 5,
  "status": 1
}
```

**字段说明：**
- `name` (可选): 标签名称
- `description` (可选): 标签描述
- `sort` (可选): 排序值
- `status` (可选): 状态，1-正常，2-禁用

**示例请求：**
```bash
curl -X PUT http://localhost:8080/api/v1/tags/1 \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "更新后的企业标签",
    "description": "更新后的描述",
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

### 5. 删除标签（需要 admin 角色）

**接口地址：** `DELETE /api/v1/tags/:id`

**认证要求：** 需要登录且具有 admin 角色

**路径参数：**
- `id`: 标签ID

**示例请求：**
```bash
curl -X DELETE http://localhost:8080/api/v1/tags/7 \
  -H "Authorization: Bearer <token>"
```

**响应示例：**
```json
{
  "message": "删除成功"
}
```

## 预设标签数据

系统初始化时会自动创建以下6个场景标签：

| ID | 名称 | 代码 | 类型 | 排序 |
|----|------|------|------|------|
| 1 | 企业 | enterprise | scene | 1 |
| 2 | 品牌 | brand | scene | 2 |
| 3 | 商品 | commodity | scene | 3 |
| 4 | 产品 | product | scene | 4 |
| 5 | 人物 | person | scene | 5 |
| 6 | 其他 | other | scene | 6 |

## 权限说明

- **查看标签**：需要登录认证
- **创建/更新/删除标签**：需要 admin 角色

## 错误码说明

- `400` - 请求参数错误
- `401` - 未认证或 token 无效
- `403` - 无权限访问（需要 admin 角色）
- `404` - 标签不存在
- `500` - 服务器内部错误

## 使用示例

### 完整流程示例

```bash
# 1. 登录获取 token
TOKEN=$(curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}' \
  | jq -r '.data.token')

# 2. 获取所有启用的场景标签
curl -X GET "http://localhost:8080/api/v1/tags?type=scene&status=active" \
  -H "Authorization: Bearer $TOKEN"

# 3. 创建新标签（需要 admin 角色）
curl -X POST http://localhost:8080/api/v1/tags \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "活动",
    "code": "activity",
    "description": "活动相关场景标签",
    "type": "scene",
    "sort": 7
  }'

# 4. 更新标签
curl -X PUT http://localhost:8080/api/v1/tags/7 \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "sort": 1,
    "status": 1
  }'

# 5. 删除标签
curl -X DELETE http://localhost:8080/api/v1/tags/7 \
  -H "Authorization: Bearer $TOKEN"
```

