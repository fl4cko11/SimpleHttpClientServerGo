package jsonGenerator

import (
	"ClientServerCP/internal/config"
	"encoding/json"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func CreateJson(numOfEvents int) ([]byte, bool) {
	var jsonArray []config.JSONStruct

	for i := 0; i < numOfEvents; i++ {
		NewUUID := uuid.New()
		CurrentTime := time.Now()
		Status := rand.Intn(2)

		JSONData := config.JSONStruct{
			Id:     NewUUID,
			Date:   CurrentTime,
			Status: Status,
		}
		jsonArray = append(jsonArray, JSONData)
	}

	JSONResult, err := json.Marshal(jsonArray) // cериализуем массив в JSON
	if err != nil {
		return nil, false
	}
	return JSONResult, true
}
