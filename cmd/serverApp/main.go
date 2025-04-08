package main

import (
	"ClientServerCP/internal/config"
	"ClientServerCP/internal/server"
	"fmt"
)

func main() {
	cfgServer, err1 := config.LoadServerConfig("config/serverConfig.yaml")
	if err1 != nil {
		fmt.Println("Ошибка открытия лога в сервере", err1)
		return
	}
	fmt.Println("ServerConfig:", cfgServer)

	s := &server.Server{NumOfProcessed: 0, NumOfDuplicates: 0, SumTime: 0, AvgTime: 0, RequestsMemory: make([][]byte, cfgServer.MemorySize)}
	server.StartServer(s, cfgServer)
}
