package memory

import (
	errs "cafe/pkg/db_manager/errors"
	"cafe/pkg/models"
	"context"
	"errors"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"reflect"
)

func init() {
	orm.RegisterTable((*models.ReviewReviewMedias)(nil))
	orm.RegisterTable((*models.AdvertAdvertMedias)(nil))
}

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

func (d *DbManager) GetAdvertsByPlaceID(ctx context.Context, placeID int, lastAdvertID int, limit int) ([]models.Advert, error) {
	var res []models.Advert
	query := d.db.Model(&res).Where("place_id = ?", placeID)
	if lastAdvertID != 0 {
		query = query.Where("advert_id < ?", lastAdvertID)
	}
	query = query.Relation("AdvertMedias", func(q *pg.Query) (*pg.Query, error) {
		q = q.OrderExpr("advert_advert_medias.order ASC")
		return q, nil
	})
	query = query.Order("publish_datetime desc")
	if limit != 0 {
		query = query.Limit(limit)
	}
	if err := query.Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) GetAdvertByID(ctx context.Context, advertID int) (models.Advert, error) {
	res := models.Advert{ID: advertID}
	if err := d.db.Model(&res).WherePK().Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		err = handleSqlError(err, reflect.TypeOf(res))
		return models.Advert{}, err
	}
	return res, nil
}

func (d *DbManager) GetUserPlaceSubscriptionsByPlaceID(ctx context.Context, placeID int) ([]models.UserPlaceSubscription, error) {
	var res []models.UserPlaceSubscription
	if err := d.db.Model(&res).Where("place_id = ?", placeID).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		err = handleSqlError(err, reflect.TypeOf(res))
		return nil, err
	}
	return res, nil
}

func (d *DbManager) GetUsersPlacesSubscriptionsByUserID(ctx context.Context, userID int) ([]models.UserPlaceSubscription, error) {
	var res []models.UserPlaceSubscription
	if err := d.db.Model(&res).Where("user_id = ?", userID).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		err = handleSqlError(err, reflect.TypeOf(res))
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
		err = wrapErr.NewWrapErr(fmt.Errorf("insert placeEvaluation=%+v into db", *placeEvaluation), err)
		return err
	}

	for i := range marks {
		marks[i].PlaceEvaluationID = placeEvaluation.ID
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

func (d *DbManager) GetReviewsByUserID(ctx context.Context, userID int, lastReviewID int, limit int) ([]models.Review, error) {
	var res []models.Review
	query := d.db.Model(&res).Where("user_id = ?", userID)
	if lastReviewID != 0 {
		query = query.Where("review_id < ?", lastReviewID)
	}
	query = query.Relation("ReviewMedias", func(q *pg.Query) (*pg.Query, error) {
		q = q.OrderExpr("review_review_medias.order ASC")
		return q, nil
	})
	query = query.Order("publish_datetime desc")
	if limit != 0 {
		query = query.Limit(limit)
	}
	if err := query.Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) GetReviewByID(ctx context.Context, reviewID int) (models.Review, error) {
	res := models.Review{ID: reviewID}
	if err := d.db.Model(&res).WherePK().Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		err = handleSqlError(err, reflect.TypeOf(res))
		return models.Review{}, err
	}
	return res, nil
}

func (d *DbManager) AddReview(ctx context.Context, review *models.Review) error {
	if _, err := d.db.Model(review).Returning("*").Insert(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("insert review=%+v into db", review), err)
		return err
	}
	return nil
}

func (d *DbManager) GetAllReviewMedias(ctx context.Context) ([]models.ReviewMedia, error) {
	var res []models.ReviewMedia
	if err := d.db.Model(&res).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) GetAllReviewReviewMedias(ctx context.Context) ([]models.ReviewReviewMedias, error) {
	var res []models.ReviewReviewMedias
	if err := d.db.Model(&res).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) AddReviewReviewMedias(ctx context.Context, reviewReviewMedias []models.ReviewReviewMedias) error {
	if _, err := d.db.Model(&reviewReviewMedias).Returning("*").Insert(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("insert reviewReviewMedias into db"), err)
		return err
	}
	return nil
}

func (d *DbManager) AddReviewMedia(ctx context.Context, reviewMedia *models.ReviewMedia) error {
	if _, err := d.db.Model(reviewMedia).Returning("*").Insert(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("insert reviewMedia=%+v into db", *reviewMedia), err)
		return err
	}
	return nil
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

func (d *DbManager) GetUserSubscriptionsByFollowedUserID(ctx context.Context, followedUserID int) ([]models.UserSubscription, error) {
	var res []models.UserSubscription
	if err := d.db.Model(&res).Where("followed_user_id = ?", followedUserID).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		err = handleSqlError(err, reflect.TypeOf(res))
		return nil, err
	}
	return res, nil
}

