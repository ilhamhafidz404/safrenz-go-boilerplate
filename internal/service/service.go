package service

// LAYER 2: SERVICE (Logika Bisnis & Mapping DTO/Model)

import (
	"fmt"
	"safrenz-go-boilerplate/internal/repository"
)

type HelloService interface {
	SayHello() string
}

type helloService struct {
	repo repository.HelloRepository
}

func NewHelloService(repo repository.HelloRepository) HelloService {
	return &helloService{repo: repo}
}

func (s *helloService) SayHello() string {
	msg := s.repo.GetMessage()
	return fmt.Sprintf("Service memproses: %s", msg)
}
