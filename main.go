package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	godotenv "github.com/joho/godotenv"
)

type WeatherResponse struct {
	List []struct {
		Name string `json:"name"`
		Main struct {
			Temp      float64 `json:"temp"`
			FeelsLike float64 `json:"feels_like"`
			Pressure  int     `json:"pressure"`
			Humidity  int     `json:"humidity"`
		} `json:"main"`
		Weather []struct {
			Main        string `json:"main"`
			Description string `json:"description"`
		} `json:"weather"`
	} `json:"list"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Не удалось открыть файл .env")
	}
	key := os.Getenv("OPENWEATHER_API_KEY")
	token := os.Getenv("TELEGRAM_TOKEN")

	var city string
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Println("Authorized on account", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			city = update.Message.Text
			url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/find?q=%s&type=like&units=metric&lang=ru&APPID=%s", city, key)
			resp, err := http.Get(url)
			if err != nil {
				log.Println("Ошибка запроса:", err)
				continue
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка чтения ответа"))
				continue
			}

			var weather WeatherResponse
			err = json.Unmarshal(body, &weather)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка обработки данных тела"))
				continue
			}
			if len(weather.List) == 0 {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Город не найден")
				bot.Send(msg)
				continue
			}
			w := weather.List[0]
			temp := w.Main.Temp
			feelslike := w.Main.FeelsLike
			desc := w.Weather[0].Description
			humidity := w.Main.Humidity
			pressure := w.Main.Pressure

			msgText := fmt.Sprintf("🌤 Погода в %s:\n%s, %.1f°C\nОщущается как %.1f°C\nВлажность: %d%%\nДавление: %d мм",
				w.Name, desc, temp, feelslike, humidity, pressure)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
			bot.Send(msg)
		}
	}
}
