FROM golang:1.21-alpine AS builder

COPY . /github.com/Ivanrumanchev/chat-server/source/
WORKDIR /github.com/Ivanrumanchev/chat-server/source/

RUN go mod download
RUN go build -o ./bin/crud_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/Ivanrumanchev/chat-server/source/bin/crud_server .

CMD ["./crud_server"]