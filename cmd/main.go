package main

import (
	echoCallback "comparasion/echoserver/callback"
	echoPointers "comparasion/echoserver/pointers"
	ginCallback "comparasion/ginserver/callback"
	ginPointers "comparasion/ginserver/pointers"
	"comparasion/resources"
	"comparasion/value"
	"comparasion/worker"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go http.ListenAndServe(":6060", nil)

	go startServer(ginPointers.NewServer(), value.GinPointers, value.GinPointersPort)
	go startServer(ginCallback.NewServer(), value.GinCallback, value.GinCallbackPort)
	go startServer(echoPointers.NewServer(), value.EchoPointers, value.EchoPointersPort)
	go startServer(echoCallback.NewServer(), value.EchoCallback, value.EchoCallbackPort)

	select {}
}

type HttpServer interface {
	SetService(service resources.Service)
	SetRouters(version string)
	Start(port string) error
}

func startServer(hs HttpServer, version string, port string) {
	repo := resources.NewRepository()
	service := resources.NewService(repo)

	hs.SetService(service)
	hs.SetRouters(version)
	if err := hs.Start(port); err != nil {
		log.Fatalf("can't start %s server:%s", version, err)
	}
}

func mainGinPointer() {
	go startServer(ginPointers.NewServer(), value.GinPointers, value.GinPointersPort)
	go worker.Do(value.GinPointers, value.GinPointersPort)
	time.Sleep(10 * time.Second)
}
func mainGinCallback() {
	go startServer(ginCallback.NewServer(), value.GinCallback, value.GinCallbackPort)
	go worker.Do(value.GinCallback, value.GinCallbackPort)
	time.Sleep(10 * time.Second)
}

func mainEchoPointer() {
	go startServer(echoPointers.NewServer(), value.EchoPointers, value.EchoPointersPort)
	go worker.Do(value.EchoPointers, value.EchoPointersPort)
	time.Sleep(10 * time.Second)
}

func mainEchoCallback() {
	go startServer(echoPointers.NewServer(), value.EchoPointers, value.EchoPointersPort)
	go worker.Do(value.EchoPointers, value.EchoPointersPort)
	time.Sleep(10 * time.Second)
}
