package rest

import (
	"cafe/pkg/client_gateway_service/usecase"
	"cafe/pkg/client_sso"
	"cafe/pkg/client_sso/models"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/gin-gonic/gin"
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
	router.POST("/auth/login", r.handlerLogin)
	router.POST("/auth/refresh-token", r.handlerRefreshToken)
	router.POST("/auth/signup", r.handlerSignUp)
	router.POST("/auth/approve-phone", r.handlerApprovePhone)

	router.GET("/places", r.handlerGetPlaces)
	router.GET("/place-by-id/:id", r.handlerGetPlaceByID)

	authorized := router.Group("/")
	authorized.Use(r.authMiddleware())
	authorized.POST("/auth/logout", r.handlerLogout)
	authorized.GET("/users/subscriptions", r.handlerGetUserSubscriptions)
	authorized.POST("/user/:id/subscribe", r.handlerSubscribeToUser)
	authorized.POST("/user/:id/unsubscribe", r.handlerUnsubscribeFromUser)
}

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

	c.JSON(http.StatusOK, models.Convert(places, modelTag))
}
