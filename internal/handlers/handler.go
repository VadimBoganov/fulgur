package handlers

import (
	"time"

	logger2 "github.com/VadimBoganov/fulgur/internal/logging"

	"github.com/VadimBoganov/fulgur/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var logger = logger2.GetLogger()

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")
	{
		api.GET("/products", h.GetAllProducts)
	}

	return router
}
