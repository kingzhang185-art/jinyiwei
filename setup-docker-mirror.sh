#!/bin/bash

# Docker 镜像加速器配置脚本

echo "正在配置 Docker 镜像加速器..."

# 备份现有配置
if [ -f ~/.docker/daemon.json ]; then
    cp ~/.docker/daemon.json ~/.docker/daemon.json.backup
    echo "已备份现有配置到 ~/.docker/daemon.json.backup"
fi

# 创建或更新 daemon.json
cat > ~/.docker/daemon.json << 'EOF'
{
  "builder": {
    "gc": {
      "defaultKeepStorage": "20GB",
      "enabled": true
    }
  },
  "experimental": false,
  "registry-mirrors": [
    "https://docker.mirrors.ustc.edu.cn",
    "https://hub-mirror.c.163.com",
    "https://mirror.baidubce.com"
  ]
}
EOF

echo "✅ Docker 镜像加速器配置完成！"
echo ""
echo "请重启 Docker Desktop 以使配置生效："
echo "1. 点击 Docker Desktop 图标"
echo "2. 选择 'Restart' 或 'Quit Docker Desktop' 然后重新启动"
echo ""
echo "配置完成后，运行以下命令启动服务："
echo "  docker-compose up -d"

