package main

import (
	"ClientServerCP/internal/server"
)

func main() {
	s := &server.Server{NumOfProcessed: 0, NumOfDuplicates: 0, SumTime: 0, AvgTime: 0, RequestsMemory: [][]byte{}}
	server.StartServer(s)
}
