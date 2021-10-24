package main

import (
	"fmt"

	"github.com/tal-tech/go-zero/core/service"
)

type Service struct {
	name string
}

func NewService(name string) service.Service {
	return &Service{
		name: name,
	}
}

func (s *Service) Start() {
	fmt.Println("Starting Service" + s.name)
}

func (s *Service) Stop() {
	fmt.Println("Stopping Service" + s.name)
}

func main() {
	group := service.NewServiceGroup()
	defer group.Stop()

	group.Add(NewService("A"))
	group.Add(NewService("B"))
	group.Start()
}
