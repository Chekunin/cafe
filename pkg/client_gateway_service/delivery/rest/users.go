package rest

import (
	errs "cafe/pkg/client_gateway_service/errors"
	"cafe/pkg/client_sso/models"
	"cafe/pkg/common"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SubscribeToUser godoc
// @Summary Подписаться на пользователя
// @Tags Пользователи
// @Accept json
// @Param Authorization header string true "Authorization token"
// @Param id path int true "id пользователя, на которого хотим подписаться"
// @Produce json
// @Success 200 {array} models.UserSubscription
// @Router /users/{id}/subscribe [post]
func (r *rest) handlerSubscribeToUser(c *gin.Context) {
	var req struct {
		UserID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userID, has := common.FromContextUserID(c.Request.Context())
	if !has {
		err := wrapErr.NewWrapErr(fmt.Errorf("userID is not in context"), nil)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := r.usecase.SubscribeToUser(c.Request.Context(), req.UserID, userID); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase SubscribeToUser"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.Status(http.StatusOK)
}

// Get user's subscriptions godoc
// @Summary Список пользователей, на которых подписан данный пользователь
// @Tags Пользователи
// @Accept json
// @Param Authorization header string true "Authorization token"
// @Param id path int true "User ID"
// @Produce json
// @Success 200 {array} models.UserSubscription
// @Router /users/:id/subscriptions [get]
func (r *rest) handlerGetUserSubscriptions(c *gin.Context) {
	var reqUri struct {
		UserID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&reqUri); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userID, has := common.FromContextUserID(c.Request.Context())
	if !has {
		err := wrapErr.NewWrapErr(fmt.Errorf("userID is not in context"), nil)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if reqUri.UserID != userID {
		err := wrapErr.NewWrapErr(fmt.Errorf("user id from path=%d and from token=%d", reqUri.UserID, userID), errs.ErrorForbidden)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	resp, err := r.usecase.GetUserSubscriptionsByFollowerID(c.Request.Context(), userID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetUserSubscriptionsByFollowerID"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UnsubscribeFromUser godoc
// @Summary Отписаться от пользователя
// @Tags Пользователи
// @Accept json
// @Param Authorization header string true "Authorization token"
// @Param id path int true "id пользователя, от которого хотим отписаться"
// @Produce json
// @Success 200
// @Router /users/{id}/unsubscribe [post]
func (r *rest) handlerUnsubscribeFromUser(c *gin.Context) {
	var req struct {
		UserID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userID, has := common.FromContextUserID(c.Request.Context())
	if !has {
		err := wrapErr.NewWrapErr(fmt.Errorf("userID is not in context"), nil)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := r.usecase.UnsubscribeFromUser(c.Request.Context(), req.UserID, userID); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase UnsubscribeFromUser"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.Status(http.StatusOK)
}

// Feed godoc
// @Summary Лента пользователя
// @Tags Пользователи
// @Accept json
// @Param Authorization header string true "Authorization token"
// @Produce json
// @Success 200 {array} models.UserFeed
// @Router /feed [get]
func (r *rest) handlerGetUserFeed(c *gin.Context) {
	var reqQuery struct {
		LastAdvertID int `form:"last_advert_id"`
		Limit        int `form:"limit,default=7" binding:"gte=0,lte=12"`
	}
	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from query"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userID, has := common.FromContextUserID(c.Request.Context())
	if !has {
		err := wrapErr.NewWrapErr(fmt.Errorf("userID is not in context"), nil)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	resp, err := r.usecase.GetFeedOfUserID(c.Request.Context(), userID, reqQuery.LastAdvertID, reqQuery.Limit)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetUsersFeedOfUserID userID=%d", userID), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.JSON(http.StatusOK, models.Convert(resp, modelTag))
}
