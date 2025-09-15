# wheather golang server
Этот проект — простой HTTP-сервер на Go, который предоставляет текущую дату, время и информацию о погоде, используя [WeatherAPI](https://www.weatherapi.com/).
## Требования
- golang 1.25
- API-ключ от WeatherAPI (бесплатно можно зарегистрировать на сайте)
## Запуск
```bash
# склонировать репозиторий
git clone <repo-url>
cd <repo-folder>

# скомпилировать бинарник
go build -o weather-service

# задать API-ключ (обязательный шаг)
export WEATHERAPI_API_KEY="<api_ключ>"

# запустить сервис
./weather-service
```
## Проверка
```bash
# Получить текущую дату
curl http://localhost:8080/date
# -> 2025-09-15

# Получить текущее время
curl http://localhost:8080/time
# -> 14:28:45

# Получить погоду
curl http://localhost:8080/weather
# -> Погода в Rostov-on-Don: 23.5°C, Облачно
```
