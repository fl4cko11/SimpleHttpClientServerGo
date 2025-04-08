package logs

import (
	"fmt"
	"os"
)

func PrintToLogFile(path string, msg string) {
	logFile, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Не удалось открыть лог файл", path)
		return
	}

	defer logFile.Close()
	logFile.WriteString(msg + "\n")
}
