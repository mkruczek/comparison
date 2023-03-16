package main

import (
	"comparasion/pointers"
	"comparasion/resources"
	"log"
)

func main() {

	repo := resources.NewRepository()
	service := resources.NewService(repo)

	s := pointers.NewServer(&service)
	s.SetRouters()

	if err := s.Start(":8088"); err != nil {
		log.Fatalf("can't start server:%s", err)
	}
}
