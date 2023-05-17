package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/VadimBoganov/fulgur/internal/domain"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllProducts(c *gin.Context) {
	products, err := h.service.GetAll()
	if err != nil {
		logger.Errorf("Error while handle get all products request: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *Handler) PostProducts(c *gin.Context) {
	var newProducts []domain.Product

	var body, err = io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Errorf("Error while read products request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &newProducts)
	if err != nil {
		logger.Errorf("Error while deserialize products request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	err = h.service.Add(newProducts)
	if err != nil {
		logger.Errorf("Error while send products to db: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, newProducts)
}
