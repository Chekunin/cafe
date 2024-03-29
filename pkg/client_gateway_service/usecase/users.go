package usecase

import (
	"cafe/pkg/common/catcherr"
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

	if err := u.feedQueueClient.AddSubscribeUser(followerUserID, followedUserID); err != nil {
		err = wraperr.NewWrapErr(fmt.Errorf("feedQueueClient AddSubscribeUser followerUserID=%d followedUserID=%d", followerUserID, followedUserID), err)
		catcherr.AsWarning().Catch(err)
	}

	// todo: здесь надо сообщать всем остальным (например, nsi), что произошло изменение состояние,
	//  можно сообщать через шину nats.

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

	userFeed := models.UserFeed{
		UserID:         followerUserID,
		FollowedUserID: followedUserID,
	}
	if err := u.dbManager.DeleteUsersFeeds(ctx, userFeed); err != nil {
		err = wraperr.NewWrapErr(fmt.Errorf("dbManager DeleteUsersFeeds userFeed=%+v", userFeed), err)
		return err
	}

	return nil
}

func (u *Usecase) GetUserSubscriptionsByFollowerID(ctx context.Context, followerUserID int) ([]models.UserSubscription, error) {
	userSubscriptions, err := u.nsi.GetUserSubscriptionsByFollowerID(ctx, followerUserID)
	if err != nil {
		err = wraperr.NewWrapErr(fmt.Errorf("nsi GetUserSubscriptionsByFollowerID followerUserID=%d", followerUserID), err)
		return nil, err
	}

	return userSubscriptions, nil
}

func (u *Usecase) GetFeedOfUserID(ctx context.Context, userID int, lastUserFeedID int, limit int) ([]models.UserFeed, error) {
	feed, err := u.dbManager.GetUsersFeedOfUserID(ctx, userID, lastUserFeedID, limit)
	if err != nil {
		err = wraperr.NewWrapErr(fmt.Errorf("dbManager GetUsersFeedOfUserID userID=%d lastUserFeedID=%d limit=%d", userID, lastUserFeedID, limit), err)
		return nil, err
	}

	return feed, nil
}
