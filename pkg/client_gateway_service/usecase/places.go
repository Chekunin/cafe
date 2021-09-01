package usecase

import (
	"cafe/pkg/common/catcherr"
	"cafe/pkg/models"
	"context"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/google/uuid"
	"io"
	"time"
)

func (u *Usecase) EvaluatePlace(ctx context.Context, placeID int, userID int, marks []models.PlaceEvaluationMark, comment string) (models.PlaceEvaluation, []models.PlaceEvaluationMark, error) {
	placeEvaluation := models.PlaceEvaluation{
		PlaceID:  placeID,
		UserID:   userID,
		DateTime: time.Now(),
		Comment:  comment,
	}
	if err := u.dbManager.AddPlaceEvaluationWithMarks(ctx, &placeEvaluation, marks); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager AddPlaceEvaluationWithMarks"), err)
		return models.PlaceEvaluation{}, nil, err
	}

	// todo: отдавать в nsi

	return placeEvaluation, marks, nil
}

func (u *Usecase) GetPlaceEvaluation(ctx context.Context, placeID int, userID int) (models.PlaceEvaluation, []models.PlaceEvaluationMark, error) {
	placeEvaluation, err := u.nsi.GetPlaceEvaluationByUserIDByPlaceID(ctx, userID, placeID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("nsi GetPlaceEvaluationByUserIDByPlaceID userID=%d, placeID=%d", userID, placeID), err)
		return models.PlaceEvaluation{}, nil, err
	}

	marks, err := u.nsi.GetPlaceEvaluationMarksByPlaceEvaluationID(ctx, placeEvaluation.ID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("nsi GetPlaceEvaluationMarksByPlaceEvaluationID placeEvaluationID=%d", placeEvaluation.ID), err)
		return models.PlaceEvaluation{}, nil, err
	}

	return placeEvaluation, marks, nil
}

func (u *Usecase) AddPlaceReviewMedia(ctx context.Context, userID int, reader io.Reader) (models.ReviewMedia, error) {
	mediaPath := fmt.Sprintf("/%d/%s", userID, uuid.New().String())
	object, err := u.reviewMediaStorage.Put(ctx, mediaPath, reader)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("reviewMediaStorage put path=%s", mediaPath), err)
		return models.ReviewMedia{}, err
	}
	mediaPath = object.Path

	//firstBytes := make([]byte, 512)
	//reader.Read(firstBytes)
	//fileType := http.DetectContentType(firstBytes)

	mediaType := models.MediaTypePhoto
	// todo: здесь определять тип

	reviewMedia := models.ReviewMedia{
		UserID:    userID,
		MediaType: mediaType,
		MediaPath: mediaPath,
	}

	if err := u.dbManager.AddReviewMedia(ctx, &reviewMedia); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager AddReviewMedia reviewMedia=%+v", reviewMedia), err)
		return models.ReviewMedia{}, err
	}

	return reviewMedia, nil
}

func (u *Usecase) GetReviewMediaData(ctx context.Context, reviewMediaID int) (io.ReadCloser, string, error) {
	reviewMedia, err := u.nsi.GetReviewMediaByID(ctx, reviewMediaID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("nsi GetReviewMediaByID reviewMediaID=%d", reviewMediaID), err)
		return nil, "", err
	}

	readCloser, contentType, err := u.reviewMediaStorage.GetStream(ctx, reviewMedia.MediaPath)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("reviewMediaStorage GetStream mediaPath=%s", reviewMedia.MediaPath), err)
		return nil, "", err
	}
	return readCloser, contentType, nil
}

