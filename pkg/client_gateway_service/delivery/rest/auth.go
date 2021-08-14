package rest

import (
	"cafe/pkg/client_gateway_service/delivery/rest/schema"
	"cafe/pkg/common"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *rest) handlerLogin(c *gin.Context) {
	var req schema.ReqLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from body"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	tokens, err := r.usecase.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase Login"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	resp := schema.RespLogin{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}

	c.JSON(http.StatusOK, resp)
}

func (r *rest) handlerLogout(c *gin.Context) {
	token, _ := common.FromContextToken(c.Request.Context())

	if err := r.usecase.Logout(c.Request.Context(), token); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase Logout"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.Status(http.StatusOK)
}

func (r *rest) handlerRefreshToken(c *gin.Context) {
	var req schema.ReqRefreshToken
	if err := c.ShouldBindJSON(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from body"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	tokens, err := r.usecase.RefreshToken(c.Request.Context(), req.RefreshToken)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase RefreshToken"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	resp := schema.RespRefreshToken{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}

	c.JSON(http.StatusOK, resp)
}

func (r *rest) handlerSignUp(c *gin.Context) {
	var req schema.ReqSignUp
	if err := c.ShouldBindJSON(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from body"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := r.usecase.SignUp(c.Request.Context(), req.Name, req.Phone, req.Password)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase SignUp username=%s, phone=%s", req.Name, req.Phone), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (r *rest) handlerApprovePhone(c *gin.Context) {
	var req schema.ApprovePhone
	if err := c.ShouldBindJSON(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from body"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := r.usecase.ApprovePhone(c.Request.Context(), req.UserID, req.Phone, req.Code); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase ApprovePhone"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.Status(http.StatusOK)
}
