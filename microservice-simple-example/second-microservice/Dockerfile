
# Этап 1: Компиляция
FROM golang:1.23-alpine AS build

WORKDIR /app

# Устанавливаем необходимые зависимости
RUN apk add --no-cache gcc musl-dev

# Копируем файлы для управления зависимостями
COPY go.mod go.sum ./
RUN ls -a /app # Выводим содержимое директории /app
RUN go mod download

# Копируем остальные файлы и собираем приложение
COPY . .
RUN ls -a /app # Проверяем, что все файлы скопированы
RUN go build -o main .

# Этап 2: Минимальный образ для запуска
FROM alpine:latest

WORKDIR /app

# Копируем скомпилированное приложение из первого этапа
COPY --from=build /app/main .
COPY --from=build /app/wait-for-it.sh .

RUN chmod a+x wait-for-it.sh


# Проверяем содержимое директории
RUN ls -a /app

# Устанавливаем необходимые зависимости для запуска, если нужно
RUN apk add --no-cache libc6-compat bash
CMD ["./wait-for-it.sh", "kafka:9092", "--", "./main"]
# ENTRYPOINT ["./wait-for-it.sh", "kafka:9092", "--"]
# CMD ["./main"]
