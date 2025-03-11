FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o crm.shopdev.com ./cmd/server

FROM alpine:latest

COPY ./config /config

COPY --from=builder /build/crm.shopdev.com /

# Set environment variables for database connections
ENV MYSQL_HOST=mysql_db
ENV MYSQL_PORT=3306
ENV REDIS_HOST=redis_db
ENV REDIS_PORT=6379

ENTRYPOINT ["/crm.shopdev.com"]