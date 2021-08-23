package rest

import (
	"bytes"
	"cafe/pkg/common/utils"
	"cafe/pkg/review_media_storage_service/api/delivery/rest/schema"
	"cafe/pkg/review_media_storage_service/api/usecase"
	"encoding/gob"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/gin-gonic/gin"
	"net/http"
)

type rest struct {
	Usecase *usecase.Usecase
}

func NewRest(router *gin.RouterGroup, usecase *usecase.Usecase) *rest {
	rest := &rest{Usecase: usecase}
	rest.routes(router)
	return rest
}

func (r *rest) routes(router *gin.RouterGroup) {
	router.GET("/medias/:dir", nil)
	router.GET("/media", r.handlerGetMedia)
	router.POST("/media", r.handlerPutMedia)
	router.DELETE("/media", nil)
}

func (r *rest) handlerGetMedia(c *gin.Context) {
	var reqQuery struct {
		Path string `form:"path" binding:"required"`
	}
	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from query"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	stream, contentType, err := r.Usecase.GetStreamMedia(c.Request.Context(), reqQuery.Path)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetMedia path=%s", reqQuery.Path), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	stream.Close()
	c.Data(http.StatusOK, contentType, buf.Bytes())
}

func (r *rest) handlerPutMedia(c *gin.Context) {
	var req schema.ReqPutMedia
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.PutStreamMedia(c.Request.Context(), req.Path, bytes.NewReader(req.Data))
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase PutStreamMedia path=%s", req.Path), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}
