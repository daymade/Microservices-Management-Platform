package main

import (
	"github.com/gin-gonic/gin"
	"catalog-service-management-api/internal/api"
	"catalog-service-management-api/internal/storage"
)

func main() {
	r := gin.Default()

	memStorage := storage.NewMemoryStorage()
	handler := api.NewHandler(memStorage)

	api.SetupRoutes(r, handler)

	r.Run(":8080")
}
