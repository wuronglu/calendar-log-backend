# 使用支持 Go 1.23 的基础镜像
FROM golang:1.23

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制项目文件
COPY . .

# 编译项目
RUN go build -o main .

# 运行应用
CMD ["./main"]
