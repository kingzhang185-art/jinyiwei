# RBAC 用户权限管理 API 文档

## 概述

本项目实现了基于角色的访问控制（RBAC）系统，包含用户管理、角色管理和权限管理功能。

## 认证方式

所有需要认证的接口都需要在请求头中携带 JWT Token：

```
Authorization: Bearer <token>
```

## API 接口

### 1. 认证相关接口

#### 1.1 用户注册

**接口地址：** `POST /api/v1/auth/register`

**请求体：**
```json
{
  "username": "admin",
  "password": "123456",
  "email": "admin@example.com",
  "nickname": "管理员"
}
```

**响应示例：**
```json
{
  "message": "注册成功",
  "data": {
    "id": 1,
    "username": "admin",
    "email": "admin@example.com",
    "nickname": "管理员"
  }
}
```

#### 1.2 用户登录

**接口地址：** `POST /api/v1/auth/login`

**请求体：**
```json
{
  "username": "admin",
  "password": "123456"
}
```

**响应示例：**
```json
{
  "message": "登录成功",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "admin",
      "email": "admin@example.com",
      "nickname": "管理员"
    }
  }
}
```

#### 1.3 获取当前用户信息

**接口地址：** `GET /api/v1/auth/me`

**请求头：**
```
Authorization: Bearer <token>
```

**响应示例：**
```json
{
  "data": {
    "id": 1,
    "username": "admin",
    "email": "admin@example.com",
    "nickname": "管理员",
    "roles": [
      {
        "id": 1,
        "name": "管理员",
        "code": "admin",
        "permissions": [...]
      }
    ]
  }
}
```

#### 1.4 修改密码

**接口地址：** `PUT /api/v1/auth/password`

**请求头：**
```
Authorization: Bearer <token>
```

**请求体：**
```json
{
  "old_password": "123456",
  "new_password": "newpassword"
}
```

### 2. 用户管理接口（需要 admin 角色）

#### 2.1 创建用户

**接口地址：** `POST /api/v1/users`

**请求体：**
```json
{
  "username": "testuser",
  "password": "123456",
  "email": "test@example.com",
  "nickname": "测试用户"
}
```

#### 2.2 获取用户列表

**接口地址：** `GET /api/v1/users?page=1&page_size=10`

**查询参数：**
- `page`: 页码（默认：1）
- `page_size`: 每页数量（默认：10）

#### 2.3 获取用户详情

**接口地址：** `GET /api/v1/users/:id`

#### 2.4 更新用户

**接口地址：** `PUT /api/v1/users/:id`

**请求体：**
```json
{
  "email": "newemail@example.com",
  "nickname": "新昵称",
  "status": 1
}
```

#### 2.5 删除用户

**接口地址：** `DELETE /api/v1/users/:id`

#### 2.6 分配角色

**接口地址：** `POST /api/v1/users/:id/roles`

**请求体：**
```json
{
  "role_ids": [1, 2]
}
```

### 3. 角色管理接口（需要 admin 角色）

#### 3.1 创建角色

**接口地址：** `POST /api/v1/roles`

**请求体：**
```json
{
  "name": "编辑",
  "code": "editor",
  "description": "内容编辑角色"
}
```

#### 3.2 获取角色列表

**接口地址：** `GET /api/v1/roles`

#### 3.3 获取角色详情

**接口地址：** `GET /api/v1/roles/:id`

#### 3.4 更新角色

**接口地址：** `PUT /api/v1/roles/:id`

**请求体：**
```json
{
  "name": "新角色名",
  "description": "新描述",
  "status": 1
}
```

#### 3.5 删除角色

**接口地址：** `DELETE /api/v1/roles/:id`

#### 3.6 分配权限

**接口地址：** `POST /api/v1/roles/:id/permissions`

**请求体：**
```json
{
  "permission_ids": [1, 2, 3]
}
```

### 4. 权限管理接口（需要 admin 角色）

#### 4.1 创建权限

**接口地址：** `POST /api/v1/permissions`

**请求体：**
```json
{
  "name": "查看舆情",
  "code": "opinion:view",
  "method": "GET",
  "path": "/api/v1/opinions",
  "description": "查看舆情列表的权限"
}
```

#### 4.2 获取权限列表

**接口地址：** `GET /api/v1/permissions`

#### 4.3 获取权限详情

**接口地址：** `GET /api/v1/permissions/:id`

#### 4.4 更新权限

**接口地址：** `PUT /api/v1/permissions/:id`

**请求体：**
```json
{
  "name": "新权限名",
  "method": "POST",
  "path": "/api/v1/opinions",
  "description": "新描述",
  "status": 1
}
```

#### 4.5 删除权限

**接口地址：** `DELETE /api/v1/permissions/:id`

## 使用示例

### 1. 注册管理员账号

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin123",
    "email": "admin@example.com",
    "nickname": "管理员"
  }'
```

### 2. 登录获取 Token

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin123"
  }'
```

### 3. 使用 Token 访问受保护接口

```bash
curl -X GET http://localhost:8080/api/v1/auth/me \
  -H "Authorization: Bearer <your_token>"
```

### 4. 创建角色并分配权限

```bash
# 创建角色
curl -X POST http://localhost:8080/api/v1/roles \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "编辑",
    "code": "editor",
    "description": "内容编辑"
  }'

# 分配权限给角色
curl -X POST http://localhost:8080/api/v1/roles/2/permissions \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "permission_ids": [9, 10]
  }'
```

### 5. 为用户分配角色

```bash
curl -X POST http://localhost:8080/api/v1/users/2/roles \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "role_ids": [2]
  }'
```

## 默认数据

系统初始化时会自动创建：

1. **默认角色：**
   - `admin` - 管理员（拥有所有权限）
   - `user` - 普通用户（基础权限）

2. **默认权限：**
   - `user:manage` - 用户管理
   - `user:create` - 创建用户
   - `user:update` - 更新用户
   - `user:delete` - 删除用户
   - `role:manage` - 角色管理
   - `role:create` - 创建角色
   - `permission:manage` - 权限管理
   - `permission:create` - 创建权限
   - `opinion:view` - 舆情查看
   - `opinion:create` - 舆情创建

## 权限说明

- **公开接口：** 注册、登录、健康检查
- **需要认证：** 所有 `/api/v1/auth/me` 及舆情相关接口
- **需要 admin 角色：** 用户管理、角色管理、权限管理相关接口

## 错误码说明

- `400` - 请求参数错误
- `401` - 未认证或 token 无效
- `403` - 无权限访问
- `404` - 资源不存在
- `500` - 服务器内部错误

