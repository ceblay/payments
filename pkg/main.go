package main

import (
	"github.com/ceblay/payments/pkg/ports"
	"github.com/ceblay/payments/pkg/service"
)

func main() {
	application := service.NewApplication()
	server := ports.NewHttpServer(application)
	server.Run()
}
