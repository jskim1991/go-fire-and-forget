package main

import (
	"github.com/labstack/echo/v4"
	"go-fire-and-forget/controller"
	"go-fire-and-forget/service"
	"go-fire-and-forget/worker"
)

func main() {
	e := echo.New()

	channel := make(chan string, 300)
	go worker.Work(channel)

	svc := service.NewService(channel)
	controller := controller.NewController(*svc)
	e.GET("/test", controller.Handle)

	e.Start(":8080")
}
