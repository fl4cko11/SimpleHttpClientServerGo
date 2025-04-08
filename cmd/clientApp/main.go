package main

import (
	"ClientServerCP/internal/client"
	"ClientServerCP/internal/config"
	"fmt"
)

func main() {
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

	client.StartPeriodicHttpReqs(cfgServer, cfgClient)
}
