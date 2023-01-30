FROM golang:alpine AS builder
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build
COPY . .
RUN go mod tidy
RUN go build -ldflags="-s -w" -o gin-mysqlbak-agent ./main.go

FROM centos
WORKDIR /app
ENV TZ Asia/Shanghai
COPY --from=builder /build/gin-mysqlbak-agent /app/gin-mysqlbak-agent
COPY --from=builder /build/domain/config/config.ini /app/domain/config/config.ini
COPY --from=builder /build/domain/template /app/domain/template
COPY --from=builder /build/docker/mysqldump /usr/bin
COPY --from=builder /build/docker/mysqladmin /usr/bin
RUN chmod 777 /usr/bin/mysqldump && chmod 777 /usr/bin/mysqladmin
CMD ["./gin-mysqlbak-agent"]