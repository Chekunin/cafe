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
	router.GET("/place-evaluation/by-user-id/:user_id/by-place-id/:place_id", r.handlerGetPlaceEvaluationByUserIDByPlaceID)
	router.GET("/place-evaluation-marks-by-place-evaluation-id/:place_evaluation_id", r.handlerGetPlaceEvaluationMarksByPlaceEvaluationID)
	router.GET("/review-media-by-id/:id", r.handlerGetReviewMediaByID)
	router.GET("/evaluation-criterions", r.handlerGetPlaceEvaluationCriterions)
	router.GET("/reviews-by-user-id/:user_id", r.handlerGetReviewsByUserID)
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

func (r *rest) handlerGetPlaceEvaluationByUserIDByPlaceID(c *gin.Context) {
	var req struct {
		UserID  int `uri:"user_id" binding:"required"`
		PlaceID int `uri:"place_id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resp, err := r.Usecase.GetPlaceEvaluationByUserIDByPlaceID(c.Request.Context(), req.UserID, req.PlaceID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetPlaceEvaluationByUserIDByPlaceID"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetPlaceEvaluationMarksByPlaceEvaluationID(c *gin.Context) {
	var req struct {
		PlaceEvaluationID int `uri:"place_evaluation_id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resp, err := r.Usecase.GetPlaceEvaluationMarksByPlaceEvaluationID(c.Request.Context(), req.PlaceEvaluationID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetPlaceEvaluationMarksByPlaceEvaluationID"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetReviewMediaByID(c *gin.Context) {
	var req struct {
		ReviewMediaID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resp, err := r.Usecase.GetReviewMediaByID(c.Request.Context(), req.ReviewMediaID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetReviewMediaByID id=%d", req.ReviewMediaID), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetPlaceEvaluationCriterions(c *gin.Context) {
	resp, err := r.Usecase.GetPlaceEvaluationCriterions(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetPlaceEvaluationCriterions"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetReviewsByUserID(c *gin.Context) {
	var req struct {
		UserID int `uri:"user_id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resp, err := r.Usecase.GetReviewsByUserID(c.Request.Context(), req.UserID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetReviewsByUserID userID=%d", req.UserID), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}
