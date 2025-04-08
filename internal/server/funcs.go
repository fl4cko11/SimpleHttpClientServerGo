package server

import (
	"ClientServerCP/internal/config"

	"crypto/sha256"
)

type Server struct {
	NumOfProcessed   int
	NumOfDuplicates  int
	SumTime          int // Время храним в int просто домнажаем на 10^(-3), чтобы получать миллисекунды
	AvgTime          int
	RequestsMemory   []string            // Очередь обработки, не сохраняем целиком, а хешируем
	DuplicateChecker map[string]struct{} // для поиска дубликатов "struct{}", тк значения не важны
}

func (s *Server) readToMemory(HashBytes string, cfgServer *config.ServerConfig) {
	if len(s.RequestsMemory) >= cfgServer.MemorySize { // Проверка переполнения
		s.RequestsMemory = s.RequestsMemory[:len(s.RequestsMemory)-1] // Удаляем самый новый элемент
	} else {
		s.RequestsMemory = append(s.RequestsMemory, HashBytes) // Добавляем новый запрос в RequestsMemory
	}
}

func (s *Server) clearProcessedEvent() {
	s.RequestsMemory = s.RequestsMemory[1:] // удаляем по FIFO
}

func (s *Server) hashBytes(reqBytes []byte) string {
	hash := sha256.New()
	hash.Write(reqBytes)
	return string(hash.Sum(nil)[:])
}

func (s *Server) countAvgProcTime(ProcTime int) {
	s.SumTime += ProcTime
	s.AvgTime = s.SumTime / s.NumOfProcessed
}

func (s *Server) checkDuplicate(HashBytes string) {
	if len(s.DuplicateChecker) != 0 {
		_, exists := s.DuplicateChecker[HashBytes]
		if exists {
			s.NumOfDuplicates += 1
		} else {
			s.DuplicateChecker[HashBytes] = struct{}{} // добавляем и значения не важны
		}
	} else {
		s.DuplicateChecker[HashBytes] = struct{}{}
	}
}

func (s *Server) serverMetrics(procTime int, HashBytes string) {
	s.NumOfProcessed += 1
	s.countAvgProcTime(procTime)
	s.checkDuplicate(HashBytes)
}
