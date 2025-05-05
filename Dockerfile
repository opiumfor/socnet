# Используем официальный образ Golang
FROM golang:1.24.2-bookworm AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем модули Go и восстанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

WORKDIR /app/cmd
# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o ../app .

# Используем минимальный образ для запуска приложения
FROM alpine:latest

# Создаем директорию для приложения
WORKDIR /root/

# Копируем собранный бинарник из предыдущего этапа
COPY --from=builder /app/app .

# Запускаем приложение
CMD ["./app"]
