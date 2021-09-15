package rest

import (
	"bytes"
	"cafe/pkg/client_gateway_service/delivery/rest/schema"
	errs "cafe/pkg/client_gateway_service/errors"
	"cafe/pkg/client_sso/models"
	"cafe/pkg/common"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/gin-gonic/gin"
	"net/http"
)

// EvaluatePlace godoc
// @Summary Оценить заведение
// @Tags Заведения
// @Accept json
// @Param evaluatePlace body schema.ReqEvaluatePlace true "evaluation place data"
// @Param Authorization header string true "Authorization token"
// @Param id path int true "id заведения"
// @Produce json
// @Success 200 {object} models.PlaceEvaluation
// @Router /places/{id}/evaluation [post]
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

// GetPlaceEvaluation godoc
// @Summary Достать оценку заведения от данного пользователя
// @Tags Заведения
// @Accept json
// @Param Authorization header string true "Authorization token"
// @Param id path int true "id заведения"
// @Produce json
// @Success 200 {object} models.PlaceEvaluation
// @Router /places/{id}/evaluation [get]
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

// AddPlaceReview godoc
// @Summary Загрузить медиа-файл для отзыва заведения
// @Tags Заведения
// @Accept json
// @Param id path int true "place id"
// @Param reviewMedia body schema.ReqAddPlaceReview true "place review data"
// @Param Authorization header string true "Authorization token"
// @Produce json
// @Success 200 {object} models.ReviewMedia
// @Router /places/{id}/review [post]
func (r *rest) handlerAddPlaceReview(c *gin.Context) {
	var req schema.ReqAddPlaceReview
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

	review, err := r.usecase.AddPlaceReview(c.Request.Context(), userID, reqUri.PlaceID, req.Text, req.ReviewMediaIDs)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase AddPlaceReview"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.JSON(http.StatusOK, models.Convert(review, modelTag))
}

// AddPlaceReviewMedia godoc
// @Summary Загрузить медиа-файл для отзыва заведения
// @Tags Заведения
// @Accept multipart/form-data
// @Param media formData file true "media file"
// @Param Authorization header string true "Authorization token"
// @Produce json
// @Success 200 {object} models.ReviewMedia
// @Router /place-review-media [post]
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

// PlaceReviewMediaData godoc
// @Summary Медиа-файл отзыва
// @Tags Заведения
// @Produce json
// @Param id path int true "review media id"
// @Success 200 {array} byte
// @Router /place-review-medias/{id}/data [get]
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

// Places godoc
// @Summary Критерии оценивания заведения
// @Tags Заведения
// @Produce json
// @Success 200 {array} models.EvaluationCriterion
// @Router /place-evaluation-criterions [get]
func (r *rest) handlerGetPlaceEvaluationCriterions(c *gin.Context) {
	resp, err := r.usecase.GetPlaceEvaluationCriterions(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetPlaceEvaluationCriterions"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.JSON(http.StatusOK, models.Convert(resp, modelTag))
}

// GetOwnPlacesReviews godoc
// @Summary Список отзывов данного пользователя
// @Tags Пользователи
// @Accept json
// @Param Authorization header string true "Authorization token"
// @Param id path int true "User ID"
// @Produce json
// @Success 200 {array} models.Review
// @Router /users/{id}/place-reviews [get]
func (r *rest) handlerGetOwnPlacesReviews(c *gin.Context) {
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

	var reqQuery struct {
		LastReviewID int `form:"last_review_id"`
		Limit        int `form:"limit,default=20" binding:"gte=0,lte=50"`
	}
	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from query"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if reqUri.UserID != userID {
		err := wrapErr.NewWrapErr(fmt.Errorf("user id from path=%d and from token=%d", reqUri.UserID, userID), errs.ErrorForbidden)
		c.AbortWithError(GetHttpCode(err), err)
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

// GetPlacesReviewsByUserID godoc
// @Summary Отзывы конкретного пользователя
// @Tags Пользователи
// @Produce json
// @Param id path int true "id пользователя"
// @Param last_review_id query int false "id последнего полученного отзыва данного пользователя"
// @Param limit query int false "лимит отдаваемых записей" minimum(1) maximum(50) default(20)
// @Success 200 {array} models.Review
// @Router /users/{id}/posts [get]
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

// GetPlaceAdvertsByPlaceID godoc
// @Summary Объявления заведения
// @Tags Заведения
// @Produce json
// @Param id path int true "id заведения"
// @Param last_review_id query int false "id последнего полученного объявления данного заведения"
// @Param limit query int false "лимит отдаваемых записей" minimum(1) maximum(50) default(20)
// @Success 200 {array} models.Advert
// @Router /places/{id}/posts [get]
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
		LastAdvertID int `form:"last_advert_id"`
		Limit        int `form:"limit,default=20" binding:"gte=0,lte=50"`
	}
	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from query"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.usecase.GetPlaceAdvertsByPlaceID(c.Request.Context(), reqUri.PlaceID, reqQuery.LastAdvertID, reqQuery.Limit)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetPlaceAdvertsByPlaceID placeID=%d", reqUri.PlaceID), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.JSON(http.StatusOK, models.Convert(resp, modelTag))
}

// GetPlaceSubscriptions godoc
// @Summary Список заведений, на которые подписан данный пользователь
// @Tags Пользователи
// @Accept json
// @Param Authorization header string true "Authorization token"
// @Param id path int true "User ID"
// @Produce json
// @Success 200 {array} models.UserPlaceSubscription
// @Router /users/{id}/place-subscriptions [get]
func (r *rest) handlerGetPlaceSubscriptions(c *gin.Context) {
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

	resp, err := r.usecase.GetPlaceSubscriptionsByUserID(c.Request.Context(), userID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetPlaceSubscriptions"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// SubscribeToPlace godoc
// @Summary Подписаться на заведение
// @Tags Пользователи
// @Accept json
// @Param Authorization header string true "Authorization token"
// @Param id path int true "id заведения"
// @Produce json
// @Success 200
// @Router /places/{id}/subscribe [post]
func (r *rest) handlerSubscribeToPlace(c *gin.Context) {
	var req struct {
		PlaceID int `uri:"id" binding:"required"`
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

	if err := r.usecase.SubscribeToPlace(c.Request.Context(), userID, req.PlaceID); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase SubscribeToPlace"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.Status(http.StatusOK)
}

// UnsubscribeFromPlace godoc
// @Summary Отписаться от заведения
// @Tags Заведения
// @Accept json
// @Param Authorization header string true "Authorization token"
// @Param id path int true "id заведения"
// @Produce json
// @Success 200
// @Router /places/{id}/unsubscribe [post]
func (r *rest) handlerUnsubscribeFromPlace(c *gin.Context) {
	var req struct {
		PlaceID int `uri:"id" binding:"required"`
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

	if err := r.usecase.UnsubscribeFromPlace(c.Request.Context(), userID, req.PlaceID); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase UnsubscribeFromUser"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.Status(http.StatusOK)
}

// GetPlaceMenu godoc
// @Summary Меню заведения
// @Tags Заведения
// @Produce json
// @Param id path int true "id заведения"
// @Success 200 {array} models.PlaceMenu
// @Router /places/{id}/menu [get]
func (r *rest) handlerGetPlaceMenu(c *gin.Context) {
	var req struct {
		PlaceID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.usecase.GetPlaceMenu(c.Request.Context(), req.PlaceID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetPlaceMenu placeID=%d", req.PlaceID), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.JSON(http.StatusOK, models.Convert(resp, modelTag))
}
