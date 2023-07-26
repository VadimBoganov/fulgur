package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/VadimBoganov/fulgur/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllProductTypes(c *gin.Context) {
	pts, err := h.service.ProductType.GetAll()
	if err != nil {
		logger.Errorf("Error while handle get all products request: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, pts)
}

func (h *Handler) PostPorductType(c *gin.Context) {
	var newPt domain.ProductType

	var body, err = io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Errorf("Error while read products request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &newPt)
	if err != nil {
		logger.Errorf("Error while deserialize products request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	iid, err := h.service.ProductType.Add(newPt)
	if err != nil {
		logger.Errorf("Error while send products to db: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	newPt.Id = int(iid)

	c.JSON(http.StatusCreated, newPt)
}

func (h *Handler) UpdateProductType(c *gin.Context) {
	var updatedPt domain.ProductType

	var body, err = io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Errorf("Error while read products request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &updatedPt)
	if err != nil {
		logger.Errorf("Error while deserialize products request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	err = h.service.ProductType.Update(updatedPt)
	if err != nil {
		logger.Errorf("Error while update product by id to db: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	c.JSON(http.StatusOK, updatedPt)
}

func (h *Handler) DeleteProductType(c *gin.Context) {
	id := c.Param("id")

	iid, err := strconv.Atoi(id)
	if err != nil {
		logger.Errorf("Error while convert id to int: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	err = h.service.ProductType.Remove(iid)
	if err != nil {
		logger.Errorf("Error while remove product by id to db: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	c.JSON(http.StatusOK, iid)
}
