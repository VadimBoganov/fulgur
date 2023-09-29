package handlers

import (
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/VadimBoganov/fulgur/internal/config"
	"github.com/VadimBoganov/fulgur/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllProductItems(c *gin.Context) {
	psts, err := h.service.ProductItem.GetAll()
	if err != nil {
		logger.Errorf("Error while handle get all product items request: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, psts)
}

func (h *Handler) PostPorductItem(c *gin.Context) {
	var newPi domain.ProductItem

	if err := c.Request.ParseMultipartForm(32 << 20); nil != err {
		logger.Errorf("Error while parse product item request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	multipartValue := c.Request.MultipartForm.Value
	newPi.Name = multipartValue["Name"][0]
	productSubTypeId, err := strconv.Atoi(multipartValue["ProductSubTypeId"][0])
	if err != nil {
		logger.Errorf("Error while read product item request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}
	newPi.ProductSubTypeId = productSubTypeId

	var headers []*multipart.FileHeader
	for _, fHeaders := range c.Request.MultipartForm.File {
		for _, hdr := range fHeaders {
			headers = append(headers, hdr)
		}
	}

	var iid int64
	if len(headers) > 0 {
		iid, err = h.service.ProductItem.Add(newPi, headers[0])

		config := config.GetConfig()
		newPi.ImageUrl = config.FtpUrl + headers[0].Filename
	} else {
		data, err := h.service.ProductItem.GetById(int(newPi.Id))
		if err != nil {
			logger.Errorf("Error while get product item from db: %s", err.Error())
			_ = c.AbortWithError(400, err)
			return
		}
		if data.ImageUrl != "" {
			newPi.ImageUrl = data.ImageUrl
		}
		iid, err = h.service.ProductItem.Add(newPi, nil)
	}

	if err != nil {
		logger.Errorf("Error while send item to db: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	newPi.Id = int(iid)

	c.JSON(http.StatusCreated, newPi)
}

func (h *Handler) UpdateProductItem(c *gin.Context) {
	var updatedPi domain.ProductItem

	if err := c.Request.ParseMultipartForm(32 << 20); nil != err {
		logger.Errorf("Error while parse product item request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	multipartValue := c.Request.MultipartForm.Value
	updatedPi.Name = multipartValue["Name"][0]

	productSubTypeId, err := strconv.Atoi(multipartValue["ProductSubTypeId"][0])
	if err != nil {
		logger.Errorf("Error while read product sub type id request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}
	updatedPi.ProductSubTypeId = productSubTypeId

	id, err := strconv.Atoi(multipartValue["Id"][0])
	if err != nil {
		logger.Errorf("Error while read product sub type id request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}
	updatedPi.Id = id

	var headers []*multipart.FileHeader
	for _, fHeaders := range c.Request.MultipartForm.File {
		for _, hdr := range fHeaders {
			headers = append(headers, hdr)
		}
	}

	if len(headers) > 0 {
		err = h.service.ProductItem.Update(updatedPi, headers[0])

		config := config.GetConfig()
		updatedPi.ImageUrl = config.FtpUrl + headers[0].Filename
	} else {
		data, err := h.service.ProductItem.GetById(int(updatedPi.Id))
		if err != nil {
			logger.Errorf("Error while get product item from db: %s", err.Error())
			_ = c.AbortWithError(400, err)
			return
		}
		if data.ImageUrl != "" {
			updatedPi.ImageUrl = data.ImageUrl
		}
		err = h.service.ProductItem.Update(updatedPi, nil)
	}

	if err != nil {
		logger.Errorf("Error while update item to db: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	c.JSON(http.StatusOK, updatedPi)
}

func (h *Handler) DeleteProductItem(c *gin.Context) {
	id := c.Param("id")

	iid, err := strconv.Atoi(id)
	if err != nil {
		logger.Errorf("Error while convert id to int: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	err = h.service.ProductItem.Remove(iid)
	if err != nil {
		logger.Errorf("Error while remove product item by id to db: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	c.JSON(http.StatusOK, iid)
}
