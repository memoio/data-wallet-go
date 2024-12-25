package server

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func loadMfileDIDMoudles(r *gin.RouterGroup, h *handle) {
	r.POST("/upload/create", h.uploadCreate)
	r.POST("/upload/confirm", h.uploadConfirm)
	r.GET("/download", h.download)

}

//	@Summary		UploadCreate
//	@Description	UploadCreate
//	@Tags			mfile
//	@Accept			json
//	@Produce		json
//	@Param			data	body		string	true	"data"
//	@Param			address	body		string	true	"address"
//	@Success		200		{string}	string	"ok"
//	@Router			/mfile/upload/create [post]
func (h *handle) uploadCreate(c *gin.Context) {
	body := make(map[string]interface{})
	c.BindJSON(&body)
	data, ok := body["data"].(string)
	if !ok {
		err := fmt.Errorf("invalid data")
		h.logger.Error(err)
		c.JSON(200, err.Error())
		return
	}

	address := body["address"].(string)
	if address == "" {
		err := fmt.Errorf("invalid address")
		h.logger.Error(err)
		c.JSON(200, err.Error())
		return
	}
	fmt.Println(data)

	// TODO: upload
	c.JSON(200, "ok")
}

//	@Summary		UploadConfirm
//	@Description	UploadConfirm
//	@Tags			mfile
//	@Accept			json
//	@Produce		json
//	@Param			sig		body		string	true	"sig"
//	@Param			address	body		string	true	"address"
//	@Success		200		{string}	string	"ok"
//	@Router			/mfile/upload/confirm [post]
func (h *handle) uploadConfirm(c *gin.Context) {
	body := make(map[string]interface{})
	c.BindJSON(&body)
	sig, ok := body["sig"].(string)
	if !ok {
		err := fmt.Errorf("invalid sig")
		h.logger.Error(err)
		c.JSON(200, err.Error())
		return
	}

	address, ok := body["address"].(string)
	if !ok {
		err := ERRINVALIDPARAMS.Error()
		h.logger.Error(err)
		c.JSON(ERRINVALIDPARAMS.Code, err.Error())
		return
	}

	fmt.Println(address, sig)

	// TODO: upload
	c.JSON(200, "ok")
}

//	@Summary		Download
//	@Description	Download
//	@Tags			mfile
//	@Accept			json
//	@Produce		json
//	@Param			mdid	query		string	true	"mdid"
//	@Param			address	query		string	true	"address"
//	@Success		200		{string}	string	"ok"
//	@Router			/mfile/download [get]
func (h *handle) download(c *gin.Context) {
	mdid := c.Query("mdid")

	address := c.Query("address")

	if mdid == "" || address == "" {
		err := fmt.Errorf("invalid cid or address")
		h.logger.Error(err)
		c.JSON(200, err.Error())
		return
	}

	// TODO: download
	c.JSON(200, "ok")
}

func createCacheFile(filepath string, r io.Reader) error {
	return nil
}
