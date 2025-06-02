# 🌤 WeatherBotGO

Простой Telegram-бот на Go, который показывает текущую погоду в указанном городе, используя [OpenWeather API](https://openweathermap.org/api).

![Go](https://img.shields.io/badge/Go-1.18%2B-blue?logo=go)
![Telegram Bot API](https://img.shields.io/badge/Telegram-Bot-blue?logo=telegram)

## 🚀 Возможности

- 📍 Определяет город по сообщению пользователя.
- 🌡 Показывает температуру, влажность, давление и описание погоды.
- 🌐 Поддержка русского языка.
- ⚙️ Использует API от OpenWeather.

## 📦 Установка и запуск

### 1. Клонировать репозиторий

```bash
git clone https://github.com/ksndcurrsed/WeatherBotGO.git
cd WeatherBotGO
```

### 2. Установить зависимости
```bash
go mod tidy
```

### 3. Создайте .env файл в корне проекта и добавьте TELEGRAM_TOKEN и OPENWEATHER_API_KEY
```.env
TELEGRAM_TOKEN = token
OPENWEATHER_API_KEY = api_key
```

### 4. Запускайте через
```bash
go run main.go
```

# Пример использования
- Сообщение: 
Kaluga

- Вывод: 
🌤 Погода в Kaluzhskaya Oblast’:
облачно с прояснениями, 20.2°C
Ощущается как 19.7°C
Влажность: 54%
Давление: 1014 мм

## Внимание, возможны неточности, т.к. поиск городов идет по типу like
 
