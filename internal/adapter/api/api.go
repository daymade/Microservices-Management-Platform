package api

import (
	"catalog-service-management-api/internal/adapter/api/handler"
	"catalog-service-management-api/internal/adapter/api/route"
	"catalog-service-management-api/internal/app/service"

	"github.com/gin-gonic/gin"
)

type API struct {
}

func NewHTTPServer() *API {
	return &API{}
}

func (a *API) Run(addr string) error {
	engine := gin.Default()

	sh := handler.NewServiceHandler(service.NewManager())
	uh := handler.NewUserHandler()

	route.SetupRoutes(engine, sh, uh)

	return engine.Run(addr)
}
