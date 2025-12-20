# сборка
FROM golang:1.23-alpine AS builder   

WORKDIR /app 

# Копируем файлы зависимостей и скачиваем модули
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -o restaurant-service ./cmd/server/main.go

# Этап запуска 
FROM alpine:3.19

WORKDIR /app

# Копируем бинарник из builder и .env из предыдущего этапа
COPY --from=builder /app/restaurant-service .
COPY --from=builder /app/.env .

EXPOSE 8082

# Команда по умолчанию
CMD ["./restaurant-service"]

