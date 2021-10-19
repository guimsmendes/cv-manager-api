package rest

import (
	"cv-manager-api/src/config/instana"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	statusOk = "Status OK"
)

type HealthCheckHandler interface {
	Ping(ctx echo.Context) error
}

type healthCheckHandler struct {
	instanaInit instana.Instana
}

func NewHealthCheckHandler(router *echo.Echo, it instana.Instana) {
	handler := &healthCheckHandler{
		instanaInit: it,
	}
	router.GET("/health", handler.Ping)
}

func (handler *healthCheckHandler) Ping(context echo.Context) error {
	handler.instanaInit.InitMetrics(context)

	return context.JSON(http.StatusOK, statusOk)
}
