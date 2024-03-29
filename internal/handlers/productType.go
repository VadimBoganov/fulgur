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
	param := c.Query("productid")
	var pts []domain.ProductType
	var err error
	if len(param) > 0 {
		productId, err := strconv.Atoi(param)
		if err != nil {
			logger.Errorf("Error while convert to int param from all product types request: %s", err.Error())
			return
		}
		pts, err = h.service.ProductType.GetByProductId(productId)
		if err != nil {
			logger.Errorf("Error while handle get param from all product product types request: %s", err.Error())
			return
		}
	} else {
		pts, err = h.service.ProductType.GetAll()
		if err != nil {
			logger.Errorf("Error while handle get all product types request: %s", err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, pts)
}

func (h *Handler) PostPorductType(c *gin.Context) {
	var newPt domain.ProductType

	var body, err = io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Errorf("Error while read product types request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &newPt)
	if err != nil {
		logger.Errorf("Error while deserialize product types request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	iid, err := h.service.ProductType.Add(newPt)
	if err != nil {
		logger.Errorf("Error while send product typs to db: %s", err.Error())
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
		logger.Errorf("Error while read product type request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &updatedPt)
	if err != nil {
		logger.Errorf("Error while deserialize product types request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	err = h.service.ProductType.Update(updatedPt)
	if err != nil {
		logger.Errorf("Error while update product type by id to db: %s", err.Error())
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
		logger.Errorf("Error while remove product type by id to db: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	c.JSON(http.StatusOK, iid)
}
