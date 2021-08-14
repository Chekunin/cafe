package rest

import (
	"cafe/pkg/client_gateway_service/delivery/rest/schema"
	"cafe/pkg/common"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *rest) handlerEvaluatePlace(c *gin.Context) {
	var req schema.ReqEvaluatePlace
	if err := c.ShouldBindJSON(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from body"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var reqUri struct {
		PlaceID int `uri:"id" binding:"required"`
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

	placeEvaluation, marks, err := r.usecase.EvaluatePlace(c.Request.Context(), reqUri.PlaceID, userID, req.PlaceEvaluationMarks, req.Comment)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase EvaluatePlace"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	placeEvaluation.PlaceEvaluationMarks = marks

	c.JSON(http.StatusOK, placeEvaluation)
}

func (r *rest) handlerGetPlaceEvaluation(c *gin.Context) {
	var reqUri struct {
		PlaceID int `uri:"id" binding:"required"`
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

	placeEvaluation, marks, err := r.usecase.GetPlaceEvaluation(c.Request.Context(), reqUri.PlaceID, userID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetPlaceEvaluation"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	placeEvaluation.PlaceEvaluationMarks = marks

	c.JSON(http.StatusOK, placeEvaluation)
}
