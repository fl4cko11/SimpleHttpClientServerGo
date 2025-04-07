package client

import (
	"math/rand"
	"time"
)

func generateSleepTime() time.Duration {
	sleepValue := rand.Intn(999) + 1                    // генерируем параметр периодичности сообщений
	return time.Duration(sleepValue) * time.Millisecond // меняем размерность
}

func generateNumOfEvents() int {
	return rand.Intn(999) + 1
}
