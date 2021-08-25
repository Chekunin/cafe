package rest

import (
	"cafe/pkg/common"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

func (r *rest) handlerGetUserSubscriptions(c *gin.Context) {
	userID, has := common.FromContextUserID(c.Request.Context())
	if !has {
		err := wrapErr.NewWrapErr(fmt.Errorf("userID is not in context"), nil)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	resp, err := r.usecase.GetUserSubscriptionsByFollowerID(c.Request.Context(), userID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetUserSubscriptionsByFollowerID"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	fmt.Printf("qqq\n%+v\n", resp)

	c.JSON(http.StatusOK, resp)
}

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

func (r *rest) handlerGetFeed(c *gin.Context) {
	userID, has := common.FromContextUserID(c.Request.Context())
	if !has {
		err := wrapErr.NewWrapErr(fmt.Errorf("userID is not in context"), nil)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err := r.usecase.GetFeedOfUserID(c.Request.Context(), userID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetFeedOfUserID userID=%d", userID), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.Status(http.StatusOK)
}
