package usecase

import (
	"cafe/pkg/db_manager"
	"cafe/pkg/models"
	"context"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
)

type Usecase struct {
	dbManager db_manager.IDBManager
}

func NewUsecase(dbManager db_manager.IDBManager) *Usecase {
	return &Usecase{dbManager: dbManager}
}

func (u *Usecase) GetAllPlaces(ctx context.Context) ([]models.Place, error) {
	res, err := u.dbManager.GetAllPlaces(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllPlaces"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) GetAllPlaceMedias(ctx context.Context) ([]models.PlaceMedia, error) {
	res, err := u.dbManager.GetAllPlaceMedias(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllPlaceMedias"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	res, err := u.dbManager.GetAllCategories(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllCategories"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) GetAllPlaceCategories(ctx context.Context) ([]models.PlaceCategory, error) {
	res, err := u.dbManager.GetAllPlaceCategories(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllPlaceCategories"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) GetAllKitchenCategories(ctx context.Context) ([]models.KitchenCategory, error) {
	res, err := u.dbManager.GetAllKitchenCategories(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllKitchenCategories"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) GetAllPlaceKitchenCategories(ctx context.Context) ([]models.PlaceKitchenCategory, error) {
	res, err := u.dbManager.GetAllPlaceKitchenCategories(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllPlaceKitchenCategories"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) GetAllPlaceSchedules(ctx context.Context) ([]models.PlaceSchedule, error) {
	res, err := u.dbManager.GetAllPlaceSchedules(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllPlaceSchedules"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) GetAllAdverts(ctx context.Context) ([]models.Advert, error) {
	res, err := u.dbManager.GetAllAdverts(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllAdverts"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) GetAllAdvertMedias(ctx context.Context) ([]models.AdvertMedia, error) {
	res, err := u.dbManager.GetAllAdvertMedias(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllAdvertMedias"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) GetAdvertsByPlaceID(ctx context.Context, placeID int, lastAdvertID int, limit int) ([]models.Advert, error) {
	res, err := u.dbManager.GetAdvertsByPlaceID(ctx, placeID, lastAdvertID, limit)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAdvertsByPlaceID"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) GetAdvertByID(ctx context.Context, advertID int) (models.Advert, error) {
	res, err := u.dbManager.GetAdvertByID(ctx, advertID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAdvertByID"), err)
		return models.Advert{}, err
	}
	return res, nil
}

func (u *Usecase) GetUserPlaceSubscriptionsByPlaceID(ctx context.Context, placeID int) ([]models.UserPlaceSubscription, error) {
	res, err := u.dbManager.GetUserPlaceSubscriptionsByPlaceID(ctx, placeID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetUserPlaceSubscriptionsByPlaceID"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) GetUsersPlacesByUserID(ctx context.Context, userID int) ([]models.UserPlaceSubscription, error) {
	res, err := u.dbManager.GetUsersPlacesSubscriptionsByUserID(ctx, userID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetUsersPlacesSubscriptionsByUserID"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) GetAllEvaluationCriterions(ctx context.Context) ([]models.EvaluationCriterion, error) {
	res, err := u.dbManager.GetAllEvaluationCriterions(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllEvaluationCriterions"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) GetAllPlaceEvaluations(ctx context.Context) ([]models.PlaceEvaluation, error) {
	res, err := u.dbManager.GetAllPlaceEvaluations(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllPlaceEvaluations"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) AddPlaceEvaluationWithMarks(ctx context.Context, placeEvaluation *models.PlaceEvaluation, marks []models.PlaceEvaluationMark) error {
	if err := u.dbManager.AddPlaceEvaluationWithMarks(ctx, placeEvaluation, marks); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager AddPlaceEvaluationWithMarks"), err)
		return err
	}
	return nil
}

func (u *Usecase) GetAllPlaceEvaluationMarks(ctx context.Context) ([]models.PlaceEvaluationMark, error) {
	res, err := u.dbManager.GetAllPlaceEvaluationMarks(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllPlaceEvaluationMarks"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) GetAllReviews(ctx context.Context) ([]models.Review, error) {
	res, err := u.dbManager.GetAllReviews(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllReviews"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) GetReviewsByUserID(ctx context.Context, userID int, lastReviewID int, limit int) ([]models.Review, error) {
	res, err := u.dbManager.GetReviewsByUserID(ctx, userID, lastReviewID, limit)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetReviewsByUserID"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) GetReviewByID(ctx context.Context, reviewID int) (models.Review, error) {
	res, err := u.dbManager.GetReviewByID(ctx, reviewID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetReviewByID"), err)
		return models.Review{}, err
	}
	return res, nil
}

func (u *Usecase) AddReview(ctx context.Context, review models.Review) (models.Review, error) {
	err := u.dbManager.AddReview(ctx, &review)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager AddReview"), err)
		return models.Review{}, err
	}
	return review, nil
}

func (u *Usecase) AddReviewMedia(ctx context.Context, reviewMedia models.ReviewMedia) (models.ReviewMedia, error) {
	err := u.dbManager.AddReviewMedia(ctx, &reviewMedia)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager AddReviewMedia"), err)
		return models.ReviewMedia{}, err
	}
	return reviewMedia, nil
}

func (u *Usecase) GetAllReviewMedias(ctx context.Context) ([]models.ReviewMedia, error) {
	res, err := u.dbManager.GetAllReviewMedias(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllReviewMedias"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) GetAllReviewReviewMedias(ctx context.Context) ([]models.ReviewReviewMedias, error) {
	res, err := u.dbManager.GetAllReviewReviewMedias(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllReviewReviewMedias"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) AddReviewReviewMedias(ctx context.Context, reviewReviewMedias []models.ReviewReviewMedias) error {
	err := u.dbManager.AddReviewReviewMedias(ctx, reviewReviewMedias)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager AddReviewReviewMedias"), err)
		return err
	}
	return nil
}

func (u *Usecase) GetAllUsers(ctx context.Context) ([]models.User, error) {
	res, err := u.dbManager.GetAllUsers(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllUsers"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) GetUserByID(ctx context.Context, userID int) (models.User, error) {
	res, err := u.dbManager.GetUserByUserID(ctx, userID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetUserByUserID"), err)
		return models.User{}, err
	}
	return res, nil
}

func (u *Usecase) GetUserByName(ctx context.Context, name string) (models.User, error) {
	res, err := u.dbManager.GetUserByUserName(ctx, name)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetUserByUserName"), err)
		return models.User{}, err
	}
	return res, nil
}

func (u *Usecase) GetUserByVerifiedPhone(ctx context.Context, phone string) (models.User, error) {
	res, err := u.dbManager.GetUserByVerifiedPhone(ctx, phone)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetUserByVerifiedPhone"), err)
		return models.User{}, err
	}
	return res, nil
}

func (u *Usecase) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	if err := u.dbManager.CreateUser(ctx, &user); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager CreateUser user=%+v", user), err)
		return models.User{}, err
	}
	return user, nil
}

func (u *Usecase) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	if err := u.dbManager.UpdateUser(ctx, &user); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager UpdateUser user=%+v", user), err)
		return models.User{}, err
	}
	return user, nil
}

func (u *Usecase) GetAllUserSubscriptions(ctx context.Context) ([]models.UserSubscription, error) {
	res, err := u.dbManager.GetAllUserSubscriptions(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllUserSubscriptions"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) AddUserSubscription(ctx context.Context, userSubscription models.UserSubscription) (models.UserSubscription, error) {
	if err := u.dbManager.AddUserSubscription(ctx, &userSubscription); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager AddUserSubscription"), err)
		return models.UserSubscription{}, err
	}
	return userSubscription, nil
}

func (u *Usecase) DeleteUserSubscription(ctx context.Context, userSubscription models.UserSubscription) error {
	if err := u.dbManager.DeleteUserSubscription(ctx, userSubscription); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager DeleteUserSubscription"), err)
		return err
	}
	return nil
}

func (u *Usecase) GetUserSubscriptionsByFollowedUserID(ctx context.Context, followedUserID int) ([]models.UserSubscription, error) {
	res, err := u.dbManager.GetUserSubscriptionsByFollowedUserID(ctx, followedUserID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetUserSubscriptionsByFollowedUserID followedUserID=%d", followedUserID), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) GetAllPlaceSubscriptions(ctx context.Context) ([]models.UserPlaceSubscription, error) {
	res, err := u.dbManager.GetAllPlaceSubscriptions(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllPlaceSubscriptions"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) AddPlaceSubscription(ctx context.Context, userPlaceSubscription models.UserPlaceSubscription) (models.UserPlaceSubscription, error) {
	if err := u.dbManager.AddPlaceSubscription(ctx, &userPlaceSubscription); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager AddPlaceSubscription"), err)
		return models.UserPlaceSubscription{}, err
	}
	return userPlaceSubscription, nil
}

func (u *Usecase) DeletePlaceSubscription(ctx context.Context, userPlaceSubscription models.UserPlaceSubscription) error {
	if err := u.dbManager.DeletePlaceSubscription(ctx, userPlaceSubscription); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager DeletePlaceSubscription"), err)
		return err
	}
	return nil
}

func (u *Usecase) GetActualUserPhoneCodeByUserID(ctx context.Context, userID int) (models.UserPhoneCode, error) {
	res, err := u.dbManager.GetActualUserPhoneCodeByUserID(ctx, userID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetActualUserPhoneCodeByUserID"), err)
		return models.UserPhoneCode{}, err
	}
	return res, nil
}

func (u *Usecase) CreateUserPhoneCode(ctx context.Context, userPhoneCode models.UserPhoneCode) (models.UserPhoneCode, error) {
	if err := u.dbManager.CreateUserPhoneCode(ctx, &userPhoneCode); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager CreateUserPhoneCode"), err)
		return models.UserPhoneCode{}, err
	}
	return userPhoneCode, nil
}

func (u *Usecase) UpdateUserPhoneCode(ctx context.Context, userPhoneCode models.UserPhoneCode) (models.UserPhoneCode, error) {
	if err := u.dbManager.UpdateUserPhoneCode(ctx, &userPhoneCode); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager UpdateUserPhoneCode"), err)
		return models.UserPhoneCode{}, err
	}
	return userPhoneCode, nil
}

func (u *Usecase) ActivateUserPhone(ctx context.Context, userPhoneCodeID int) error {
	if err := u.dbManager.ActivateUserPhone(ctx, userPhoneCodeID); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager ActivateUserPhone"), err)
		return err
	}
	return nil
}

func (u *Usecase) GetFeedOfUserID(ctx context.Context, userID int, lastUserFeedID int, limit int) ([]models.UserFeed, error) {
	res, err := u.dbManager.GetUsersFeedOfUserID(ctx, userID, lastUserFeedID, limit)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetUsersFeedOfUserID"), err)
		return nil, err
	}
	return res, nil
}

func (u *Usecase) AddUsersFeed(ctx context.Context, usersFeeds []models.UserFeed) error {
	if len(usersFeeds) == 0 {
		return nil
	}
	err := u.dbManager.AddUsersFeed(ctx, usersFeeds)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager AddUsersFeed"), err)
		return err
	}
	return nil
}

func (u *Usecase) AddFeedAdvertQueue(ctx context.Context, feedAdvertQueue models.FeedAdvertQueue) (models.FeedAdvertQueue, error) {
	err := u.dbManager.AddFeedAdvertQueue(ctx, &feedAdvertQueue)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager feedAdvertQueue"), err)
		return models.FeedAdvertQueue{}, err
	}
	return feedAdvertQueue, nil
}

