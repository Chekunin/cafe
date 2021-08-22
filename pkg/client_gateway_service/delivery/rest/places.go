package rest

import (
	"bytes"
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

// данный ендпоинт будет возвращать список всех отзывов к данному заведение
// возможно он не нужен, тк пользователю будем отдавать только агрегированную инфу по оценкам рестиков
// полезно будет подчеркнуть как какой-то конкретный пользователь (или друзья) оценивают это место
func (r *rest) handlerGetPlaceEvaluations(c *gin.Context) {
	var reqUri struct {
		PlaceID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&reqUri); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var reqQuery struct {
		OfSubscribes     bool `form:"of_subscribes"`
		LastEvaluationID int  `form:"last_evaluation_id"`
	}
	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from query"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//userID, has := common.FromContextUserID(c.Request.Context())
	//if !has {
	//	err := wrapErr.NewWrapErr(fmt.Errorf("userID is not in context"), nil)
	//	c.AbortWithError(http.StatusInternalServerError, err)
	//	return
	//}

	// здесь должны отдавать отзывы всех пользователей к данному заведение (с пагинацией)
	// или только друзей
	panic("implement me") // todo: под вопросом, нужен ли вообще метод
}

func (r *rest) handlerAddPlaceReview(c *gin.Context) {
	panic("implement me")
}

func (r *rest) handlerAddPlaceReviewMedia(c *gin.Context) {
	multipartForm, err := c.MultipartForm()
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("call c.MultipartForm()"), err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	fileHeaders, has := multipartForm.File["media"]
	if !has {
		err := wrapErr.NewWrapErr(fmt.Errorf("prices files are not attached"), nil)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// todo: исправить
	if len(fileHeaders) > 1 {
		err = wrapErr.NewWrapErr(fmt.Errorf("extracting file"), nil)
		c.AbortWithError(http.StatusBadRequest, err)
	}

	file, err := fileHeaders[0].Open()
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("open file=%s", fileHeaders[0].Filename), err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer file.Close()

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(file); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("read to buffer from filename=%s", fileHeaders[0].Filename), err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
