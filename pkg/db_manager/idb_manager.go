package db_manager

import (
	"cafe/pkg/models"
	"context"
)

type IDBManager interface {
	GetAllPlaces(ctx context.Context) ([]models.Place, error)
	GetAllPlaceMedias(ctx context.Context) ([]models.PlaceMedia, error)
	GetAllCategories(ctx context.Context) ([]models.Category, error)
	GetAllPlaceCategories(ctx context.Context) ([]models.PlaceCategory, error)
	GetAllKitchenCategories(ctx context.Context) ([]models.KitchenCategory, error)
	GetAllPlaceKitchenCategories(ctx context.Context) ([]models.PlaceKitchenCategory, error)
	GetAllPlaceSchedules(ctx context.Context) ([]models.PlaceSchedule, error)
	GetAllAdverts(ctx context.Context) ([]models.Advert, error)
	GetAllAdvertMedias(ctx context.Context) ([]models.AdvertMedia, error)
	GetAllEvaluationCriterions(ctx context.Context) ([]models.EvaluationCriterion, error)
	GetAllPlaceEvaluations(ctx context.Context) ([]models.PlaceEvaluation, error)
	GetAllPlaceEvaluationMarks(ctx context.Context) ([]models.PlaceEvaluationMark, error)
	GetAllReviews(ctx context.Context) ([]models.Review, error)
	GetAllReviewMedias(ctx context.Context) ([]models.ReviewMedia, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	GetUserByUserID(ctx context.Context, userID int) (models.User, error)
	GetUserByUserName(ctx context.Context, userName string) (models.User, error)
	GetUserByVerifiedPhone(ctx context.Context, phone string) (models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	GetAllUserSubscriptions(ctx context.Context) ([]models.UserSubscription, error)
	GetActualUserPhoneCodeByUserID(ctx context.Context, userID int) (models.UserPhoneCode, error)
	CreateUserPhoneCode(ctx context.Context, userPhoneCode *models.UserPhoneCode) error
	ActivateUserPhone(ctx context.Context, userPhoneCodeID int) error
}
