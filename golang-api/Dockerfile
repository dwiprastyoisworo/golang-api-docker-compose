FROM golang:1.23.3-alpine AS builder

# Mengatur direktori kerja
WORKDIR /app

# Menyalin semua file ke dalam container
COPY . .

# Mengunduh dependencies Go
RUN go mod download

#build binary
RUN go build -o main

FROM alpine:3.20 AS base

WORKDIR /app

#create user
RUN addgroup -S goapi && adduser -S goapi -G goapi

#use new user
USER goapi

# Menambahkan environment variables untuk database
ENV POSTGRES_DB_HOST=postgres \
    POSTGRES_DB_USERNAME=root \
    POSTGRES_DB_PASSWORD=root \
    POSTGRES_DB_NAME=mydatabase \
    POSTGRES_DB_SCHEMA=public \
    POSTGRES_DB_PORT=5432 \
    POSTGRES_DB_SSL=disable \
    PORT=8070

#copy main binary and env
COPY --from=builder /app/main /app/.env ./

#copy file migration
COPY --from=builder /app/migration ./migration

#expose port
EXPOSE 8087

CMD ["./main"]