package main

import (
	"comparasion/callback"
	"comparasion/resources"
	"log"

	"net/http"
	_ "net/http/pprof"
)

func main() {
	go http.ListenAndServe(":6060", nil)

	repo := resources.NewRepository()
	service := resources.NewService(repo)

	//s := pointers.NewServer(&service)
	s := callback.NewServer(service)
	s.SetRouters()

	if err := s.Start(":8088"); err != nil {
		log.Fatalf("can't start server:%s", err)
	}
}
