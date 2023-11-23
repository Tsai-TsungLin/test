# 第一階段
FROM golang:1.19.1-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

#ENV GOPROXY=https://goproxy.io // go mod download 失敗請把註解拿掉

COPY go.mod go.sum ./

RUN go mod download

# 將本機的 pkg 目錄複製到容器中
# COPY ./inventories/host_metadata.go /go/pkg/mod/github.com/!data!dog/datadog-agent@v0.0.0-20230726050554-2ee47b918b01/pkg/metadata/inventories/host_metadata.go

COPY . .

RUN go build -o test

# 第二階段
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/test .
# COPY --from=builder /app/yaml ./yaml
# COPY --from=builder /app/datadog ./datadog

EXPOSE 50051

CMD ["./test"]
