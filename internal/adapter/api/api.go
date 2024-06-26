package api

import (
	"github.com/gin-gonic/gin"
	"catalog-service-management-api/internal/adapter/api/handler"
	"catalog-service-management-api/internal/adapter/api/route"
	"catalog-service-management-api/internal/app/service"
)

type API struct {
}

func NewHTTPServer() *API {
	return &API{}
}

func (a *API) Run(addr string) error {
	engine := gin.Default()

	sh := handler.NewServiceHandler(service.NewManager())

	route.SetupRoutes(engine, sh)

	return engine.Run(addr)
}
