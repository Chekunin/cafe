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
	GetAdvertsByPlaceID(ctx context.Context, placeID int, lastAdvertID int, limit int) ([]models.Advert, error)
	GetAdvertByID(ctx context.Context, advertID int) (models.Advert, error)
	GetAllEvaluationCriterions(ctx context.Context) ([]models.EvaluationCriterion, error)
	AddPlaceEvaluationWithMarks(ctx context.Context, placeEvaluation *models.PlaceEvaluation, marks []models.PlaceEvaluationMark) error
	GetAllPlaceEvaluations(ctx context.Context) ([]models.PlaceEvaluation, error)
	GetAllPlaceEvaluationMarks(ctx context.Context) ([]models.PlaceEvaluationMark, error)
	GetAllReviews(ctx context.Context) ([]models.Review, error)
	GetReviewsByUserID(ctx context.Context, userID int, lastReviewID int, limit int) ([]models.Review, error)
	GetReviewByID(ctx context.Context, reviewID int) (models.Review, error)
	AddReview(ctx context.Context, review *models.Review) error
	GetAllReviewMedias(ctx context.Context) ([]models.ReviewMedia, error)
	AddReviewMedia(ctx context.Context, reviewMedia *models.ReviewMedia) error
	GetAllReviewReviewMedias(ctx context.Context) ([]models.ReviewReviewMedias, error)
	AddReviewReviewMedias(ctx context.Context, reviewReviewMedias []models.ReviewReviewMedias) error
	GetAllUsers(ctx context.Context) ([]models.User, error)
	GetUserByUserID(ctx context.Context, userID int) (models.User, error)
	GetUserByUserName(ctx context.Context, userName string) (models.User, error)
	GetUserByVerifiedPhone(ctx context.Context, phone string) (models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error

	GetAllUserSubscriptions(ctx context.Context) ([]models.UserSubscription, error)
	AddUserSubscription(ctx context.Context, userSubscription *models.UserSubscription) error
	DeleteUserSubscription(ctx context.Context, userSubscription models.UserSubscription) error
	GetUserSubscriptionsByFollowedUserID(ctx context.Context, followedUserID int) ([]models.UserSubscription, error)

	GetAllPlaceSubscriptions(ctx context.Context) ([]models.UserPlaceSubscription, error)
	AddPlaceSubscription(ctx context.Context, userPlaceSubscription *models.UserPlaceSubscription) error
	DeletePlaceSubscription(ctx context.Context, userPlaceSubscription models.UserPlaceSubscription) error

	GetActualUserPhoneCodeByUserID(ctx context.Context, userID int) (models.UserPhoneCode, error)
	CreateUserPhoneCode(ctx context.Context, userPhoneCode *models.UserPhoneCode) error
	UpdateUserPhoneCode(ctx context.Context, userPhoneCode *models.UserPhoneCode) error
	ActivateUserPhone(ctx context.Context, userPhoneCodeID int) error
	GetUsersFeedOfUserID(ctx context.Context, userID int, lastUserFeedID int, limit int) ([]models.UserFeed, error)
	AddUsersFeed(ctx context.Context, usersFeed []models.UserFeed) error
	//DeleteUsersFeedsByUserIDAndFollowedUserID(ctx context.Context, userID int, followedUserID int) error
	DeleteUsersFeeds(ctx context.Context, userFeeds models.UserFeed) error
	GetUserPlaceSubscriptionsByPlaceID(ctx context.Context, placeID int) ([]models.UserPlaceSubscription, error)
	GetUsersPlacesSubscriptionsByUserID(ctx context.Context, userID int) ([]models.UserPlaceSubscription, error)

	AddFeedAdvertQueue(ctx context.Context, feedAdvertQueue *models.FeedAdvertQueue) error
	PollFeedAdvertQueue(ctx context.Context) (models.FeedAdvertQueue, error)
	CompleteFeedAdvertQueue(ctx context.Context, advertID int) error

	AddFeedReviewQueue(ctx context.Context, feedReviewQueue *models.FeedReviewQueue) error
	PollFeedReviewQueue(ctx context.Context) (models.FeedReviewQueue, error)
	CompleteFeedReviewQueue(ctx context.Context, reviewID int) error

	AddFeedUserSubscribeQueue(ctx context.Context, feedUserSubscribeQueue *models.FeedUserSubscribeQueue) error
	PollFeedUserSubscribeQueue(ctx context.Context) (models.FeedUserSubscribeQueue, error)
	CompleteFeedUserSubscribeQueue(ctx context.Context, followerUserID int, followedUserID int) error

	GetFullPlaceMenu(ctx context.Context, placeID int) (models.PlaceMenu, error)
}
