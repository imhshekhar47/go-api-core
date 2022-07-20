package service

import "sync"

type SequenceService interface {
	Next() uint64
}

type DefaultSequenceService struct {
	lock     sync.Mutex
	sequence uint64
}

func NewDefaultSequenceService() *DefaultSequenceService {
	return &DefaultSequenceService{
		lock:     sync.Mutex{},
		sequence: 0,
	}
}

func (s *DefaultSequenceService) Next() uint64 {
	s.lock.Lock()
	s.sequence = s.sequence + 1
	s.lock.Unlock()

	return s.sequence
}
