-- 创建数据库（如果不存在）
CREATE DATABASE IF NOT EXISTS opinion_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 使用数据库
USE opinion_db;

-- 创建用户表
CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE COMMENT '用户名',
    password VARCHAR(255) NOT NULL COMMENT '密码（加密）',
    email VARCHAR(100) UNIQUE COMMENT '邮箱',
    nickname VARCHAR(50) COMMENT '昵称',
    status TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1-正常，2-禁用',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX idx_username (username),
    INDEX idx_email (email),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 创建角色表
CREATE TABLE IF NOT EXISTS roles (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL COMMENT '角色名称',
    code VARCHAR(50) NOT NULL UNIQUE COMMENT '角色代码',
    description VARCHAR(255) COMMENT '角色描述',
    status TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1-正常，2-禁用',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX idx_code (code),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

-- 创建权限表
CREATE TABLE IF NOT EXISTS permissions (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL COMMENT '权限名称',
    code VARCHAR(100) NOT NULL UNIQUE COMMENT '权限代码',
    method VARCHAR(10) COMMENT 'HTTP方法',
    path VARCHAR(255) COMMENT 'API路径',
    description VARCHAR(255) COMMENT '权限描述',
    status TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1-正常，2-禁用',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX idx_code (code),
    INDEX idx_method_path (method, path),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='权限表';

-- 创建用户角色关联表
CREATE TABLE IF NOT EXISTS user_roles (
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    role_id BIGINT UNSIGNED NOT NULL COMMENT '角色ID',
    PRIMARY KEY (user_id, role_id),
    INDEX idx_user_id (user_id),
    INDEX idx_role_id (role_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户角色关联表';

-- 创建角色权限关联表
CREATE TABLE IF NOT EXISTS role_permissions (
    role_id BIGINT UNSIGNED NOT NULL COMMENT '角色ID',
    permission_id BIGINT UNSIGNED NOT NULL COMMENT '权限ID',
    PRIMARY KEY (role_id, permission_id),
    INDEX idx_role_id (role_id),
    INDEX idx_permission_id (permission_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色权限关联表';

-- 创建标签表
CREATE TABLE IF NOT EXISTS tags (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL COMMENT '标签名称',
    code VARCHAR(50) NOT NULL UNIQUE COMMENT '标签代码',
    description VARCHAR(255) COMMENT '标签描述',
    type VARCHAR(20) NOT NULL DEFAULT 'scene' COMMENT '标签类型:scene-场景标签',
    sort INT NOT NULL DEFAULT 0 COMMENT '排序',
    status TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1-正常，2-禁用',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX idx_code (code),
    INDEX idx_type (type),
    INDEX idx_status (status),
    INDEX idx_sort (sort)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='标签表';

-- 创建渠道表
CREATE TABLE IF NOT EXISTS channels (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL COMMENT '渠道名称',
    code VARCHAR(50) NOT NULL UNIQUE COMMENT '渠道代码',
    description VARCHAR(255) COMMENT '渠道描述',
    icon VARCHAR(255) COMMENT '图标URL或图标标识',
    sort INT NOT NULL DEFAULT 0 COMMENT '排序',
    status TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1-正常，2-禁用',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX idx_code (code),
    INDEX idx_status (status),
    INDEX idx_sort (sort)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='渠道表';

-- 创建场景表
CREATE TABLE IF NOT EXISTS scenarios (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL COMMENT '场景名称',
    tag_id BIGINT UNSIGNED NOT NULL COMMENT '场景标签ID',
    status TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1-正常，2-禁用',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX idx_tag_id (tag_id),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='监测场景表';

-- 创建监测组表
CREATE TABLE IF NOT EXISTS monitoring_groups (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    scenario_id BIGINT UNSIGNED NOT NULL COMMENT '所属场景ID',
    name VARCHAR(100) NOT NULL COMMENT '监测组名称',
    sort INT NOT NULL DEFAULT 0 COMMENT '排序',
    status TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1-正常，2-禁用',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX idx_scenario_id (scenario_id),
    INDEX idx_status (status),
    INDEX idx_sort (sort)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='监测组表';

-- 创建监测组-渠道关联表
CREATE TABLE IF NOT EXISTS group_channels (
    group_id BIGINT UNSIGNED NOT NULL COMMENT '监测组ID',
    channel_id BIGINT UNSIGNED NOT NULL COMMENT '渠道ID',
    PRIMARY KEY (group_id, channel_id),
    INDEX idx_group_id (group_id),
    INDEX idx_channel_id (channel_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='监测组-渠道关联表';

-- 创建监测组关键词表
CREATE TABLE IF NOT EXISTS group_keywords (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    group_id BIGINT UNSIGNED NOT NULL COMMENT '监测组ID',
    keyword VARCHAR(255) NOT NULL COMMENT '关键词',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    INDEX idx_group_id (group_id),
    INDEX idx_keyword (keyword)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='监测组关键词表';

-- 创建监测组排除词表
CREATE TABLE IF NOT EXISTS group_exclusion_words (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    group_id BIGINT UNSIGNED NOT NULL COMMENT '监测组ID',
    word VARCHAR(255) NOT NULL COMMENT '排除词',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    INDEX idx_group_id (group_id),
    INDEX idx_word (word)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='监测组排除词表';

-- 创建舆情表
CREATE TABLE IF NOT EXISTS opinions (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    content TEXT NOT NULL COMMENT '舆情内容',
    source VARCHAR(255) NOT NULL COMMENT '来源',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX idx_source (source),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='舆情表';

-- 插入默认管理员角色
INSERT INTO roles (name, code, description, status) VALUES
('管理员', 'admin', '系统管理员，拥有所有权限', 1),
('普通用户', 'user', '普通用户，基础权限', 1)
ON DUPLICATE KEY UPDATE name=VALUES(name);

-- 插入默认权限
INSERT INTO permissions (name, code, method, path, description, status) VALUES
('用户管理', 'user:manage', 'GET', '/api/v1/users', '查看用户列表', 1),
('创建用户', 'user:create', 'POST', '/api/v1/users', '创建用户', 1),
('更新用户', 'user:update', 'PUT', '/api/v1/users/:id', '更新用户信息', 1),
('删除用户', 'user:delete', 'DELETE', '/api/v1/users/:id', '删除用户', 1),
('角色管理', 'role:manage', 'GET', '/api/v1/roles', '查看角色列表', 1),
('创建角色', 'role:create', 'POST', '/api/v1/roles', '创建角色', 1),
('权限管理', 'permission:manage', 'GET', '/api/v1/permissions', '查看权限列表', 1),
('创建权限', 'permission:create', 'POST', '/api/v1/permissions', '创建权限', 1),
('舆情查看', 'opinion:view', 'GET', '/api/v1/opinions', '查看舆情列表', 1),
('舆情创建', 'opinion:create', 'POST', '/api/v1/opinions', '创建舆情', 1)
ON DUPLICATE KEY UPDATE name=VALUES(name);

-- 为管理员角色分配所有权限
INSERT INTO role_permissions (role_id, permission_id)
SELECT 1, id FROM permissions WHERE status = 1
ON DUPLICATE KEY UPDATE role_id=VALUES(role_id);

-- 插入预设场景标签数据
INSERT INTO tags (name, code, description, type, sort, status) VALUES
('企业', 'enterprise', '企业相关场景标签', 'scene', 1, 1),
('品牌', 'brand', '品牌相关场景标签', 'scene', 2, 1),
('商品', 'commodity', '商品相关场景标签', 'scene', 3, 1),
('产品', 'product', '产品相关场景标签', 'scene', 4, 1),
('人物', 'person', '人物相关场景标签', 'scene', 5, 1),
('其他', 'other', '其他场景标签', 'scene', 6, 1)
ON DUPLICATE KEY UPDATE name=VALUES(name);

-- 插入预设渠道数据
INSERT INTO channels (name, code, description, icon, sort, status) VALUES
('小红书', 'xiaohongshu', '小红书平台', 'xiaohongshu', 1, 1),
('微博', 'weibo', '微博平台', 'weibo', 2, 1),
('快手', 'kuaishou', '快手平台', 'kuaishou', 3, 1),
('抖音', 'douyin', '抖音平台', 'douyin', 4, 1),
('西瓜视频', 'xigua', '西瓜视频平台', 'xigua', 5, 1),
('百度贴吧', 'tieba', '百度贴吧平台', 'tieba', 6, 1),
('知乎', 'zhihu', '知乎平台', 'zhihu', 7, 1),
('虎扑', 'hupu', '虎扑平台', 'hupu', 8, 1)
ON DUPLICATE KEY UPDATE name=VALUES(name);

