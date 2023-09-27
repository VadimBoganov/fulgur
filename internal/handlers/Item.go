package handlers

import (
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/VadimBoganov/fulgur/internal/config"
	"github.com/VadimBoganov/fulgur/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllItems(c *gin.Context) {
	items, err := h.service.Item.GetAll()
	if err != nil {
		logger.Errorf("Error while handle get all product items request: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) PostItem(c *gin.Context) {
	var item domain.Item

	if err := c.Request.ParseMultipartForm(32 << 20); nil != err {
		logger.Errorf("Error while parse product item request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	multipartValue := c.Request.MultipartForm.Value
	item.Name = multipartValue["Name"][0]

	productItemId, err := strconv.Atoi(multipartValue["ProductItemId"][0])
	if err != nil {
		logger.Errorf("Error while read product item id: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}
	item.ProductItemId = int32(productItemId)

	itemPrice, err := strconv.ParseFloat(multipartValue["Price"][0], 32)
	itemPrice = toFixed(itemPrice, 2)
	if err != nil {
		logger.Errorf("Error while read item price: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}
	item.Price = float64(itemPrice)

	var headers []*multipart.FileHeader
	for _, fHeaders := range c.Request.MultipartForm.File {
		for _, hdr := range fHeaders {
			headers = append(headers, hdr)
		}
	}

	var iid int64
	if len(headers) > 0 {
		iid, err = h.service.Item.Add(item, headers[0])

		config := config.GetConfig()
		item.ImageUrl = config.FtpUrl + headers[0].Filename
	} else {
		iid, err = h.service.Item.Add(item, nil)
	}

	if err != nil {
		logger.Errorf("Error while send item to db: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	item.Id = int32(iid)

	c.JSON(http.StatusCreated, item)
}

func (h *Handler) UpdateItem(c *gin.Context) {
	var item domain.Item

	if err := c.Request.ParseMultipartForm(32 << 20); nil != err {
		logger.Errorf("Error while parse product item request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	multipartValue := c.Request.MultipartForm.Value
	item.Name = multipartValue["Name"][0]

	productItemId, err := strconv.Atoi(multipartValue["ProductItemId"][0])
	if err != nil {
		logger.Errorf("Error while read product sub type id request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}
	item.ProductItemId = int32(productItemId)

	itemPrice, err := strconv.ParseFloat(multipartValue["Price"][0], 32)
	itemPrice = toFixed(itemPrice, 2)
	if err != nil {
		logger.Errorf("Error while read item price: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}
	item.Price = float64(itemPrice)

	id, err := strconv.Atoi(multipartValue["Id"][0])
	if err != nil {
		logger.Errorf("Error while read item id request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}
	item.Id = int32(id)

	var headers []*multipart.FileHeader
	for _, fHeaders := range c.Request.MultipartForm.File {
		for _, hdr := range fHeaders {
			headers = append(headers, hdr)
		}
	}

	if len(headers) > 0 {
		err = h.service.Item.Update(item, headers[0])

		config := config.GetConfig()
		item.ImageUrl = config.FtpUrl + headers[0].Filename
	} else {
		err = h.service.Item.Update(item, nil)
	}

	if err != nil {
		logger.Errorf("Error while update item to db: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) DeleteItem(c *gin.Context) {
	id := c.Param("id")

	iid, err := strconv.Atoi(id)
	if err != nil {
		logger.Errorf("Error while convert id to int: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	err = h.service.Item.Remove(iid)
	if err != nil {
		logger.Errorf("Error while remove product item by id to db: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	c.JSON(http.StatusOK, iid)
}
