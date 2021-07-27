package memory

import (
	errs "cafe/pkg/db_manager/errors"
	"cafe/pkg/models"
	"context"
	"errors"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/go-pg/pg/v9"
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

func (d *DbManager) GetAllUserSubscriptions(ctx context.Context) ([]models.UserSubscription, error) {
	var res []models.UserSubscription
	if err := d.db.Model(&res).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}