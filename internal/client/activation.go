package client

import (
	"ClientServerCP/internal/config"
	"ClientServerCP/internal/jsonGenerator"
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func StartPeriodicHttpReqs() {
	client := &http.Client{}

	config, err := config.LoadConfig("config/serverConfig.yaml") // запускается относительно директории, где вызвали исполняемый
	fmt.Printf("URL: %s", config.Url)
	if err != nil {
		fmt.Println("Ошибка чтения лога в клиенте\n", err)
		return
	}

	count := 0
	sleepTime := generateSleepTime()

	var jsonSlice [][]byte // запоминаем json запросы, чтобы на каждый 100й отправлять случайный уже отправленный

	for {
		if count == 100 {
			randIndex := rand.Intn(len(jsonSlice))                                                 // рандомный индекс в срезе
			req, err := http.NewRequest("POST", config.Url, bytes.NewBuffer(jsonSlice[randIndex])) // создаём случайный уже существующий HTTP запрос типа POST на переданный URL, содержащий JSON объект
			if err != nil {
				fmt.Println("Ошибка создания запроса:", err)
			} else {
				client.Do(req) // отправляем запрос
				fmt.Print("Отправлен дупликат !\n")
				count = 1                 // заново отсчёт
				jsonSlice = jsonSlice[:0] // очищаем срез

			}
		} else {
			jsonData, ok := jsonGenerator.CreateJson(generateNumOfEvents()) // формируем массив json событий
			jsonSlice = append(jsonSlice, jsonData)                         // запоминаем
			if ok {
				req, err := http.NewRequest("POST", config.Url, bytes.NewBuffer(jsonData))
				if err != nil {
					fmt.Println("Ошибка создания запроса:\n", err)
				} else {
					client.Do(req)
					count += 1
				}
			} else {
				fmt.Println("Ошибка создания JSON")
			}
		}
		time.Sleep(sleepTime)
	}
}