func (u *Usecase) AddPlaceReview(ctx context.Context, userID int, placeID int, text string, reviewMediaIDs []int) (models.Review, error) {
	// создаём новое ревью в БД
	review := models.Review{
		UserID:          userID,
		PlaceID:         placeID,
		Text:            text,
		PublishDateTime: time.Now(),
	}
	if err := u.dbManager.AddReview(ctx, &review); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager AddReview review=%+v", review), err)
		return models.Review{}, err
	}

	// создаём связки review с media в БД
	var reviewReviewMedias []models.ReviewReviewMedias
	for i, v := range reviewMediaIDs {
		reviewReviewMedias = append(reviewReviewMedias, models.ReviewReviewMedias{
			ReviewID:      review.ID,
			ReviewMediaID: v,
			Order:         i + 1,
		})
	}
	if err := u.dbManager.AddReviewReviewMedias(ctx, reviewReviewMedias); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager AddReviewReviewMedias"), err)
		return models.Review{}, err
	}

	// достаём из БД нужные review_media и кладём в объект review
	for _, v := range reviewMediaIDs {
		reviewMedia, err := u.nsi.GetReviewMediaByID(ctx, v)
		if err != nil {
			err = wrapErr.NewWrapErr(fmt.Errorf("nsi GetReviewMediaByID reviewMediaID=%d", v), err)
			return models.Review{}, err
		}
		review.ReviewMedias = append(review.ReviewMedias, reviewMedia)
	}

	if err := u.feedQueueClient.AddNewReview(review.ID); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("feedQueueClient AddNewReview reviewID=%d", review.ID), err)
		catcherr.AsWarning().Catch(err)
	}

	return review, nil
}

func (u *Usecase) GetPlaceEvaluationCriterions(ctx context.Context) ([]models.EvaluationCriterion, error) {
	evaluationCriterion, err := u.nsi.GetPlaceEvaluationCriterions(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("nsi GetPlaceEvaluationCriterions"), err)
		return nil, err
	}
	return evaluationCriterion, nil
}

func (u *Usecase) GetPlacesReviewsOfUserID(ctx context.Context, userID int, lastReviewID int, limit int) ([]models.Review, error) {
	reviews, err := u.dbManager.GetReviewsByUserID(ctx, userID, lastReviewID, limit)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetReviewsByUserID userID=%d lastReviewID=%d limit=%d", userID, lastReviewID, limit), err)
		return nil, err
	}

	return reviews, nil
}

func (u *Usecase) GetPlaceAdvertsByPlaceID(ctx context.Context, placeID int, lastAdvertID int, limit int) ([]models.Advert, error) {
	adverts, err := u.dbManager.GetAdvertsByPlaceID(ctx, placeID, lastAdvertID, limit)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAdvertsByPlaceID placeID=%d lastAdvertID=%d limit=%d", placeID, lastAdvertID, limit), err)
		return nil, err
	}

	return adverts, nil
}

func (u *Usecase) SubscribeToPlace(ctx context.Context, userID int, placeID int) error {
	userPlaceSubscription := models.UserPlaceSubscription{
		UserID:  userID,
		PlaceID: placeID,
	}

	if err := u.dbManager.AddPlaceSubscription(ctx, &userPlaceSubscription); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager AddPlaceSubscription userPlaceSubscription = %+v", userPlaceSubscription), err)
		// todo: если уже подписка такая есть, то обрабатывать подписку
		return err
	}

	// todo: здесь надо сообщать всем остальным (например, nsi), что произошло изменение состояние,
	//  можно сообщать через шину nats.

	return nil
}

func (u *Usecase) UnsubscribeFromPlace(ctx context.Context, userID int, placeID int) error {
	userPlaceSubscription := models.UserPlaceSubscription{
		UserID:  userID,
		PlaceID: placeID,
	}

	if err := u.dbManager.DeletePlaceSubscription(ctx, userPlaceSubscription); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager DeletePlaceSubscription userPlaceSubscription = %+v", userPlaceSubscription), err)
		return err
	}

	return nil
}

func (u *Usecase) GetPlaceSubscriptionsByUserID(ctx context.Context, userID int) ([]models.UserPlaceSubscription, error) {
	userSubscriptions, err := u.dbManager.GetUsersPlacesSubscriptionsByUserID(ctx, userID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetUsersPlacesSubscriptionsByUserID userID=%d", userID), err)
		return nil, err
	}

	return userSubscriptions, nil
}
