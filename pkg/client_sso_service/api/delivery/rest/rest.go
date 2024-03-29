package rest

import (
	"cafe/pkg/client_sso/models"
	"cafe/pkg/client_sso_service/api/usecase"
	"cafe/pkg/common/utils"
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
	router.POST("/login", r.handlerLogin)
	router.POST("/logout", r.handlerLogout)
	router.POST("/refresh-token", r.handlerRefreshToken)
	router.POST("/check-permission", r.handlerCheckPermission)
}

func (r rest) handlerLogin(c *gin.Context) {
	var req models.ReqLogin
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.Login(c.Request.Context(), req.UserName, req.Password)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase login username=%s, password=%s", req.UserName, req.Password), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r rest) handlerLogout(c *gin.Context) {
	var req models.ReqLogout
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := r.Usecase.Logout(c.Request.Context(), req.Token)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase Logout token=%s", req.Token), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Status(http.StatusOK)
}

func (r rest) handlerRefreshToken(c *gin.Context) {
	var req models.ReqRefreshToken
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.RefreshToken(c.Request.Context(), req.RefreshToken)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase RefreshToken refreshToken=%s", req.RefreshToken), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r rest) handlerCheckPermission(c *gin.Context) {
	var req models.ReqCheckPermission
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.CheckPermission(c.Request.Context(), req.Method, req.Path, req.Token)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase CheckPermission method=%s, path=%s, token=%s", req.Method, req.Path, req.Token), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}
