# сборка
FROM golang:1.25.3-alpine AS builder   

WORKDIR /app 

# Копируем файлы зависимостей и скачиваем модули
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -o user-service ./cmd/main.go

# Этап запуска 
FROM alpine:3.19

WORKDIR /app

# Копируем бинарник из builder и .env из предыдущего этапа
COPY --from=builder /app/user-service .
COPY --from=builder /app/.env .

EXPOSE 8081

# Команда по умолчанию
CMD ["./user-service"]