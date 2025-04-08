package client

import (
	"ClientServerCP/internal/config"
	"ClientServerCP/internal/jsonGenerator"
	"ClientServerCP/logs"
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func StartPeriodicHttpReqs() {
	client := &http.Client{}

	cfgServer, err := config.LoadServerConfig("config/serverConfig.yaml")
	if err != nil {
		fmt.Println("Ошибка чтения лога в клиенте\n", err)
		return
	}
	fmt.Println("ServerConfig:", cfgServer)

	cfgClient, err1 := config.LoadClientConfig("config/clientConfig.yaml")
	if err1 != nil {
		fmt.Println("Ошибка чтения лога в клиенте\n", err1)
		return
	}
	fmt.Println("ClientConfig:", cfgClient)

	count := 0
	sleepTime := generateSleepTime()

	jsonSlice := make([][]byte, cfgClient.Period) // сразу 100 для устранения лишних ресайзов

	for {
		if count == cfgClient.Period { // на каждый период повторяем
			randIndex := rand.Intn(len(jsonSlice))                                                    // рандомный индекс в срезе
			req, err := http.NewRequest("POST", cfgServer.Url, bytes.NewBuffer(jsonSlice[randIndex])) // создаём случайный уже существующий HTTP запрос типа POST на переданный URL, содержащий JSON объект
			if err != nil {
				logs.PrintToLogFile(cfgClient.LogStorage, "Ошибка создания запроса: "+err.Error())
			} else {
				client.Do(req) // отправляем запрос
				logs.PrintToLogFile(cfgClient.LogStorage, "Отправлен дупликат !")
				count = 1                 // заново отсчёт
				jsonSlice = jsonSlice[:0] // очищаем срез
			}
		} else {
			jsonData, ok := jsonGenerator.CreateJson(generateNumOfEvents()) // формируем массив json событий
			if ok {
				jsonSlice = append(jsonSlice, jsonData) // запоминаем
				req, err := http.NewRequest("POST", cfgServer.Url, bytes.NewBuffer(jsonData))
				if err != nil {
					logs.PrintToLogFile(cfgClient.LogStorage, "Ошибка создания запроса: "+err.Error())
				} else {
					client.Do(req)
					count += 1
				}
			} else {
				logs.PrintToLogFile(cfgClient.LogStorage, "Ошибка создания JSON")
			}
		}
		time.Sleep(sleepTime)
	}
}
