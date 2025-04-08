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

type ServerConfig struct {
	Url        string `yaml:"url"`
	LogStorage string `yaml:"logs_storage"`
}

type ClientConfig struct {
	LogStorage string `yaml:"logs_storage"`
	Period     int    `yaml:"period"`
}

func LoadServerConfig(path string) (*ServerConfig, error) {
	var config ServerConfig
	data, err := os.ReadFile(path) // байты из файла
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

func LoadClientConfig(path string) (*ClientConfig, error) {
	var config ClientConfig
	data, err := os.ReadFile(path) // байты из файла
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
