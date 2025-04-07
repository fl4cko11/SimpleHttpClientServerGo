package config

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"gopkg.in/yaml.v3"
)

type JSONStruct struct {
	Id     uuid.UUID `json:"id"`
	Date   time.Time `json:"date"`
	Status int       `json:"status"`
}

type YAMLStruct struct {
	Url string `yaml:"url"`
}

func LoadConfig(filename string) (*YAMLStruct, error) {
	var config YAMLStruct
	data, err := os.ReadFile(filename) // байты из файла
	if err != nil {
		fmt.Println("Ошибка считывания конфига", err)
		return nil, err
	}

	err = yaml.Unmarshal(data, &config) // байты в соответсвющую структуру
	if err != nil {
		return nil, err
	}

	return &config, nil
}
