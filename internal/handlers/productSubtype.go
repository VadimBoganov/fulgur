package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/VadimBoganov/fulgur/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllProductSubtypes(c *gin.Context) {
	psts, err := h.service.ProductSubtype.GetAll()
	if err != nil {
		logger.Errorf("Error while handle get all products request: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, psts)
}

func (h *Handler) PostPorductSubtype(c *gin.Context) {
	var newPst domain.ProductSubType

	var body, err = io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Errorf("Error while read product sub types request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &newPst)
	if err != nil {
		logger.Errorf("Error while deserialize product sub types request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	iid, err := h.service.ProductSubtype.Add(newPst)
	if err != nil {
		logger.Errorf("Error while send product sub types to db: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	newPst.Id = int(iid)

	c.JSON(http.StatusCreated, newPst)
}

func (h *Handler) UpdateProductSubtype(c *gin.Context) {
	var updatedPst domain.ProductSubType

	var body, err = io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Errorf("Error while read product sub type request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &updatedPst)
	if err != nil {
		logger.Errorf("Error while deserialize product sub types request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	err = h.service.ProductSubtype.Update(updatedPst)
	if err != nil {
		logger.Errorf("Error while update product sub type by id to db: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	c.JSON(http.StatusOK, updatedPst)
}

func (h *Handler) DeleteProductSubype(c *gin.Context) {
	id := c.Param("id")

	iid, err := strconv.Atoi(id)
	if err != nil {
		logger.Errorf("Error while convert id to int: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	err = h.service.ProductSubtype.Remove(iid)
	if err != nil {
		logger.Errorf("Error while remove product sub type by id to db: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	c.JSON(http.StatusOK, iid)
}
