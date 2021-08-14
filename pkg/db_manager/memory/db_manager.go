package memory

import (
	errs "cafe/pkg/db_manager/errors"
	"cafe/pkg/models"
	"context"
	"errors"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/go-pg/pg/v9"
	"reflect"
)

type DbManager struct {
	db *pg.DB
}

type NewDbManagerParams struct {
	DB *pg.DB
}

func NewDbManager(params NewDbManagerParams) (*DbManager, error) {
	return &DbManager{
		db: params.DB,
	}, nil
}

func (d *DbManager) GetAllPlaces(ctx context.Context) ([]models.Place, error) {
	var res []models.Place
	if err := d.db.Model(&res).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) GetAllPlaceMedias(ctx context.Context) ([]models.PlaceMedia, error) {
	var res []models.PlaceMedia
	if err := d.db.Model(&res).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	var res []models.Category
	if err := d.db.Model(&res).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) GetAllPlaceCategories(ctx context.Context) ([]models.PlaceCategory, error) {
	var res []models.PlaceCategory
	if err := d.db.Model(&res).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) GetAllKitchenCategories(ctx context.Context) ([]models.KitchenCategory, error) {
	var res []models.KitchenCategory
	if err := d.db.Model(&res).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) GetAllPlaceKitchenCategories(ctx context.Context) ([]models.PlaceKitchenCategory, error) {
	var res []models.PlaceKitchenCategory
	if err := d.db.Model(&res).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) GetAllPlaceSchedules(ctx context.Context) ([]models.PlaceSchedule, error) {
	var res []models.PlaceSchedule
	if err := d.db.Model(&res).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) GetAllAdverts(ctx context.Context) ([]models.Advert, error) {
	var res []models.Advert
	if err := d.db.Model(&res).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) GetAllAdvertMedias(ctx context.Context) ([]models.AdvertMedia, error) {
	var res []models.AdvertMedia
	if err := d.db.Model(&res).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) GetAllEvaluationCriterions(ctx context.Context) ([]models.EvaluationCriterion, error) {
	var res []models.EvaluationCriterion
	if err := d.db.Model(&res).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) AddPlaceEvaluationWithMarks(ctx context.Context, placeEvaluation *models.PlaceEvaluation, marks []models.PlaceEvaluationMark) error {
	// todo: это выполнять в транзакции
	if _, err := d.db.Model(placeEvaluation).Returning("*").Insert(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("insert placeEvaluation=%+v into db", placeEvaluation), err)
		return err
	}

	if _, err := d.db.Model(&marks).Returning("*").Insert(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("insert marks=%+v into db", marks), err)
		return err
	}

	return nil
}

func (d *DbManager) GetAllPlaceEvaluations(ctx context.Context) ([]models.PlaceEvaluation, error) {
	var res []models.PlaceEvaluation
	if err := d.db.Model(&res).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) GetAllPlaceEvaluationMarks(ctx context.Context) ([]models.PlaceEvaluationMark, error) {
	var res []models.PlaceEvaluationMark
	if err := d.db.Model(&res).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) GetAllReviews(ctx context.Context) ([]models.Review, error) {
	var res []models.Review
	if err := d.db.Model(&res).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) GetAllReviewMedias(ctx context.Context) ([]models.ReviewMedia, error) {
	var res []models.ReviewMedia
	if err := d.db.Model(&res).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var res []models.User
	if err := d.db.Model(&res).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) GetUserByUserID(ctx context.Context, userID int) (models.User, error) {
	var res models.User
	if err := d.db.Model(&res).Where("user_id = ?", userID).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		if errors.Is(err, pg.ErrNoRows) {
			err = wrapErr.NewWrapErr(errs.ErrorEntityNotFound, err)
		}
		return models.User{}, err
	}
	return res, nil
}

func (d *DbManager) GetUserByUserName(ctx context.Context, userName string) (models.User, error) {
	var res models.User
	if err := d.db.Model(&res).Where("name = ?", userName).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		if errors.Is(err, pg.ErrNoRows) {
			err = wrapErr.NewWrapErr(errs.ErrorEntityNotFound, err)
		}
		return models.User{}, err
	}
	return res, nil
}

