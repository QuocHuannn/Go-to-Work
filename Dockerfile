FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o crm.shopdev.com ./cmd/server

FROM alpine:latest

RUN apk add --no-cache bash netcat-openbsd

# Copy .env file and config directory
COPY .env /
COPY ./config /config
COPY ./templates /templates
# Copy wait script
COPY --from=builder /build/crm.shopdev.com /

# Set environment variables for database connections
ENV MYSQL_HOST=mysql_db
ENV MYSQL_PORT=3306
ENV REDIS_HOST=redis
ENV REDIS_PORT=6379

# Add a script to wait for MySQL to be ready
RUN printf '#!/bin/sh\n\
echo "Waiting for MySQL to be ready..."\n\
until nc -z $MYSQL_HOST $MYSQL_PORT; do\n\
  echo "MySQL not available yet - sleeping"\n\
  sleep 2\n\
done\n\
echo "MySQL is up - starting app"\n\
exec /crm.shopdev.com\n' > /wait-for-mysql.sh && chmod +x /wait-for-mysql.sh

ENTRYPOINT ["/wait-for-mysql.sh"]