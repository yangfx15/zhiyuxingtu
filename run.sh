#!/bin/bash

# 默认参数
COVER_CONTAINER=${1:-"true"}
IMAGE_NAME="zhiyuxingtu"
CONTAINER_NAME="zhiyuxingtu"
PORT=18080

# 构建 Docker 镜像
echo "🛠️  Building Docker image..."
docker build -t $IMAGE_NAME:latest .

# 覆盖容器逻辑
if [ "$COVER_CONTAINER" = "true" ]; then
  echo "🔫 Stopping and removing existing container..."
  docker rm -f $CONTAINER_NAME 2>/dev/null || true
fi

# 运行新容器
echo "🚀 Starting new container..."
docker run -d \
  --name $CONTAINER_NAME \
  -p $PORT:18080 \
  --restart unless-stopped \
  $IMAGE_NAME:latest

echo "✅ Deployment completed! Access via http://localhost:$PORT"