func (d *DbManager) GetUserByVerifiedPhone(ctx context.Context, phone string) (models.User, error) {
	var res models.User
	if err := d.db.Model(&res).Where("phone = ? and phone_verified is true", phone).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		if errors.Is(err, pg.ErrNoRows) {
			err = wrapErr.NewWrapErr(errs.ErrorEntityNotFound, err)
		}
		return models.User{}, err
	}
	return res, nil
}

func (d *DbManager) CreateUser(ctx context.Context, user *models.User) error {
	if _, err := d.db.Model(user).Insert(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("insert user = %+v to db", *user), err)
		err = handleSqlError(err, reflect.TypeOf(*user))
		return err
	}
	return nil
}

func (d *DbManager) UpdateUser(ctx context.Context, user *models.User) error {
	if _, err := d.db.Model(user).WherePK().Update(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("update user = %+v in db", *user), err)
		err = handleSqlError(err, reflect.TypeOf(*user))
		return err
	}
	return nil
}

func (d *DbManager) GetAllUserSubscriptions(ctx context.Context) ([]models.UserSubscription, error) {
	var res []models.UserSubscription
	if err := d.db.Model(&res).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) AddUserSubscription(ctx context.Context, userSubscription *models.UserSubscription) error {
	if _, err := d.db.Model(userSubscription).Returning("*").Insert(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("into into db userSubscription=%+v", userSubscription), err)
		err = handleSqlError(err, reflect.TypeOf(*userSubscription))
		if errors.Is(err, errs.ErrUniqueViolation(nil)) {
			err = wrapErr.NewWrapErr(errs.ErrorEntityAlreadyExists, err)
		}
		return err
	}
	return nil
}

func (d *DbManager) DeleteUserSubscription(ctx context.Context, userSubscription models.UserSubscription) error {
	if _, err := d.db.Model(&userSubscription).Delete(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("delete from db userSubscription=%+v", userSubscription), err)
		return err
	}
	return nil
}

func (d *DbManager) GetActualUserPhoneCodeByUserID(ctx context.Context, userID int) (models.UserPhoneCode, error) {
	var res models.UserPhoneCode
	if err := d.db.Model(&res).Where("user_id = ? and actual is true", userID).Select(); err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			err = wrapErr.NewWrapErr(errs.ErrorEntityNotFound, err)
		}
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return models.UserPhoneCode{}, err
	}
	return res, nil
}

func (d *DbManager) CreateUserPhoneCode(ctx context.Context, userPhoneCode *models.UserPhoneCode) error {
	if _, err := d.db.Model(userPhoneCode).Insert(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("insert userPhoneCode = %+v to db", *userPhoneCode), err)
		err = handleSqlError(err, reflect.TypeOf(*userPhoneCode))
		return err
	}
	return nil
}

func (d *DbManager) UpdateUserPhoneCode(ctx context.Context, userPhoneCode *models.UserPhoneCode) error {
	if _, err := d.db.Model(userPhoneCode).WherePK().Update(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("update userPhoneCode = %+v in db", *userPhoneCode), err)
		err = handleSqlError(err, reflect.TypeOf(*userPhoneCode))
		return err
	}
	return nil
}

func (d *DbManager) ActivateUserPhone(ctx context.Context, userPhoneCodeID int) error {
	userPhoneCode := models.UserPhoneCode{ID: userPhoneCodeID}
	if err := d.db.Model(&userPhoneCode).WherePK().Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return err
	}
	if _, err := d.db.Model(&userPhoneCode).WherePK().Set("actual = false").Update(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("set actual to false userPhoneCode = %+v in db", userPhoneCode), err)
		err = handleSqlError(err, reflect.TypeOf(userPhoneCode))
		return err
	}
	user := models.User{ID: userPhoneCode.UserID}
	if _, err := d.db.Model(&user).WherePK().Set("phone_verified = true").Update(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("set phone_verified to true userID = %d in db", userPhoneCode.UserID), err)
		err = handleSqlError(err, reflect.TypeOf(user))
		return err
	}
	return nil
}
