# 第一阶段：构建 Go 二进制文件
FROM golang:1.21-alpine AS builder

# 安装必要的构建工具
RUN apk add --no-cache git ca-certificates

# 设置工作目录
WORKDIR /app

# 复制依赖文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建可执行文件（根据你的项目入口调整）
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./main.go

# 第二阶段：生成最小化运行镜像
FROM alpine:3.19

# 复制可执行文件和证书
COPY --from=builder /app/main /app/main
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# 设置工作目录和暴露端口
WORKDIR /app
EXPOSE 18080

# 运行服务
CMD ["/app/main"]
