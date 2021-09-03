package client

import (
	"cafe/pkg/feed_queue/models"
	"encoding/json"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/hibiken/asynq"
)

type Client struct {
	client *asynq.Client
}

type NewClientParams struct {
	RedisClientOpt asynq.RedisClientOpt
}

func NewClient(params NewClientParams) *Client {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: params.RedisClientOpt.Addr,
	})

	return &Client{client: client}
}

func (c *Client) AddNewAdvert(advertID int) error {
	payload, err := json.Marshal(models.NewAdvertTaskPayload{AdvertID: advertID})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("json marshal"), err)
		return err
	}
	task := asynq.NewTask(models.TypeNewAdvert, payload)

	_, err = c.client.Enqueue(task)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("client Enqueue"), err)
		return err
	}

	return nil
}

func (c *Client) AddNewReview(reviewID int) error {
	payload, err := json.Marshal(models.NewReviewTaskPayload{ReviewID: reviewID})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("json marshal"), err)
		return err
	}
	task := asynq.NewTask(models.TypeNewReview, payload)

	_, err = c.client.Enqueue(task)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("client Enqueue"), err)
		return err
	}

	return nil
}

func (c *Client) AddSubscribeUser(followerUserID int, followedUserID int) error {
	payload, err := json.Marshal(models.SubscribeUserTaskPayload{
		FollowerUserID: followerUserID,
		FollowedUserID: followedUserID,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("json marshal"), err)
		return err
	}
	task := asynq.NewTask(models.TypeSubscribeUser, payload)

	_, err = c.client.Enqueue(task)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("client Enqueue"), err)
		return err
	}

	return nil
}

func (c *Client) AddSubscribePlace(userID int, placeID int) error {
	payload, err := json.Marshal(models.SubscribePlaceTaskPayload{
		UserID:  userID,
		PlaceID: placeID,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("json marshal"), err)
		return err
	}
	task := asynq.NewTask(models.TypeSubscribePlace, payload)

	_, err = c.client.Enqueue(task)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("client Enqueue"), err)
		return err
	}

	return nil
}
