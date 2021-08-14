package rest

import (
	"cafe/pkg/common/utils"
	"cafe/pkg/nsi_service/api/usecase"
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
	router.GET("/place-by-id/:id", r.handlerGetPlaceByID)
	router.GET("/places-inside-bound", r.handlerGetPlacesInsideBound)
	router.GET("/users-subscriptions/by-follower-id/:id", r.handlerGetUserSubscriptionsByFollowerID)
}

func (r *rest) handlerGetPlaceByID(c *gin.Context) {
	var req struct {
		PlaceID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resp, err := r.Usecase.GetPlaceByID(c.Request.Context(), req.PlaceID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetPlaceByID id=%d", req.PlaceID), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetPlacesInsideBound(c *gin.Context) {
	var req struct {
		LeftLng   float64 `form:"left_lng"`
		RightLng  float64 `form:"right_lng"`
		TopLat    float64 `form:"top_lat"`
		BottomLat float64 `form:"bottom_lat"`
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from query"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.GetPlacesInsideBound(c.Request.Context(), req.LeftLng, req.RightLng, req.TopLat, req.BottomLat)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetPlacesInsideBound"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetUserSubscriptionsByFollowerID(c *gin.Context) {
	var req struct {
		FollowerID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resp, err := r.Usecase.GetUserSubscriptionsByFollowerID(c.Request.Context(), req.FollowerID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetUserSubscriptionsByFollowerID followerID=%d", req.FollowerID), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}
