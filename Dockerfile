# Используем базовый образ Go
FROM golang:1.23-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum (если они есть)
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь код проекта в рабочую директорию
COPY . .

# Собираем приложение
RUN GOOS=linux go build -o /main ./cmd/avitoapi-server/main.go

# Финальный образ
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /main .
EXPOSE 8080
CMD ["./main"]
