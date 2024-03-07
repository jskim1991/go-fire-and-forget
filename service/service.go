package service

import (
	"fmt"
	"github.com/google/uuid"
)

type Service struct {
	ch chan string
}

func NewService(ch chan string) *Service {
	return &Service{
		ch: ch,
	}
}

func (m *Service) DoSomething() error {
	id := uuid.NewString()
	m.ch <- id

	fmt.Println("Service: Return", id)
	return nil
}
