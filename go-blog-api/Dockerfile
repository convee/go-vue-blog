# Stage 1: Build the Go app
FROM golang:alpine AS builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod tidy

COPY . .

RUN go build -ldflags="-s -w" -o /app/blog ./main.go

# Stage 2: Run the Go app
FROM alpine:latest
ENV TZ Asia/Shanghai

WORKDIR /app
EXPOSE 80

COPY --from=builder /app/blog /app/blog

CMD ["./blog"]