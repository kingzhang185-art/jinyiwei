# Docker 镜像加速器配置指南

## 问题说明

如果遇到以下错误：
```
Error response from daemon: failed to resolve reference "docker.io/library/mysql:8.0": 
failed to do request: Head "https://registry-1.docker.io/v2/library/mysql/manifests/8.0": 
context deadline exceeded
```

这是因为无法连接到 Docker Hub，需要配置国内镜像加速器。

## 解决方案

### 方法一：使用配置脚本（推荐）

运行项目根目录下的配置脚本：

```bash
./setup-docker-mirror.sh
```

然后重启 Docker Desktop 使配置生效。

### 方法二：手动配置

#### macOS / Linux

1. 编辑或创建 `~/.docker/daemon.json` 文件：

```bash
mkdir -p ~/.docker
nano ~/.docker/daemon.json
```

2. 添加以下内容：

```json
{
  "registry-mirrors": [
    "https://docker.mirrors.ustc.edu.cn",
    "https://hub-mirror.c.163.com",
    "https://mirror.baidubce.com"
  ]
}
```

3. 重启 Docker Desktop（macOS）或 Docker 服务（Linux）：

**macOS:**
- 点击 Docker Desktop 图标 → Restart

**Linux:**
```bash
sudo systemctl restart docker
```

### 方法三：使用 Docker Desktop GUI（macOS）

1. 打开 Docker Desktop
2. 点击设置图标（齿轮）
3. 选择 "Docker Engine"
4. 在 JSON 配置中添加：

```json
{
  "registry-mirrors": [
    "https://docker.mirrors.ustc.edu.cn",
    "https://hub-mirror.c.163.com",
    "https://mirror.baidubce.com"
  ]
}
```

5. 点击 "Apply & Restart"

## 验证配置

配置完成后，运行以下命令验证：

```bash
docker info | grep -A 10 "Registry Mirrors"
```

应该能看到配置的镜像加速器地址。

## 启动服务

配置完成后，运行：

```bash
docker-compose up -d
```

## 常用镜像加速器地址

- 中科大镜像：`https://docker.mirrors.ustc.edu.cn`
- 网易镜像：`https://hub-mirror.c.163.com`
- 百度云镜像：`https://mirror.baidubce.com`
- 阿里云镜像：需要登录阿里云获取专属地址

## 如果仍然无法连接

1. 检查网络连接
2. 尝试使用代理
3. 手动拉取镜像：
   ```bash
   docker pull docker.mirrors.ustc.edu.cn/library/mysql:8.0
   docker tag docker.mirrors.ustc.edu.cn/library/mysql:8.0 mysql:8.0
   ```