func (u *Usecase) PollFeedAdvertQueue(ctx context.Context) (models.FeedAdvertQueue, error) {
	feedAdvertQueue, err := u.dbManager.PollFeedAdvertQueue(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager PollFeedAdvertQueue"), err)
		return models.FeedAdvertQueue{}, err
	}
	return feedAdvertQueue, nil
}

func (u *Usecase) CompleteFeedAdvertQueue(ctx context.Context, advertID int) error {
	if err := u.dbManager.CompleteFeedAdvertQueue(ctx, advertID); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager CompleteFeedAdvertQueue"), err)
		return err
	}
	return nil
}

func (u *Usecase) AddFeedReviewQueue(ctx context.Context, feedReviewQueue models.FeedReviewQueue) (models.FeedReviewQueue, error) {
	err := u.dbManager.AddFeedReviewQueue(ctx, &feedReviewQueue)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager AddFeedReviewQueue"), err)
		return models.FeedReviewQueue{}, err
	}
	return feedReviewQueue, nil
}

func (u *Usecase) PollFeedReviewQueue(ctx context.Context) (models.FeedReviewQueue, error) {
	feedReviewQueue, err := u.dbManager.PollFeedReviewQueue(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager PollFeedReviewQueue"), err)
		return models.FeedReviewQueue{}, err
	}
	return feedReviewQueue, nil
}

