package rest

import (
	"cafe/pkg/common/utils"
	"cafe/pkg/db_manager_service/api/delivery/rest/schema"
	"cafe/pkg/db_manager_service/api/usecase"
	"cafe/pkg/models"
	"encoding/gob"
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
	router.GET("/places", r.handlerGetAllPlaces)
	router.GET("/place-medias", r.handlerGetAllPlaceMedias)
	router.GET("/categories", r.handlerGetAllCategories)
	router.GET("/place-categories", r.handlerGetAllPlaceCategories)
	router.GET("/kitchen-categories", r.handlerGetAllKitchenCategories)
	router.GET("/place-kitchen-categories", r.handlerGetAllPlaceKitchenCategories)
	router.GET("/place-schedules", r.handlerGetAllPlaceSchedules)
	router.GET("/adverts", r.handlerGetAllAdverts)
	router.GET("/advert-medias", r.handlerGetAllAdvertMedias)
	router.GET("/adverts-by-place-id/:place_id", r.handlerGetAdvertsByPlaceID)
	router.GET("/advert-by-id/:advert_id", r.handlerGetAdvertByID)
	router.GET("/users-places-by-place-id/:place_id", r.handlerGetUsersPlacesByPlaceID)
	router.GET("/users-places-by-user-id/:user_id", r.handlerGetUsersPlacesByUserID)
	router.GET("/evaluation-criterions", r.handlerGetAllEvaluationCriterions)
	router.GET("/place-evaluations", r.handlerGetAllPlaceEvaluations)
	router.POST("/place-evaluation-with-marks", r.handlerAddPlaceEvaluationWithMarks)
	router.GET("/place-evaluation-marks", r.handlerGetAllPlaceEvaluationMarks)
	router.GET("/reviews", r.handlerGetAllReviews)
	router.GET("/reviews-by-user-id/:user_id", r.handlerGetReviewsByUserID)
	router.GET("/review-by-id/:review_id", r.handlerGetGetReviewByID)
	router.POST("/review", r.handlerAddReview)
	router.POST("/review-media", r.handlerAddReviewMedia)
	router.GET("/review-medias", r.handlerGetAllReviewMedias)
	router.GET("/review-review-medias", r.handlerGetAllReviewReviewMedias)
	router.POST("/review-review-medias", r.handlerAddReviewReviewMedias)
	router.GET("/users", r.handlerGetAllUsers)
	router.GET("/user-by-id/:user_id", r.handlerGetUserByID)
	router.GET("/user-by-name/:name", r.handlerGetUserByName)
	router.GET("/user-by-phone/:phone", r.handlerGetUserByVerifiedPhone)
	router.POST("/user", r.handlerCreateUser)
	router.POST("/user/:id", r.handlerUpdateUser)

	router.GET("/user-subscriptions", r.handlerGetAllUserSubscriptions)
	router.POST("/user-subscription", r.handlerAddUserSubscription)
	router.DELETE("/user-subscription", r.handlerDeleteUserSubscription)
	router.GET("/users-subscriptions-by-followed-user-id/:followed_user_id", r.handlerGetUserSubscriptionsByFollowedUserID)

	router.GET("/place-subscriptions", r.handlerGetAllPlaceSubscriptions)
	router.POST("/place-subscription", r.handlerAddPlaceSubscription)
	router.DELETE("/place-subscription", r.handlerDeletePlaceSubscription)

	router.GET("/actual-user-phone-code-by-user-id/:user_id", r.handlerGetActualUserPhoneCodeByUserID)
	router.POST("/user-phone-code", r.handlerCreateUserPhoneCode)
	router.POST("/user-phone-code/:id", r.handlerUpdateUserPhoneCode)
	router.POST("/activate-user-phone", r.handlerActivateUserPhone)
	router.GET("/feed-by-user-id/:user_id", r.handlerGetFeedOfUserID)
	router.POST("/users-feeds", r.handlerAddUsersFeed)
	router.DELETE("/users-feeds", r.handlerDeleteUsersFeed)

	router.POST("/add-feed-advert-queue", r.handlerAddFeedAdvertQueue)
	router.POST("/poll-feed-advert-queue", r.handlerPollFeedAdvertQueue)
	router.POST("/complete-feed-advert-queue/:advert_id", r.handlerCompleteFeedAdvertQueue)

	router.POST("/add-feed-review-queue", r.handlerAddFeedReviewQueue)
	router.POST("/poll-feed-review-queue", r.handlerPollFeedReviewQueue)
	router.POST("/complete-feed-review-queue/:advert_id", r.handlerCompleteFeedReviewQueue)

	router.POST("/add-feed-user-subscribe-queue", r.handlerAddFeedUserSubscribeQueue)
	router.POST("/poll-feed-user-subscribe-queue", r.handlerPollFeedUserSubscribeQueue)
	router.POST("/complete-feed-user-subscribe-queue/:follower_user_id/:followed_user_id", r.handlerCompleteFeedUserSubscribeQueue)

	router.GET("/places/:place_id/menu", r.handlerGetPlaceMenu)
}

