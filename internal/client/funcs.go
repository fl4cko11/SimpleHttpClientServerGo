package client

import (
	"math/rand"
	"time"
)

func generateSleepTime() time.Duration {
	return time.Duration(rand.Intn(999)+1) * time.Millisecond // генерируем параметр периодичности сообщений и меняем размерность на мс
}

func generateNumOfEvents() int {
	return rand.Intn(999) + 1
}
