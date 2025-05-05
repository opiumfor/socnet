### Инструкция по запуску приложения через `docker-compose`

1. **Клонируйте репозиторий**:
   Убедитесь, что у вас есть доступ к исходному коду приложения и файлы `docker-compose.yml` и `Dockerfile`.

1. **Запустите контейнеры**:
   Выполните команду для сборки и запуска всех сервисов:
   ```bash
   docker-compose up --build
   ```
    - `--build` пересобирает образы, если они уже существуют.

1. **Проверьте работу сервисов**:
    - База данных PostgreSQL будет доступна на порту `5444`.
    - API-сервер будет доступен на порту `8787`.

1. **Остановите контейнеры**:
   Чтобы остановить контейнеры, нажмите `Ctrl+C` в терминале или выполните:
   ```bash
   docker-compose down
   ```

---

### Примеры curl-запросов для проверки работы API

#### 1. **Регистрация нового пользователя**
```bash
curl -X POST http://localhost:8787/user/register \
     -H "Content-Type: application/json" \
     -d '{
           "first_name": "Имя",
           "second_name": "Фамилия",
           "birthdate": "2017-02-01",
           "biography": "Хобби, интересы и т.п.",
           "city": "Москва",
           "password": "Секретная строка"
         }'
```
- **Ответ** (идентификатор нового пользователя):
  ```json
  {"message":"User registered successfully","user_id":13}
  ```

#### 2. **Авторизация пользователя**
```bash
curl -X POST http://localhost:8787/login \
     -H "Content-Type: application/json" \
     -d '{"id": "13", "password": "Секретная строка"}'
```
- **Ответ** (JWT-токен):
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDY1NTk1MzAsInN1YiI6IjEzIn0.2rOzh0CmcE7vgWY_maS7mTojwFgZeh8HldTm9M66uzo"
  }
  ```

#### 3. **Получение анкеты пользователя**
```bash
curl -X GET http://localhost:8787/user/get/13
```
- **Ответ**:
  ```json
  {
    "id": "1",
    "first_name": "Сержант",
    "second_name": "Махоуни",
    "birthdate": "1970-01-01",
    "biography": "Не был, не состоял, не привлекался",
    "city": "х. Гуляй-Борисовка"
  }
  ```
