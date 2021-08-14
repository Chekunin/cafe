package usecase

import (
	"cafe/pkg/models"
	"context"
	"fmt"
	"github.com/Chekunin/wraperr"
)

func (u *Usecase) SubscribeToUser(ctx context.Context, followerUserID int, followedUserID int) error {
	userSubscription := models.UserSubscription{
		FollowerUserID: followerUserID,
		FollowedUserID: followedUserID,
	}

	if err := u.dbManager.AddUserSubscription(ctx, &userSubscription); err != nil {
		err = wraperr.NewWrapErr(fmt.Errorf("dbManager AddUserSubscription userSubscription = %+v", userSubscription), err)
		// todo: если уже подписка такая есть, то обрабатывать подписку
		return err
	}

	return nil
}

func (u *Usecase) UnsubscribeFromUser(ctx context.Context, followerUserID int, followedUserID int) error {
	userSubscription := models.UserSubscription{
		FollowerUserID: followerUserID,
		FollowedUserID: followedUserID,
	}

	if err := u.dbManager.DeleteUserSubscription(ctx, userSubscription); err != nil {
		err = wraperr.NewWrapErr(fmt.Errorf("dbManager DeleteUserSubscription userSubscription = %+v", userSubscription), err)
		return err
	}

	return nil
}