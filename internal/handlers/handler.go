package handlers

import (
	"math"
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
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
	router.Use(corsConfig)

	router.Static("/public", "./public")

	api := router.Group("/api")
	{
		api.GET("user", h.User)
		api.POST("/register", h.Register)
		api.POST("/login", h.Login)
		api.POST("/logout", h.Logout)

		api.GET("/products", h.GetAllProducts)
		api.POST("/products", h.PostProducts)
		api.PUT("/products/:id", h.UpdateProduct)
		api.DELETE("/products/:id", h.DeleteProduct)

		api.GET("/producttypes", h.GetAllProductTypes)
		api.POST("/producttypes", h.PostPorductType)
		api.PUT("/producttypes", h.UpdateProductType)
		api.DELETE("/producttypes/:id", h.DeleteProductType)

		api.GET("/productsubtypes", h.GetAllProductSubtypes)
		api.POST("/productsubtypes", h.PostPorductSubtype)
		api.PUT("/productsubtypes", h.UpdateProductSubtype)
		api.DELETE("/productsubtypes/:id", h.DeleteProductSubype)

		api.GET("/productitems", h.GetAllProductItems)
		api.POST("/productitems", h.PostPorductItem)
		api.PUT("/productitems", h.UpdateProductItem)
		api.DELETE("/productitems/:id", h.DeleteProductItem)

		api.GET("/items", h.GetAllItems)
		api.POST("/items", h.PostItem)
		api.PUT("/items", h.UpdateItem)
		api.DELETE("/items/:id", h.DeleteItem)
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

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