func (d *DbManager) GetAllPlaceSubscriptions(ctx context.Context) ([]models.UserPlaceSubscription, error) {
	var res []models.UserPlaceSubscription
	if err := d.db.Model(&res).Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) AddPlaceSubscription(ctx context.Context, userPlaceSubscription *models.UserPlaceSubscription) error {
	if _, err := d.db.Model(userPlaceSubscription).Returning("*").Insert(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("into into db userPlaceSubscription=%+v", userPlaceSubscription), err)
		err = handleSqlError(err, reflect.TypeOf(*userPlaceSubscription))
		if errors.Is(err, errs.ErrUniqueViolation(nil)) {
			err = wrapErr.NewWrapErr(errs.ErrorEntityAlreadyExists, err)
		}
		return err
	}
	return nil
}

func (d *DbManager) DeletePlaceSubscription(ctx context.Context, userPlaceSubscription models.UserPlaceSubscription) error {
	if _, err := d.db.Model(&userPlaceSubscription).WherePK().Delete(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("delete from db userPlaceSubscription=%+v", userPlaceSubscription), err)
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

func (d *DbManager) GetUsersFeedOfUserID(ctx context.Context, userID int, lastUserFeedID int, limit int) ([]models.UserFeed, error) {
	var res []models.UserFeed
	query := d.db.Model(&res).Where("user_feed.user_id = ?", userID)
	if lastUserFeedID != 0 {
		query = query.Where("users_feed_id < ?", lastUserFeedID)
	}
	query = query.OrderExpr("publish_datetime DESC").
		Limit(limit)
	query = query.Relation("Advert")
	query = query.Relation("Review")
	if err := query.Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select from db"), err)
		return nil, err
	}
	return res, nil
}

func (d *DbManager) AddUsersFeed(ctx context.Context, usersFeed []models.UserFeed) error {
	if _, err := d.db.Model(&usersFeed).Returning("*").Insert(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("insert usersFeed into db"), err)
		return err
	}
	return nil
}

//func (d *DbManager) DeleteUsersFeedsByUserIDAndFollowedUserID(ctx context.Context, userID int, followedUserID int) error {
//	userFeeds := models.UserFeed{
//		UserID:         userID,
//		FollowedUserID: followedUserID,
//	}
//	if _, err := d.db.Model(&userFeeds).Delete(); err != nil {
//		err = wrapErr.NewWrapErr(fmt.Errorf("delete usersFeed with userID=%d followedUserID=%d from db", userID, followedUserID), err)
//		return err
//	}
//	return nil
//}

func (d *DbManager) DeleteUsersFeeds(ctx context.Context, userFeeds models.UserFeed) error {
	query := d.db.Model(&models.UserFeed{})
	if userFeeds.UserID != 0 {
		query = query.Where("user_id = ?", userFeeds.UserID)
	}
	if userFeeds.FollowedUserID != 0 {
		query = query.Where("followed_user_id = ?", userFeeds.FollowedUserID)
	}
	if userFeeds.PlaceID != 0 {
		query = query.Where("place_id = ?", userFeeds.PlaceID)
	}
	if _, err := query.Delete(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("delete usersFeed userFeeds=%+v from db", userFeeds), err)
		return err
	}
	return nil
}

func (d *DbManager) AddFeedAdvertQueue(ctx context.Context, feedAdvertQueue *models.FeedAdvertQueue) error {
	if _, err := d.db.Model(feedAdvertQueue).OnConflict("(advert_id) do nothing").Insert(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("insert feedAdvertQueue=%+v into db", *feedAdvertQueue), err)
		return err
	}
	return nil
}

func (d *DbManager) PollFeedAdvertQueue(ctx context.Context) (models.FeedAdvertQueue, error) {
	var feedAdvertQueue models.FeedAdvertQueue
	if _, err := d.db.Query(&feedAdvertQueue, `
		with next_task as (
			select advert_id from main.feed_advert_queue
			where status = 0
			limit 1
				for update skip locked
		)
		update main.feed_advert_queue
		set
			status = 1,
			change_status_datetime = now()
		from next_task
		where main.feed_advert_queue.advert_id = next_task.advert_id
		returning main.feed_advert_queue.advert_id;
	`); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("query to db"), err)
		if errors.Is(err, pg.ErrNoRows) {
			err = wrapErr.NewWrapErr(errs.ErrorQueueIsEmpty, err)
		}
		return models.FeedAdvertQueue{}, err
	}
	if feedAdvertQueue.AdvertID == 0 {
		return models.FeedAdvertQueue{}, wrapErr.NewWrapErr(errs.ErrorQueueIsEmpty, nil)
	}
	return feedAdvertQueue, nil
}

func (d *DbManager) CompleteFeedAdvertQueue(ctx context.Context, advertID int) error {
	feedAdvertQueue := models.FeedAdvertQueue{AdvertID: advertID}
	if _, err := d.db.Model(&feedAdvertQueue).WherePK().Set("status = ?, change_status_datetime = now()", 2).Update(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("set status to 2 feedAdvertQueue = %+v in db", feedAdvertQueue), err)
		err = handleSqlError(err, reflect.TypeOf(feedAdvertQueue))
		return err
	}
	return nil
}

