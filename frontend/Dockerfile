# Сборка
FROM node:20-alpine AS builder

# Рабочая директория
WORKDIR /app

# Копируем package.json и package-lock.json (если есть)
COPY package*.json ./

# Устанавливаем зависимости
RUN npm ci

# Копируем весь проект
COPY . .

# Собираем production-версию
RUN npm run build

# 2. Этап выполнения (runtime)
FROM nginx:alpine

# Копируем собранный билд во внутреннюю директорию nginx
COPY --from=builder /app/dist /usr/share/nginx/html

# Копируем дефолтный конфиг nginx (опционально)
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Открываем порт
EXPOSE 80

# Запускаем nginx
CMD ["nginx", "-g", "daemon off;"]
