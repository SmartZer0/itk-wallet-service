# ITK Wallet Service

Простой сервис для работы с кошельками: пополнение и снятие средств через REST API.
Сервис хранит данные в PostgreSQL, поддерживает конкурентную работу и запускается в Docker.

---

## Описание

Сервис предоставляет два основных API:

- `POST /api/v1/wallet` — изменить баланс (операции `DEPOSIT` или `WITHDRAW`)
- `GET  /api/v1/wallets/{uuid}` — получить текущий баланс кошелька

Особенности реализации:
- Поддержка до 1000 RPS на один кошелёк
- Защита от гонок (используется транзакция с уровнем изоляции `SERIALIZABLE`)
- Все переменные окружения читаются из `.env`
- Поддержка запуска в `Docker` с `PostgreSQL`
- Покрытие основными юнит-тестами
- Структура проекта по принципам чистой архитектуры

---

## Используемый стек

- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [mux](https://github.com/gorilla/mux)
- [sqlmock](https://github.com/DATA-DOG/go-sqlmock)
- [testify](https://github.com/stretchr/testify)

---

## Структура проекта

```
itk/
├── cmd/                # Точка входа
├── internal/
│   ├── handler/        # HTTP-хендлеры
│   ├── models/         # Общие модели
│   ├── repository/     # Работа с базой данных
│   └── service/        # Бизнес-логика
├── migrations/         # Скрипты миграций БД
├── tests/              # Unit-тесты
├── .env                # Конфигурация
├── docker-compose.yml  # Поднятие инфраструктуры
├── Dockerfile          # Описание сборки docker-образа приложения
├── go.mod              # Модули Go и зависимости
├── go.sum              # Контрольные суммы зависимостей
├── README.md           # Документация
└── request.json        # Пример тестового запроса к API
```

---

## Как запустить

### 1. Склонируйте проект
```bash
git clone https://github.com/yourusername/itk-wallet-service.git
cd itk-wallet-service
```

### 2. Заполните `.env`
Создайте файл `.env` в корне проекта:

```
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=wallets
SERVER_PORT=8080
```

### 3. Запустите Docker
```bash
docker-compose up --build
```

Сервис будет доступен на `http://localhost:8080`.

---

## Тесты

Запуск всех тестов:
```bash
go test ./tests/...
```

---

## Примеры запросов

### POST `/api/v1/wallet`

**Тело запроса:**
```json
{
  "walletId": "afb8ad3a-338c-45d0-ace4-f93715f3c234",
  "operationType": "DEPOSIT",
  "amount": 500
}
```

### GET `/api/v1/wallets/{uuid}`

**Пример:**
```bash
curl http://localhost:8080/api/v1/wallets/afb8ad3a-338c-45d0-ace4-f93715f3c234
```
