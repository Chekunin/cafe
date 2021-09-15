package rest

import (
	"cafe/pkg/client_gateway_service/delivery/rest/schema"
	"cafe/pkg/common"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login godoc
// @Summary Логин
// @Tags Авторизация
// @Accept json
// @Produce json
// @Param login body schema.ReqLogin true "Login data"
// @Success 200 {object} schema.RespLogin
// @Router /auth/login [post]
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

// Logout godoc
// @Summary Деавторизация
// @Tags Авторизация
// @Accept json
// @Param Authorization header string true "Authorization token"
// @Produce json
// @Success 200
// @Router /auth/logout [post]
func (r *rest) handlerLogout(c *gin.Context) {
	token, _ := common.FromContextToken(c.Request.Context())

	if err := r.usecase.Logout(c.Request.Context(), token); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase Logout"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.Status(http.StatusOK)
}

// Refresh token godoc
// @Summary Обновление токена
// @Tags Авторизация
// @Accept json
// @Produce json
// @Param login body schema.ReqRefreshToken true "Refresh token data"
// @Success 200 {object} schema.RespRefreshToken
// @Router /auth/refresh-token [post]
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

// Signup godoc
// @Summary Регистрация пользователя
// @Tags Авторизация
// @Accept json
// @Produce json
// @Param login body schema.ReqSignUp true "Signup data"
// @Success 200 {object} models.User
// @Router /auth/signup [post]
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

// ApprovePhone godoc
// @Summary Подтверждение номера телефона
// @Tags Авторизация
// @Accept json
// @Produce json
// @Param login body schema.ApprovePhone true "ApprovePhone data"
// @Success 200
// @Router /auth/approve-phone [post]
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
