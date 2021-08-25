package rest

import (
	"bytes"
	"cafe/pkg/client_gateway_service/delivery/rest/schema"
	"cafe/pkg/client_sso/models"
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
	var req schema.ReqAddPlaceReview
	if err := c.ShouldBindJSON(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from body"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userID, has := common.FromContextUserID(c.Request.Context())
	if !has {
		err := wrapErr.NewWrapErr(fmt.Errorf("userID is not in context"), nil)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	review, err := r.usecase.AddPlaceReview(c.Request.Context(), userID, req.Text, req.ReviewMediaIDs)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase AddPlaceReview"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.JSON(http.StatusOK, models.Convert(review, modelTag))
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

	userID, has := common.FromContextUserID(c.Request.Context())
	if !has {
		err := wrapErr.NewWrapErr(fmt.Errorf("userID is not in context"), nil)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	reviewMedia, err := r.usecase.AddPlaceReviewMedia(c.Request.Context(), userID, file)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase AddPlaceReviewMedia"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.JSON(http.StatusOK, models.Convert(reviewMedia, modelTag))
}

func (r *rest) handlerGetPlaceReviewMediaData(c *gin.Context) {
	var reqUri struct {
		ReviewMediaID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&reqUri); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	stream, contentType, err := r.usecase.GetReviewMediaData(c.Request.Context(), reqUri.ReviewMediaID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetReviewMediaData reviewMediaID=%d", reqUri.ReviewMediaID), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(stream); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("buf ReadFrom"), err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	stream.Close()
	c.Writer.Header().Set("Content-type", contentType)
	c.Data(http.StatusOK, contentType, buf.Bytes())
}

func (r *rest) handlerGetPlaceEvaluationCriterions(c *gin.Context) {
	resp, err := r.usecase.GetPlaceEvaluationCriterions(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetPlaceEvaluationCriterions"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.JSON(http.StatusOK, models.Convert(resp, modelTag))
}

func (r *rest) handlerGetOwnPlacesReviews(c *gin.Context) {
	userID, has := common.FromContextUserID(c.Request.Context())
	if !has {
		err := wrapErr.NewWrapErr(fmt.Errorf("userID is not in context"), nil)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var reqQuery struct {
		LastReviewID int `form:"last_review_id"`
		Limit        int `form:"limit,default=20" binding:"gte=0,lte=50"`
	}
	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from query"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.usecase.GetPlacesReviewsOfUserID(c.Request.Context(), userID, reqQuery.LastReviewID, reqQuery.Limit)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetPlacesReviewsOfUserID userID=%d", userID), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.JSON(http.StatusOK, models.Convert(resp, modelTag))
}

func (r *rest) handlerGetPlacesReviewsByUserID(c *gin.Context) {
	var reqUri struct {
		UserID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&reqUri); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var reqQuery struct {
		LastReviewID int `form:"last_review_id"`
		Limit        int `form:"limit,default=20" binding:"gte=0,lte=50"`
	}
	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from query"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.usecase.GetPlacesReviewsOfUserID(c.Request.Context(), reqUri.UserID, reqQuery.LastReviewID, reqQuery.Limit)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetPlacesReviewsOfUserID userID=%d", reqUri.UserID), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.JSON(http.StatusOK, models.Convert(resp, modelTag))
}

func (r *rest) handlerGetPlaceAdvertsByPlaceID(c *gin.Context) {
	var reqUri struct {
		PlaceID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&reqUri); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var reqQuery struct {
		LastUserFeedID int `form:"last_user_feed_id"`
		Limit          int `form:"limit,default=20" binding:"gte=0,lte=50"`
	}
	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from query"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.usecase.GetPlaceAdvertsByPlaceID(c.Request.Context(), reqUri.PlaceID, reqQuery.LastUserFeedID, reqQuery.Limit)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetPlaceAdvertsByPlaceID placeID=%d", reqUri.PlaceID), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.JSON(http.StatusOK, models.Convert(resp, modelTag))
}
