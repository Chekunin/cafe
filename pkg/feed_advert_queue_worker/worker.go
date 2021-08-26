package feed_advert_queue_worker

import (
	"cafe/pkg/common/catcherr"
	"cafe/pkg/db_manager"
	dbManagerErrs "cafe/pkg/db_manager/errors"
	"cafe/pkg/models"
	"context"
	"errors"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/gammazero/workerpool"
	"time"
)

type Config struct {
	CheckDBPeriod time.Duration `yaml:"check_db_period"`
	WorkerAmount  int           `yaml:"worker_amount"`
}

type WorkerPool struct {
	cfg       Config
	dbManager db_manager.IDBManager
	wp        *workerpool.WorkerPool
	done      chan struct{}
}

func NewWorkerPool(cfg Config, dbManager db_manager.IDBManager) (*WorkerPool, error) {
	wp := workerpool.New(cfg.WorkerAmount)

	worker := WorkerPool{
		cfg:       cfg,
		dbManager: dbManager,
		wp:        wp,
		done:      make(chan struct{}),
	}

	return &worker, nil
}

func (w *WorkerPool) Start() {
	go func() {
		for {
			if w.wp.Stopped() {
				break
			}
			// todo: надо доставать сразу по много задач
			task, err := w.dbManager.PollFeedAdvertQueue(context.TODO())
			if err != nil {
				err = wrapErr.NewWrapErr(fmt.Errorf("dbManager PollFeedAdvertQueue"), err)
				if errors.Is(err, dbManagerErrs.ErrorQueueIsEmpty) {
					time.Sleep(w.cfg.CheckDBPeriod)
					fmt.Printf("queue is empty\n")
					continue
				}
				catcherr.Catch(err)
			}
			w.wp.Submit(func() {
				if err := w.execFunc(task.AdvertID); err != nil {
					err = wrapErr.NewWrapErr(fmt.Errorf("execFunc advertID=%d", task.AdvertID), err)
					catcherr.Catch(err)
				}
			})
		}
	}()
}

func (w *WorkerPool) execFunc(advertID int) error {
	fmt.Printf("start execute task AdvertID=%d\n", advertID)

	// достаём всех пользователей, которые подписаны на данное заведение
	advert, err := w.dbManager.GetAdvertByID(context.TODO(), advertID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAdvertByID advertID=%d", advertID), err)
		return err
	}

	// достаём всех пользователей, которые подписаны на данное заведение и добавляем запись им в users_feeds
	userPlaces, err := w.dbManager.GetUsersPlacesSubscriptionsByPlaceID(context.TODO(), advert.PlaceID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetUsersPlacesSubscriptionsByPlaceID placeID=%d", advert.PlaceID), err)
		return err
	}

	usersFeed := make([]models.UserFeed, 0, len(userPlaces))
	for _, v := range userPlaces {
		usersFeed = append(usersFeed, models.UserFeed{
			UserID:          v.UserID,
			AdvertID:        advert.ID,
			PublishDatetime: advert.PublishDateTime,
		})
	}

	if err := w.dbManager.AddUsersFeed(context.TODO(), usersFeed); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager AddUsersFeed"), err)
		return err
	}

	fmt.Printf("finish execute task advertID=%d\n", advertID)
	return nil
}

func (w *WorkerPool) Stop() {
	w.wp.Stop()
}
