package controller

import (
	"github.com/labstack/echo/v4"
	"go-fire-and-forget/service"
	"net/http"
	"time"
)

type Controller struct {
	Service service.Service
}

func NewController(svc service.Service) *Controller {
	return &Controller{
		Service: svc,
	}
}

func (m *Controller) Handle(c echo.Context) error {
	now := time.Now()
	m.Service.DoSomething()

	return c.String(http.StatusOK, time.Since(now).String())
}
