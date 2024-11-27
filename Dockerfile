# 使用 Go 官方镜像作为构建环境
FROM golang:1.21 AS builder

# 设置工作目录
WORKDIR /app

# 将 go.mod 和 go.sum 复制到容器中
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 将整个项目复制到容器
COPY . .

# 编译应用
RUN GOOS=linux GOARCH=amd64 go build -o calendar-log-backend .

# 使用更小的镜像运行应用
FROM alpine:latest

# 安装 ca-certificates（为了 SSL 支持）
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /root/

# 从构建镜像复制二进制文件
COPY --from=builder /app/calendar-log-backend .

# 设置容器启动命令
CMD ["./calendar-log-backend"]

# 暴露后端服务端口
EXPOSE 9191
