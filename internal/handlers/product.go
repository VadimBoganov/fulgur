package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

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

	lid, err := h.service.Add(newProducts)
	if err != nil {
		logger.Errorf("Error while send products to db: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	if len(newProducts) < 2 {
		newProducts[0].Id = int(lid)
		c.JSON(http.StatusCreated, newProducts)
	} else {
		c.JSON(http.StatusCreated, newProducts)
	}
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	var updatedProduct domain.Product

	var body, err = io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Errorf("Error while read products request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &updatedProduct)
	if err != nil {
		logger.Errorf("Error while deserialize products request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	err = h.service.UpdateById(updatedProduct.Id, updatedProduct.Name)
	if err != nil {
		logger.Errorf("Error while update product by id to db: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	iid, err := strconv.Atoi(id)
	if err != nil {
		logger.Errorf("Error while convert id to int: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	err = h.service.RemoveById(iid)
	if err != nil {
		logger.Errorf("Error while remove product by id to db: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	c.JSON(http.StatusOK, iid)
}
