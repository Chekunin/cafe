package rest

import (
	"cafe/pkg/client_gateway_service/usecase"
	"cafe/pkg/client_sso"
	"cafe/pkg/client_sso/models"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

const modelTag string = "api"

type rest struct {
	usecase   *usecase.Usecase
	clientSso client_sso.IClientSso
}

func NewRest(router *gin.RouterGroup, usecase *usecase.Usecase, clientSso client_sso.IClientSso) *rest {
	rest := &rest{
		usecase:   usecase,
		clientSso: clientSso,
	}
	rest.routes(router)
	return rest
}

func (r *rest) routes(router *gin.RouterGroup) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/auth/login", r.handlerLogin)
	router.POST("/auth/refresh-token", r.handlerRefreshToken)
	router.POST("/auth/signup", r.handlerSignUp)
	router.POST("/auth/approve-phone", r.handlerApprovePhone)

	router.GET("/places", r.handlerGetPlaces)
	router.GET("/places/:id", r.handlerGetPlaceByID)

	router.GET("/place-review-medias/:id/data", r.handlerGetPlaceReviewMediaData)

	router.GET("/place-evaluation-criterions", r.handlerGetPlaceEvaluationCriterions)

	// записи конкретного пользователя
	router.GET("/users/:id/posts", r.handlerGetPlacesReviewsByUserID)

	// записи конкретного заведения
	router.GET("/places/:id/posts", r.handlerGetPlaceAdvertsByPlaceID)

	router.GET("/places/:id/menu", r.handlerGetPlaceMenu)

	authorized := router.Group("/")
	authorized.Use(r.authMiddleware())
	authorized.POST("/auth/logout", r.handlerLogout)
	authorized.GET("/users/:id/subscriptions", r.handlerGetUserSubscriptions)
	authorized.POST("/users/:id/subscribe", r.handlerSubscribeToUser)
	authorized.POST("/users/:id/unsubscribe", r.handlerUnsubscribeFromUser)

	authorized.GET("/users/:id/place-subscriptions", r.handlerGetPlaceSubscriptions)
	authorized.POST("/places/:id/subscribe", r.handlerSubscribeToPlace)
	authorized.POST("/places/:id/unsubscribe", r.handlerUnsubscribeFromPlace)

	authorized.POST("/places/:id/evaluation", r.handlerEvaluatePlace)
	authorized.GET("/places/:id/evaluation", r.handlerGetPlaceEvaluation)

	authorized.GET("/places/:id/evaluations", r.handlerGetPlaceEvaluations)

	authorized.POST("/place-review-media", r.handlerAddPlaceReviewMedia)
	authorized.POST("/places/:id/review", r.handlerAddPlaceReview)

	// запрос на получение своих собственных записей
	authorized.GET("/users/:id/place-reviews", r.handlerGetOwnPlacesReviews)
	// запрос на получение своей ленты (главной страницы с чужими записями)
	authorized.GET("/feed", r.handlerGetUserFeed)
}

// Places godoc
// @Summary Список ресторанов в заданном диапазоне координат.
// @Tags Заведения
// @Produce json
// @Param left_lng query float64 true "left_lng"
// @Param right_lng query float64 true "right_lng"
// @Param bottom_lat query float64 true "bottom_lat"
// @Param top_lat query float64 true "top_lat"
// @Success 200 {array} models.Place
// @Router /places [get]
func (r *rest) handlerGetPlaces(c *gin.Context) {
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

	places, err := r.usecase.GetPlaces(c.Request.Context(), req.LeftLng, req.RightLng, req.TopLat, req.BottomLat)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetPlaces"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	c.JSON(http.StatusOK, models.Convert(places, modelTag))
}

// Places godoc
// @Summary Ресторан по id
// @Tags Заведения
// @Produce json
// @Param id path int true "Place ID"
// @Success 200 {object} models.Place
// @Router /places/{id} [get]
func (r *rest) handlerGetPlaceByID(c *gin.Context) {
	var req struct {
		PlaceID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from uri"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	places, err := r.usecase.GetPlaceByID(c.Request.Context(), req.PlaceID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("usecase GetPlaceByID"), err)
		c.AbortWithError(GetHttpCode(err), err)
		return
	}

	// todo: здесь надо возвращать уникальный id картинки
	c.JSON(http.StatusOK, models.Convert(places, modelTag))
}
