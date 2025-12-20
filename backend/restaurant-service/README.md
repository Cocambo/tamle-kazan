# Restaurant Service

Микросервис для управления ресторанами, отзывами, избранным и фотографиями.

## Технологии

- Go 1.21+
- Gin (HTTP framework)
- GORM (ORM)
- PostgreSQL

## Структура проекта

```
restaurant-service/
├── cmd/
│   └── server/
│       └── main.go          # Точка входа
├── config/
│   └── config.go            # Конфигурация
├── internal/
│   ├── models/              # Модели данных
│   ├── repository/          # Слой доступа к данным
│   ├── service/             # Бизнес-логика
│   ├── handlers/            # HTTP обработчики
│   └── router/              # Маршрутизация
├── seed/                    # Seed данные
└── go.mod
```

## Архитектура

Слоистая архитектура:
- **Handlers** - HTTP слой (Gin)
- **Service** - Бизнес-логика
- **Repository** - Работа с БД (GORM)

## Конфигурация

Переменные окружения (файл `.env`):

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=restaurant_db
SERVER_PORT=8080
PHOTOS_DIR=./storage/photos
```

## API Endpoints

### Public

- `GET /restaurants` - Список ресторанов с фильтрацией
  - Query params: `search`, `cuisine`, `min_bill`, `max_bill`, `min_rating`, `limit`, `offset`
- `GET /restaurants/:id` - Карточка ресторана
- `GET /restaurants/:id/reviews` - Отзывы ресторана

### Authorized (требуется заголовок X-User-ID от API Gateway)

- `POST /restaurants/:id/reviews` - Создать отзыв
- `POST /restaurants/:id/favorite` - Добавить в избранное
- `DELETE /restaurants/:id/favorite` - Удалить из избранного

## Запуск

1. Установите зависимости:
```bash
go mod download
```

2. Запустите через Docker Compose (конфигурация должна быть уже настроена)

3. Сервис автоматически:
   - Выполнит миграции
   - Загрузит seed данные (10 ресторанов Казани)

## Модели данных

- **Restaurant**: ресторан с рейтингом и количеством отзывов
- **Review**: отзыв пользователя (unique по user_id, restaurant_id)
- **Favorite**: избранное (unique по user_id, restaurant_id)
- **Photo**: фотография ресторана (одна может быть главной)

## Особенности

- Рейтинг ресторана рассчитывается как среднее всех отзывов и обновляется при создании нового отзыва
- Фильтрация и поиск поддерживают комбинацию параметров
- Фотографии хранятся в локальной файловой системе (Docker volume)
- Seed данные загружаются автоматически при первом запуске

