# step 1
FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/goaway

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY config /app/config
COPY data /app/data
RUN go build -ldflags="-s -w" -o /app/goaway main.go

# step 2
FROM alpine

RUN sed -i 's/https/http/' /etc/apk/repositories
RUN apk add curl
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/goaway /app/goaway
COPY --from=builder /app/config /app/config
COPY --from=builder /app/data /app/data

CMD ["./goaway", "-f", "config/config.toml"]
