
# Subscription Service API

REST API на Go (с использованием [Gin](https://github.com/gin-gonic/gin)) для управления подписками.  
Поддерживается документация через Swagger.

## 🚀 Возможности
- CRUD-операции для подписок.
- Получение списка подписок с фильтрацией.
- Получение сводки по подпискам.
- Swagger-документация для удобного тестирования.

## 🛠 Стек
- [Go](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin) — HTTP-фреймворк
- [Swaggo](https://github.com/swaggo/gin-swagger) — Swagger UI
- [GORM](https://gorm.io/) — ORM для работы с БД
- [godotenv](https://github.com/joho/godotenv) — загрузка переменных окружения

## 📂 Структура проекта
```

.
├── controllers/      # Обработчики API
├── docs/             # Swagger-документация
├── initializers/     # Подключение к БД, загрузка переменных окружения, миграции
├── main.go           # Точка входа
└── README.md

````

## ⚙️ Запуск

### 1. Установить зависимости
```bash
go mod tidy
````

### 2. Сгенерировать Swagger-документацию

```bash
swag init
```

### 3. Запустить сервер

```bash
go run main.go
```

По умолчанию сервер стартует на `http://localhost:8080`.

### 4. Открыть Swagger UI

```
http://localhost:8080/swagger/index.html
```

## 📌 Эндпоинты

### CRUD

* `POST   /subs` — создать подписку
* `GET    /subs` — получить список подписок (с фильтрами: `user_id`, `service_name`)
* `GET    /subs/:id` — получить подписку по ID
* `PUT    /subs/:id` — обновить подписку
* `DELETE /subs/:id` — удалить подписку

### Дополнительно

* `GET /subs/summary` — получить сводку по подпискам

## 🛠 Переменные окружения

В `.env` файле необходимо указать:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=subscriptions
```
