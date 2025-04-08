package server

import (
	"ClientServerCP/internal/config"
	"ClientServerCP/logs"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func StartServer(s *Server, cfgServer *config.ServerConfig) {
	if s == nil {
		fmt.Println("Некорреткная переменная сервера")
		return
	}

	http.HandleFunc("/endpoint", func(resp http.ResponseWriter, req *http.Request) {
		procTime := rand.Intn(490) + 10 // Задаём время обработки

		defer req.Body.Close()

		bodyBytes, err := io.ReadAll(req.Body) // Читаем тело запроса
		if err != nil {
			logs.PrintToLogFile(cfgServer.LogStorage, "Ошибка получения тела: "+err.Error())
			http.Error(resp, "Ошибка получения тела", http.StatusBadRequest)
			return
		}

		if s.readToMemory(bodyBytes, cfgServer) { // читаем эффективно в память
			serverMetrics(s, procTime, s.checkDuplicate())

			var decodedJson []config.JSONStruct // для декодирования JSON
			errJson := json.Unmarshal(s.RequestsMemory[len(s.RequestsMemory)-1], &decodedJson)
			if errJson != nil {
				logs.PrintToLogFile(cfgServer.LogStorage, "Ошибка декодирования JSON: "+errJson.Error())
				http.Error(resp, "Ошибка декодирования JSON", http.StatusBadRequest)
				return
			}

			logs.PrintToLogFile(cfgServer.LogStorage, fmt.Sprintf("Декодированная информация из %d слота RAM: %+v", len(s.RequestsMemory)-1, decodedJson))
			fmt.Printf("[Метрики сервера]\n Число обработанных: %v\n Число дупликатов: %v\n Среднее время обработки: %v мс\n", s.NumOfProcessed, s.NumOfDuplicates, s.AvgTime)
		} else {
			logs.PrintToLogFile(cfgServer.LogStorage, "Ошибка чтения в память")
			http.Error(resp, "Ошибка чтения в память", http.StatusInternalServerError)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logs.PrintToLogFile(cfgServer.LogStorage, "Ошибка при запуске сервера: "+err.Error())
		return
	}
}