func (u *Usecase) CompleteFeedReviewQueue(ctx context.Context, reviewID int) error {
	if err := u.dbManager.CompleteFeedReviewQueue(ctx, reviewID); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager CompleteFeedReviewQueue"), err)
		return err
	}
	return nil
}

func (u *Usecase) AddFeedUserSubscribeQueue(ctx context.Context, feedUserSubscribeQueue models.FeedUserSubscribeQueue) (models.FeedUserSubscribeQueue, error) {
	err := u.dbManager.AddFeedUserSubscribeQueue(ctx, &feedUserSubscribeQueue)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager AddFeedReviewQueue"), err)
		return models.FeedUserSubscribeQueue{}, err
	}
	return feedUserSubscribeQueue, nil
}

func (u *Usecase) PollFeedUserSubscribeQueue(ctx context.Context) (models.FeedUserSubscribeQueue, error) {
	feedUserSubscribeQueue, err := u.dbManager.PollFeedUserSubscribeQueue(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager PollFeedUserSubscribeQueue"), err)
		return models.FeedUserSubscribeQueue{}, err
	}
	return feedUserSubscribeQueue, nil
}

func (u *Usecase) CompleteFeedUserSubscribeQueue(ctx context.Context, followerUserID int, followedUserID int) error {
	if err := u.dbManager.CompleteFeedUserSubscribeQueue(ctx, followerUserID, followedUserID); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager CompleteFeedUserSubscribeQueue"), err)
		return err
	}
	return nil
}
