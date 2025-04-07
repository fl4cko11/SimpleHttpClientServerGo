package server

import (
	"ClientServerCP/internal/config"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func StartServer(s *Server) {
	if s == nil {
		return
	}

	http.HandleFunc("/endpoint", func(resp http.ResponseWriter, req *http.Request) {
		procTime := rand.Intn(490) + 10 // Задаём время обработки

		// Процесс обработки тела ---------------------------------------------------------
		defer req.Body.Close()                 // Закрываем тело запроса после чтения
		bodyBytes, err := io.ReadAll(req.Body) // Читаем тело запроса
		if err != nil {
			fmt.Println("Ошибка получения тела:", err)
			http.Error(resp, "Ошибка получения тела", http.StatusBadRequest)
			return
		}

		if s.readToMemory(bodyBytes) { // читаем эффективно в память
			serverMetrics(s, procTime, s.checkDuplicate())

			var decodedJson []config.JSONStruct // для декодирования JSON
			errJson := json.Unmarshal(s.RequestsMemory[len(s.RequestsMemory)-1], &decodedJson)
			if errJson != nil {
				fmt.Println("Ошибка декодирования JSON:", errJson)
				http.Error(resp, "Ошибка декодирования JSON", http.StatusBadRequest)
				return
			}

			fmt.Printf("Декодированная информация из %d слота RAM: ", len(s.RequestsMemory)-1)
			fmt.Printf("%+v \n", decodedJson)

			fmt.Printf("[Метрики сервера]\n Число обработанных: %v\n Число дупликатов: %v\n Среднее время обработки: %v мс\n ", s.NumOfProcessed, s.NumOfDuplicates, s.AvgTime)
		} else {
			fmt.Println("Ошибка чтения в память")
			http.Error(resp, "Ошибка чтения в память", http.StatusInternalServerError)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		return
	}
}
