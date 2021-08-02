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
	router.GET("/evaluation-criterions", r.handlerGetAllEvaluationCriterions)
	router.GET("/place-evaluations", r.handlerGetAllPlaceEvaluations)
	router.GET("/place-evaluation-marks", r.handlerGetAllPlaceEvaluationMarks)
	router.GET("/reviews", r.handlerGetAllReviews)
	router.GET("/review-medias", r.handlerGetAllReviewMedias)
	router.GET("/users", r.handlerGetAllUsers)
	router.GET("/user-by-id/:user_id", r.handlerGetUserByID)
	router.GET("/user-by-name/:name", r.handlerGetUserByName)
	router.GET("/user-by-phone/:phone", r.handlerGetUserByVerifiedPhone)
	router.POST("/user", r.handlerCreateUser)
	router.GET("/user-subscriptions", r.handlerGetAllUserSubscriptions)
	router.GET("/actual-user-phone-code-by-user-id/:user_id", r.handlerGetActualUserPhoneCodeByUserID)
	router.POST("/user-phone-code", r.handlerCreateUserPhoneCode)
	router.POST("/activate-user-phone", r.handlerActivateUserPhone)
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

func (r *rest) handlerGetAllReviewMedias(c *gin.Context) {
	resp, err := r.Usecase.GetAllReviewMedias(c.Request.Context())
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetAllReviewMedias"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}
	c.Data(http.StatusOK, "application/x-gob", utils.ToGobBytes(resp))
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
