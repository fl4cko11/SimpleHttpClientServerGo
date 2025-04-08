package server

import (
	"ClientServerCP/internal/config"
	"bytes"
)

const RAMSize = 1000

type Server struct {
	NumOfProcessed  int
	NumOfDuplicates int
	SumTime         int // Время храним в int просто домнажаем на 10^(-3), чтобы получать миллисекунды
	AvgTime         int
	RequestsMemory  [][]byte
}

func (s *Server) addNumOfProcessed() {
	s.NumOfProcessed += 1
}

func (s *Server) addNumOfDuplicates() {
	s.NumOfDuplicates += 1
}

func (s *Server) countAvgProcTime(ProcTime int) {
	s.SumTime += ProcTime
	s.AvgTime = s.SumTime / s.NumOfProcessed
}

func (s *Server) readToMemory(BodyBites []byte, cfgServer *config.ServerConfig) bool {
	if len(s.RequestsMemory) >= cfgServer.MemorySize { // Проверка переполнения
		s.RequestsMemory = s.RequestsMemory[:len(s.RequestsMemory)-1] // Удаляем самый новый элемент
	}
	s.RequestsMemory = append(s.RequestsMemory, BodyBites) // Добавляем новый запрос в RequestsMemory
	return true
}

func (s *Server) checkDuplicate() bool {
	for i := 0; i < len(s.RequestsMemory)-1; i++ {
		if bytes.Equal(s.RequestsMemory[len(s.RequestsMemory)-1], s.RequestsMemory[i]) {
			return true
		}
	}
	return false
}

func serverMetrics(s *Server, procTime int, duplicateFlag bool) {
	s.addNumOfProcessed()
	s.countAvgProcTime(procTime)
	if duplicateFlag {
		s.addNumOfDuplicates()
	}
}
