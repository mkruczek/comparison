package main

import (
	"comparasion/common"
	echoCallback "comparasion/echoserver/callback"
	echoPointers "comparasion/echoserver/pointers"
	ginCallback "comparasion/ginserver/callback"
	ginPointers "comparasion/ginserver/pointers"
	"comparasion/resources"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go http.ListenAndServe(":6060", nil)

	go startServer(ginPointers.NewServer(), common.GinPointers, common.GinPointersPort)
	go startServer(ginCallback.NewServer(), common.GinCallback, common.GinCallbackPort)
	go startServer(echoPointers.NewServer(), common.EchoPointers, common.EchoPointersPort)
	go startServer(echoCallback.NewServer(), common.EchoCallback, common.EchoCallbackPort)

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
