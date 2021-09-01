package workers

import (
	"cafe/pkg/db_manager"
	"cafe/pkg/feed_queue/models"
	commonModels "cafe/pkg/models"
	"context"
	"encoding/json"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/hibiken/asynq"
)

type Workers struct {
	srv       *asynq.Server
	mux       *asynq.ServeMux
	dbManager db_manager.IDBManager
}

type NewWorkersParams struct {
	RedisClientOpt asynq.RedisClientOpt
	DbManager      db_manager.IDBManager
}

func NewWorkers(params NewWorkersParams) *Workers {
	srv := asynq.NewServer(
		params.RedisClientOpt,
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	workers := Workers{
		srv:       srv,
		mux:       mux,
		dbManager: params.DbManager,
	}

	mux.HandleFunc(models.TypeNewAdvert, workers.HandleNewAdvertTask)
	mux.HandleFunc(models.TypeNewReview, workers.HandleNewReviewTask)

	return &workers
}

func (w *Workers) Run() error {
	if err := w.srv.Start(w.mux); err != nil {
		return err
	}
	return nil
}

func (w *Workers) Shutdown() {
	w.srv.Shutdown()
}

func (w *Workers) HandleNewAdvertTask(ctx context.Context, t *asynq.Task) error {
	var payload models.NewAdvertTaskPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	advert, err := w.dbManager.GetAdvertByID(context.TODO(), payload.AdvertID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAdvertByID advertID=%d", payload.AdvertID), err)
		return err
	}

	// достаём всех пользователей, которые подписаны на данное заведение и добавляем запись им в users_feeds
	userPlaces, err := w.dbManager.GetUserPlaceSubscriptionsByPlaceID(context.TODO(), advert.PlaceID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetUserPlaceSubscriptionsByPlaceID placeID=%d", advert.PlaceID), err)
		return err
	}

	usersFeed := make([]commonModels.UserFeed, 0, len(userPlaces))
	for _, v := range userPlaces {
		usersFeed = append(usersFeed, commonModels.UserFeed{
			UserID:          v.UserID,
			AdvertID:        advert.ID,
			PublishDatetime: advert.PublishDateTime,
		})
	}

	if err := w.dbManager.AddUsersFeed(context.TODO(), usersFeed); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager AddUsersFeed"), err)
		return err
	}

	fmt.Printf("finish execute task advertID=%d\n", payload.AdvertID)
	return nil
}

func (w *Workers) HandleNewReviewTask(ctx context.Context, t *asynq.Task) error {
	var payload models.NewReviewTaskPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	review, err := w.dbManager.GetReviewByID(context.TODO(), payload.ReviewID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetReviewByID reviewID=%d", payload.ReviewID), err)
		return err
	}

	// достаём всех пользователей, которые подписаны на данное заведение и добавляем запись им в users_feeds
	userSubscriptions, err := w.dbManager.GetUserSubscriptionsByFollowedUserID(context.TODO(), review.UserID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetUserSubscriptionsByFollowedUserID reviewID=%d", review.UserID), err)
		return err
	}

	usersFeed := make([]commonModels.UserFeed, 0, len(userSubscriptions))
	for _, v := range userSubscriptions {
		usersFeed = append(usersFeed, commonModels.UserFeed{
			UserID:          v.FollowerUserID,
			ReviewID:        review.ID,
			PublishDatetime: review.PublishDateTime,
		})
	}

	if err := w.dbManager.AddUsersFeed(context.TODO(), usersFeed); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager AddUsersFeed"), err)
		return err
	}

	fmt.Printf("finish execute task reviewID=%d\n", payload.ReviewID)
	return nil
}
