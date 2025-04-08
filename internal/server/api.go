package server

import (
	"ClientServerCP/internal/config"
	"ClientServerCP/logs"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
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

		hashedBytes := s.hashBytes(bodyBytes) // Т.к. метрика сервера предполагает только наличие дубликатов, то не храним события целиком в памяти, а хешируем их

		s.readToMemory(hashedBytes, cfgServer) // Эффективное чтение в память
		fmt.Printf("[Состояние памяти сервера]\n Ёмкость ReqMem: %d\n Размер ReqMem: %d\n", cap(s.RequestsMemory), len(s.RequestsMemory))

		// За обработку считаем дешифрование JSON (можно дописать логику, что с ней делать, я пишу в лог просто (закомменченная строка)) ---------------------------------------
		var decodedJson []config.JSONStruct // для декодирования JSON
		errJson := json.Unmarshal(bodyBytes, &decodedJson)
		if errJson != nil {
			logs.PrintToLogFile(cfgServer.LogStorage, "Ошибка декодирования JSON: "+errJson.Error())
			http.Error(resp, "Ошибка декодирования JSON", http.StatusBadRequest)
			return
		}
		time.Sleep(time.Duration(procTime) * time.Millisecond) // даём время на обработку
		// logs.PrintToLogFile(cfgServer.LogStorage, fmt.Sprintf("Декодированная информация из %d слота RAM: %+v", len(s.RequestsMemory)-1, decodedJson))
		// --------------------------------------------------------------------------------------------------------------------------------------------

		s.serverMetrics(procTime, hashedBytes)
		fmt.Printf("[Метрики сервера]\n Число обработанных: %v\n Число дубликатов: %v\n Среднее время обработки: %v мс\n", s.NumOfProcessed, s.NumOfDuplicates, s.AvgTime)

		if len(s.RequestsMemory) >= cfgServer.MemorySize { // начинаем очистку только когда уже переполняемся, а пока держим
			s.clearProcessedEvent() // обработанное событие удаляем из памяти
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logs.PrintToLogFile(cfgServer.LogStorage, "Ошибка при запуске сервера: "+err.Error())
		return
	}
}
