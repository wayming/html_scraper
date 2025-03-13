# Golang 客户端 Dockerfile
FROM golang:1.24

WORKDIR /app
COPY client/ /app

RUN go mod tidy

CMD ["go", "run", "main.go"]
