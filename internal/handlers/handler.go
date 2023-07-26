package handlers

import (
	"time"

	logger2 "github.com/VadimBoganov/fulgur/pkg/logging"
	"github.com/gin-contrib/cors"

	"github.com/VadimBoganov/fulgur/internal/services"
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
	corsConfig := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
	router.Use(corsConfig)

	api := router.Group("/api")
	{
		api.GET("/products", h.GetAllProducts)
		api.POST("/products", h.PostProducts)
		api.PUT("/products/:id", h.UpdateProduct)
		api.DELETE("/products/:id", h.DeleteProduct)

		api.GET("/producttypes", h.GetAllProductTypes)
		api.POST("/producttypes", h.PostPorductType)
		api.PUT("/producttypes", h.UpdateProductType)
		api.DELETE("/producttypes/:id", h.DeleteProductType)
	}

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
