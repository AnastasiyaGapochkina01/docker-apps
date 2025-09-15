# Server Catalog (Go)
Простое REST API на Go для управления каталогом серверов. Хранение данных осуществляется в памяти без базы данных.
## Запуск
```bash
# склонировать репозиторий
git clone <repo-url>
cd <repo-folder>

# инициализировать модуль
go mod init <имя_модуля>
# установить зависимости
go mod tidy

# скомпилировать бинарник
go build -o <имя_бинаря>

# запустить сервис
./<имя_бинаря>
```
## Проверка
```bash
# добавить сервер
curl -X POST http://localhost:8080/servers \
  -H "Content-Type: application/json" \
  -d '{"name":"web-01","cpu":4,"ram":8192,"region":"eu-central"}'

# получить список серверов
curl http://localhost:8080/servers

```
