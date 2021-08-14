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

func (u *Usecase) GetAllReviewMedias(ctx context.Context) ([]models.ReviewMedia, error) {
	res, err := u.dbManager.GetAllReviewMedias(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllReviewMedias"), err)
		return nil, err
	}
	return res, nil
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