func (r *rest) handlerGetAllPlaces(c *gin.Context) {
	resp, err := r.Usecase.GetAllPlaces(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAllPlaces"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetAllPlaceMedias(c *gin.Context) {
	resp, err := r.Usecase.GetAllPlaceMedias(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAllPlaceMedias"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetAllCategories(c *gin.Context) {
	resp, err := r.Usecase.GetAllCategories(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAllCategories"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetAllPlaceCategories(c *gin.Context) {
	resp, err := r.Usecase.GetAllPlaceCategories(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAllPlaceCategories"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetAllKitchenCategories(c *gin.Context) {
	resp, err := r.Usecase.GetAllKitchenCategories(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAllKitchenCategories"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetAllPlaceKitchenCategories(c *gin.Context) {
	resp, err := r.Usecase.GetAllPlaceKitchenCategories(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAllPlaceKitchenCategories"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetAllPlaceSchedules(c *gin.Context) {
	resp, err := r.Usecase.GetAllPlaceSchedules(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAllPlaceSchedules"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetAllAdverts(c *gin.Context) {
	resp, err := r.Usecase.GetAllAdverts(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAllAdverts"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetAllAdvertMedias(c *gin.Context) {
	resp, err := r.Usecase.GetAllAdvertMedias(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAllAdvertMedias"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetAdvertsByPlaceID(c *gin.Context) {
	var req struct {
		PlaceID int `uri:"place_id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var reqQuery struct {
		LastAdvertID int `form:"last_adver_id"`
		Limit        int `form:"limit,default=20"`
	}
	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from query"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.GetAdvertsByPlaceID(c.Request.Context(), req.PlaceID, reqQuery.LastAdvertID, reqQuery.Limit)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAdvertsByPlaceID"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetAdvertByID(c *gin.Context) {
	var req struct {
		AdvertID int `uri:"advert_id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.GetAdvertByID(c.Request.Context(), req.AdvertID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAdvertByID"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetUsersPlacesByPlaceID(c *gin.Context) {
	var req struct {
		PlaceID int `uri:"place_id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.GetUserPlaceSubscriptionsByPlaceID(c.Request.Context(), req.PlaceID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetUserPlaceSubscriptionsByPlaceID"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetUsersPlacesByUserID(c *gin.Context) {
	var req struct {
		UserID int `uri:"user_id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.GetUsersPlacesByUserID(c.Request.Context(), req.UserID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetUsersPlacesByUserID"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetAllEvaluationCriterions(c *gin.Context) {
	resp, err := r.Usecase.GetAllEvaluationCriterions(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAllEvaluationCriterions"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetAllPlaceEvaluations(c *gin.Context) {
	resp, err := r.Usecase.GetAllPlaceEvaluations(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAllPlaceEvaluations"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerAddPlaceEvaluationWithMarks(c *gin.Context) {
	var req schema.ReqEvaluatePlace
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := r.Usecase.AddPlaceEvaluationWithMarks(c.Request.Context(), &req.PlaceEvaluation, req.Marks); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase AddPlaceEvaluationWithMarks"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(req))
}

func (r *rest) handlerGetAllPlaceEvaluationMarks(c *gin.Context) {
	resp, err := r.Usecase.GetAllPlaceEvaluationMarks(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAllPlaceEvaluationMarks"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetAllReviews(c *gin.Context) {
	resp, err := r.Usecase.GetAllReviews(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAllReviews"), err)
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

	var reqQuery struct {
		LastReviewID int `form:"last_review_id"`
		Limit        int `form:"limit,default=20"`
	}
	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from query"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.GetReviewsByUserID(c.Request.Context(), req.UserID, reqQuery.LastReviewID, reqQuery.Limit)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetReviewsByUserID"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetGetReviewByID(c *gin.Context) {
	var req struct {
		ReviewID int `uri:"review_id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.GetReviewByID(c.Request.Context(), req.ReviewID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetReviewByID reviewID=%d", req.ReviewID), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerAddReview(c *gin.Context) {
	var req models.Review
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	review, err := r.Usecase.AddReview(c.Request.Context(), req)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase AddReview"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(review))
}

func (r *rest) handlerAddReviewMedia(c *gin.Context) {
	var req models.ReviewMedia
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.AddReviewMedia(c.Request.Context(), req)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase AddReviewMedia"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetAllReviewMedias(c *gin.Context) {
	resp, err := r.Usecase.GetAllReviewMedias(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAllReviewMedias"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetAllReviewReviewMedias(c *gin.Context) {
	resp, err := r.Usecase.GetAllReviewReviewMedias(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAllReviewReviewMedias"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerAddReviewReviewMedias(c *gin.Context) {
	var req []models.ReviewReviewMedias
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := r.Usecase.AddReviewReviewMedias(c.Request.Context(), req)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase AddReviewReviewMedias"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(req))
}

func (r *rest) handlerGetAllUsers(c *gin.Context) {
	resp, err := r.Usecase.GetAllUsers(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAllUsers"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerCreateUser(c *gin.Context) {
	var req models.User
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respUser, err := r.Usecase.CreateUser(c.Request.Context(), req)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase CreateUser"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(respUser))
}

func (r *rest) handlerUpdateUser(c *gin.Context) {
	var reqUri struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&reqUri); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var req models.User
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	req.ID = reqUri.ID

	respUser, err := r.Usecase.UpdateUser(c.Request.Context(), req)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase UpdateUser"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(respUser))
}

func (r *rest) handlerGetUserByID(c *gin.Context) {
	var req struct {
		UserID int `uri:"user_id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.GetUserByID(c.Request.Context(), req.UserID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetUserByID"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetUserByName(c *gin.Context) {
	var req struct {
		UserName string `uri:"name" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.GetUserByName(c.Request.Context(), req.UserName)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetUserByName"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetUserByVerifiedPhone(c *gin.Context) {
	var req struct {
		Phone string `uri:"phone" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.GetUserByVerifiedPhone(c.Request.Context(), req.Phone)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetUserByVerifiedPhone"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetAllUserSubscriptions(c *gin.Context) {
	resp, err := r.Usecase.GetAllUserSubscriptions(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAllUserSubscriptions"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerAddUserSubscription(c *gin.Context) {
	var req models.UserSubscription
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.AddUserSubscription(c.Request.Context(), req)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase AddUserSubscription"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerDeleteUserSubscription(c *gin.Context) {
	var req models.UserSubscription
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := r.Usecase.DeleteUserSubscription(c.Request.Context(), req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase DeleteUserSubscription"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Status(http.StatusOK)
}

func (r *rest) handlerGetAllPlaceSubscriptions(c *gin.Context) {
	resp, err := r.Usecase.GetAllPlaceSubscriptions(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAllPlaceSubscriptions"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerGetUserSubscriptionsByFollowedUserID(c *gin.Context) {
	var req struct {
		FollowedUserID int `uri:"followed_user_id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.GetUserSubscriptionsByFollowedUserID(c.Request.Context(), req.FollowedUserID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetUserSubscriptionsByFollowedUserID followedUserID=%d", req.FollowedUserID), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerAddPlaceSubscription(c *gin.Context) {
	var req models.UserPlaceSubscription
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.AddPlaceSubscription(c.Request.Context(), req)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase AddUserSubscription"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerDeletePlaceSubscription(c *gin.Context) {
	var req models.UserPlaceSubscription
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := r.Usecase.DeletePlaceSubscription(c.Request.Context(), req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase DeletePlaceSubscription"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Status(http.StatusOK)
}

func (r *rest) handlerGetActualUserPhoneCodeByUserID(c *gin.Context) {
	var req struct {
		UserID int `uri:"user_id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resp, err := r.Usecase.GetActualUserPhoneCodeByUserID(c.Request.Context(), req.UserID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetActualUserPhoneCodeByUserID"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerUpdateUserPhoneCode(c *gin.Context) {
	var reqUri struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&reqUri); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var req models.UserPhoneCode
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	req.ID = reqUri.ID

	userPhoneCode, err := r.Usecase.UpdateUserPhoneCode(c.Request.Context(), req)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase UpdateUserPhoneCode"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(userPhoneCode))
}

func (r *rest) handlerCreateUserPhoneCode(c *gin.Context) {
	var req models.UserPhoneCode
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	userPhoneCode, err := r.Usecase.CreateUserPhoneCode(c.Request.Context(), req)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase CreateUserPhoneCode"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(userPhoneCode))
}

func (r *rest) handlerActivateUserPhone(c *gin.Context) {
	var req schema.ReqActivateUserPhone
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := r.Usecase.ActivateUserPhone(c.Request.Context(), req.UserPhoneCodeID); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase ActivateUserPhone"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Status(http.StatusOK)
}

func (r *rest) handlerGetFeedOfUserID(c *gin.Context) {
	var req struct {
		UserID int `uri:"user_id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var reqQuery struct {
		LastUserFeedID int `form:"last_user_feed_id"`
		Limit          int `form:"limit,default=8"`
	}
	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from query"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.GetFeedOfUserID(c.Request.Context(), req.UserID, reqQuery.LastUserFeedID, reqQuery.Limit)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetUsersFeedOfUserID"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerAddUsersFeed(c *gin.Context) {
	var req []models.UserFeed
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := r.Usecase.AddUsersFeed(c.Request.Context(), req)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetUsersFeedOfUserID"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(req))
}

func (r *rest) handlerDeleteUsersFeed(c *gin.Context) {
	var req models.UserFeed
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := r.Usecase.DeleteUsersFeed(c.Request.Context(), req)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase DeleteUsersFeed"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(req))
}

func (r *rest) handlerAddFeedAdvertQueue(c *gin.Context) {
	var req models.FeedAdvertQueue
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.AddFeedAdvertQueue(c.Request.Context(), req)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase AddFeedAdvertQueue"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerPollFeedAdvertQueue(c *gin.Context) {
	resp, err := r.Usecase.PollFeedAdvertQueue(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase PollFeedAdvertQueue"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerCompleteFeedAdvertQueue(c *gin.Context) {
	var req struct {
		AdvertID int `uri:"advert_id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := r.Usecase.CompleteFeedAdvertQueue(c.Request.Context(), req.AdvertID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase CompleteFeedAdvertQueue advertID=%d", req.AdvertID), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Status(http.StatusOK)
}

func (r *rest) handlerAddFeedReviewQueue(c *gin.Context) {
	var req models.FeedReviewQueue
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.AddFeedReviewQueue(c.Request.Context(), req)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase AddFeedReviewQueue"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerPollFeedReviewQueue(c *gin.Context) {
	resp, err := r.Usecase.PollFeedReviewQueue(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase PollFeedReviewQueue"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerCompleteFeedReviewQueue(c *gin.Context) {
	var req struct {
		ReviewID int `uri:"review_id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := r.Usecase.CompleteFeedReviewQueue(c.Request.Context(), req.ReviewID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase CompleteFeedReviewQueue reviewID=%d", req.ReviewID), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Status(http.StatusOK)
}

func (r *rest) handlerAddFeedUserSubscribeQueue(c *gin.Context) {
	var req models.FeedUserSubscribeQueue
	dec := gob.NewDecoder(c.Request.Body)
	if err := dec.Decode(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("decode data"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := r.Usecase.AddFeedUserSubscribeQueue(c.Request.Context(), req)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase AddFeedUserSubscribeQueue"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerPollFeedUserSubscribeQueue(c *gin.Context) {
	resp, err := r.Usecase.PollFeedUserSubscribeQueue(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase PollFeedUserSubscribeQueue"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}

func (r *rest) handlerCompleteFeedUserSubscribeQueue(c *gin.Context) {
	var req struct {
		FollowerUserID int `uri:"follower_user_id" binding:"required"`
		FollowedUserID int `uri:"followed_user_id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := r.Usecase.CompleteFeedUserSubscribeQueue(c.Request.Context(), req.FollowerUserID, req.FollowedUserID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase CompleteFeedUserSubscribeQueue followerUserID=%d followedUserID=%d", req.FollowerUserID, req.FollowedUserID), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Status(http.StatusOK)
}

func (r *rest) handlerGetPlaceMenu(c *gin.Context) {
	var req struct {
		PlaceID int `uri:"place_id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resp, err := r.Usecase.GetPlaceMenu(c.Request.Context(), req.PlaceID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetPlaceMenu placeID=%d", req.PlaceID), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
}