func (d *DbManager) AddFeedReviewQueue(ctx context.Context, feedReviewQueue *models.FeedReviewQueue) error {
	if _, err := d.db.Model(feedReviewQueue).OnConflict("(review_id) do nothing").Insert(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("insert feedReviewQueue=%+v into db", *feedReviewQueue), err)
		return err
	}
	return nil
}

func (d *DbManager) PollFeedReviewQueue(ctx context.Context) (models.FeedReviewQueue, error) {
	var feedReviewQueue models.FeedReviewQueue
	if _, err := d.db.Query(&feedReviewQueue, `
		with next_task as (
			select review_id from main.feed_review_queue
			where status = 0
			limit 1
				for update skip locked
		)
		update main.feed_review_queue
		set
			status = 1,
			change_status_datetime = now()
		from next_task
		where main.feed_review_queue.review_id = next_task.review_id
		returning main.feed_review_queue.review_id;
	`); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("query to db"), err)
		if errors.Is(err, pg.ErrNoRows) {
			err = wrapErr.NewWrapErr(errs.ErrorQueueIsEmpty, err)
		}
		return models.FeedReviewQueue{}, err
	}
	if feedReviewQueue.ReviewID == 0 {
		return models.FeedReviewQueue{}, wrapErr.NewWrapErr(errs.ErrorQueueIsEmpty, nil)
	}
	return feedReviewQueue, nil
}

func (d *DbManager) CompleteFeedReviewQueue(ctx context.Context, reviewID int) error {
	feedReviewQueue := models.FeedReviewQueue{ReviewID: reviewID}
	if _, err := d.db.Model(&feedReviewQueue).WherePK().Set("status = ?, change_status_datetime = now()", 2).Update(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("set status to 2 feedReviewQueue = %+v in db", feedReviewQueue), err)
		err = handleSqlError(err, reflect.TypeOf(feedReviewQueue))
		return err
	}
	return nil
}

func (d *DbManager) AddFeedUserSubscribeQueue(ctx context.Context, feedUserSubscribeQueue *models.FeedUserSubscribeQueue) error {
	if _, err := d.db.Model(feedUserSubscribeQueue).OnConflict("(follower_user_id, followed_user_id) do nothing").Insert(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("insert feedUserSubscribeQueue=%+v into db", *feedUserSubscribeQueue), err)
		return err
	}
	return nil
}

func (d *DbManager) PollFeedUserSubscribeQueue(ctx context.Context) (models.FeedUserSubscribeQueue, error) {
	var feedUserSubscribeQueue models.FeedUserSubscribeQueue
	if _, err := d.db.Query(&feedUserSubscribeQueue, `
		with next_task as (
			select follower_user_id, followed_user_id from main.feed_user_subscribe_queue
			where status = 0
			limit 1
				for update skip locked
		)
		update main.feed_user_subscribe_queue
		set
			status = 1,
			change_status_datetime = now()
		from next_task
		where feed_user_subscribe_queue.follower_user_id = next_task.follower_user_id and feed_user_subscribe_queue.followed_user_id = next_task.followed_user_id
		returning feed_user_subscribe_queue.follower_user_id, feed_user_subscribe_queue.followed_user_id;
	`); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("query to db"), err)
		if errors.Is(err, pg.ErrNoRows) {
			err = wrapErr.NewWrapErr(errs.ErrorQueueIsEmpty, err)
		}
		return models.FeedUserSubscribeQueue{}, err
	}
	if feedUserSubscribeQueue.FollowerUserID == 0 && feedUserSubscribeQueue.FollowedUserID == 0 {
		return models.FeedUserSubscribeQueue{}, wrapErr.NewWrapErr(errs.ErrorQueueIsEmpty, nil)
	}
	return feedUserSubscribeQueue, nil
}

func (d *DbManager) CompleteFeedUserSubscribeQueue(ctx context.Context, followerUserID int, followedUserID int) error {
	feedUserSubscribeQueue := models.FeedUserSubscribeQueue{FollowerUserID: followerUserID, FollowedUserID: followedUserID}
	if _, err := d.db.Model(&feedUserSubscribeQueue).WherePK().Set("status = ?, change_status_datetime = now()", 2).Update(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("set status to 2 feedUserSubscribeQueue = %+v in db", feedUserSubscribeQueue), err)
		err = handleSqlError(err, reflect.TypeOf(feedUserSubscribeQueue))
		return err
	}
	return nil
}

func (d *DbManager) GetFullPlaceMenu(ctx context.Context, placeID int) (models.PlaceMenu, error) {
	placeMenu := models.PlaceMenu{
		PlaceID:             placeID,
		PlaceMenuCategories: nil,
	}

	query := d.db.Model(&placeMenu.PlaceMenuCategories).
		Where("place_id = ?", placeID).
		Order("order")

	query.Relation("PlaceMenuItems", func(query *orm.Query) (*orm.Query, error) {
		query = query.Relation("PlaceMenuItemMedia")
		query = query.Where("publish_datetime <= now()").Order("order")
		return query, nil
	})

	if err := query.Select(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("select place menu from db"), err)
		return models.PlaceMenu{}, err
	}

	return placeMenu, nil
}
