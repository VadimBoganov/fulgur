package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllProducts(c *gin.Context) {
	products, err := h.service.GetAll()
	if err != nil {
		logger.Fatalf("Error while handle get all products request: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}
