package http

import (
	"cafe/pkg/common"
	"cafe/pkg/db_manager_service/api/delivery/rest/schema"
	"cafe/pkg/models"
	"cafe/pkg/transport/http"
	"context"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"io"
	"strconv"
)

type HttpDbManager struct {
	httpClient *http.HttpClient
}

// todo: не забыть делать маппинг ошибок на локальные,
//  а то будем отдавать пользователю ошибки внутренних сервисов

func NewHttpDbManager(uri string) (*HttpDbManager, error) {
	httpClient := http.NewHttpClient(http.HttpClientParams{
		BaseUrl: uri,
		Headers: map[string]string{},
		ErrorHandler: func(reader io.Reader) error {
			var err common.Err
			if err2 := http.GobDecoder(reader, &err); err2 != nil {
				err2 = wrapErr.NewWrapErr(fmt.Errorf("http GobDecoder"), err2)
				return err2
			}
			return common.NewErr(err.Code, err.Message, err.Meta)
		},
		RequestPayloadEncoder: http.GobEncoder,
		RequestPayloadDecoder: http.GobDecoder,
	})

	return &HttpDbManager{httpClient: httpClient}, nil
}

func (h HttpDbManager) GetAllPlaces(ctx context.Context) ([]models.Place, error) {
	var resp []models.Place
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/places",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllPlaceMedias(ctx context.Context) ([]models.PlaceMedia, error) {
	var resp []models.PlaceMedia
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/place-medias",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	var resp []models.Category
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/categories",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllPlaceCategories(ctx context.Context) ([]models.PlaceCategory, error) {
	var resp []models.PlaceCategory
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/place-categories",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllKitchenCategories(ctx context.Context) ([]models.KitchenCategory, error) {
	var resp []models.KitchenCategory
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/kitchen-categories",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllPlaceKitchenCategories(ctx context.Context) ([]models.PlaceKitchenCategory, error) {
	var resp []models.PlaceKitchenCategory
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/place-kitchen-categories",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllPlaceSchedules(ctx context.Context) ([]models.PlaceSchedule, error) {
	var resp []models.PlaceSchedule
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/place-schedules",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllAdverts(ctx context.Context) ([]models.Advert, error) {
	var resp []models.Advert
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/adverts",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllAdvertMedias(ctx context.Context) ([]models.AdvertMedia, error) {
	var resp []models.AdvertMedia
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/advert-medias",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllEvaluationCriterions(ctx context.Context) ([]models.EvaluationCriterion, error) {
	var resp []models.EvaluationCriterion
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/evaluation-criterions",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllPlaceEvaluations(ctx context.Context) ([]models.PlaceEvaluation, error) {
	var resp []models.PlaceEvaluation
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/place-evaluations",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) AddPlaceEvaluationWithMarks(ctx context.Context, placeEvaluation *models.PlaceEvaluation, marks []models.PlaceEvaluationMark) error {
	payload := schema.ReqEvaluatePlace{
		PlaceEvaluation: *placeEvaluation,
		Marks:           marks,
	}
	var resp schema.ReqEvaluatePlace
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:     ctx,
		Method:  "POST",
		Url:     "/place-evaluation-with-marks",
		Payload: payload,
		Result:  &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	*placeEvaluation = resp.PlaceEvaluation
	for i, _ := range marks {
		marks[i] = resp.Marks[i]
	}
	return nil
}

func (h HttpDbManager) GetAllPlaceEvaluationMarks(ctx context.Context) ([]models.PlaceEvaluationMark, error) {
	var resp []models.PlaceEvaluationMark
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/place-evaluation-marks",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllReviews(ctx context.Context) ([]models.Review, error) {
	var resp []models.Review
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/reviews",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetReviewsByUserID(ctx context.Context, userID int, lastReviewID int, limit int) ([]models.Review, error) {
	var resp []models.Review
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    fmt.Sprintf("/reviews-by-user-id/%d", userID),
		UrlParams: map[string]string{
			"last_review_id": strconv.Itoa(lastReviewID),
			"limit":          strconv.Itoa(limit),
		},
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetReviewByID(ctx context.Context, reviewID int) (models.Review, error) {
	var resp models.Review
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    fmt.Sprintf("/review-by-id/%d", reviewID),
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return models.Review{}, err
	}
	return resp, nil
}

func (h HttpDbManager) AddReview(ctx context.Context, review *models.Review) error {
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:     ctx,
		Method:  "POST",
		Url:     "/review",
		Payload: *review,
		Result:  review,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpDbManager) AddReviewMedia(ctx context.Context, reviewMedia *models.ReviewMedia) error {
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:     ctx,
		Method:  "POST",
		Url:     "/review-media",
		Payload: *reviewMedia,
		Result:  reviewMedia,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpDbManager) GetAllReviewMedias(ctx context.Context) ([]models.ReviewMedia, error) {
	var resp []models.ReviewMedia
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/review-medias",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllReviewReviewMedias(ctx context.Context) ([]models.ReviewReviewMedias, error) {
	var resp []models.ReviewReviewMedias
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/review-review-medias",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAdvertsByPlaceID(ctx context.Context, placeID int, lastAdvertID int, limit int) ([]models.Advert, error) {
	var resp []models.Advert
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    fmt.Sprintf("/adverts-by-place-id/%d", placeID),
		UrlParams: map[string]string{
			"last_advert_id": strconv.Itoa(lastAdvertID),
			"limit":          strconv.Itoa(limit),
		},
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAdvertByID(ctx context.Context, advertID int) (models.Advert, error) {
	var resp models.Advert
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    fmt.Sprintf("/advert-by-id/%d", advertID),
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return models.Advert{}, err
	}
	return resp, nil
}

func (h HttpDbManager) GetUserPlaceSubscriptionsByPlaceID(ctx context.Context, placeID int) ([]models.UserPlaceSubscription, error) {
	var resp []models.UserPlaceSubscription
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    fmt.Sprintf("/users-places-by-place-id/%d", placeID),
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetUsersPlacesSubscriptionsByUserID(ctx context.Context, userID int) ([]models.UserPlaceSubscription, error) {
	var resp []models.UserPlaceSubscription
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    fmt.Sprintf("/users-places-by-user-id/%d", userID),
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) AddReviewReviewMedias(ctx context.Context, reviewReviewMedias []models.ReviewReviewMedias) error {
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:     ctx,
		Method:  "POST",
		Url:     "/review-review-medias",
		Payload: reviewReviewMedias,
		Result:  &reviewReviewMedias,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpDbManager) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var resp []models.User
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/users",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetUserByUserID(ctx context.Context, userID int) (models.User, error) {
	var resp models.User
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    fmt.Sprintf("/user-by-id/%d", userID),
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return models.User{}, err
	}
	return resp, nil
}

func (h HttpDbManager) GetUserByUserName(ctx context.Context, userName string) (models.User, error) {
	var resp models.User
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    fmt.Sprintf("/user-by-name/%s", userName),
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return models.User{}, err
	}
	return resp, nil
}

func (h HttpDbManager) GetUserByVerifiedPhone(ctx context.Context, phone string) (models.User, error) {
	var resp models.User
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    fmt.Sprintf("/user-by-phone/%s", phone),
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return models.User{}, err
	}
	return resp, nil
}

func (h HttpDbManager) CreateUser(ctx context.Context, user *models.User) error {
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:     ctx,
		Method:  "POST",
		Url:     "/user",
		Result:  user,
		Payload: *user,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpDbManager) UpdateUser(ctx context.Context, user *models.User) error {
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:     ctx,
		Method:  "POST",
		Url:     fmt.Sprintf("/user/%d", user.ID),
		Result:  user,
		Payload: *user,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpDbManager) GetAllUserSubscriptions(ctx context.Context) ([]models.UserSubscription, error) {
	var resp []models.UserSubscription
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/user-subscriptions",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) AddUserSubscription(ctx context.Context, userSubscription *models.UserSubscription) error {
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:     ctx,
		Method:  "POST",
		Url:     "/user-subscription",
		Payload: *userSubscription,
		Result:  userSubscription,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpDbManager) DeleteUserSubscription(ctx context.Context, userSubscription models.UserSubscription) error {
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:     ctx,
		Method:  "DELETE",
		Url:     "/user-subscription",
		Payload: userSubscription,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpDbManager) GetUserSubscriptionsByFollowedUserID(ctx context.Context, followedUserID int) ([]models.UserSubscription, error) {
	var resp []models.UserSubscription
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    fmt.Sprintf("/users-subscriptions-by-followed-user-id/%d", followedUserID),
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllPlaceSubscriptions(ctx context.Context) ([]models.UserPlaceSubscription, error) {
	var resp []models.UserPlaceSubscription
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/place-subscriptions",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) AddPlaceSubscription(ctx context.Context, userPlaceSubscription *models.UserPlaceSubscription) error {
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:     ctx,
		Method:  "POST",
		Url:     "/place-subscription",
		Payload: *userPlaceSubscription,
		Result:  userPlaceSubscription,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpDbManager) DeletePlaceSubscription(ctx context.Context, userPlaceSubscription models.UserPlaceSubscription) error {
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:     ctx,
		Method:  "DELETE",
		Url:     "/place-subscription",
		Payload: userPlaceSubscription,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpDbManager) GetActualUserPhoneCodeByUserID(ctx context.Context, userID int) (models.UserPhoneCode, error) {
	var resp models.UserPhoneCode
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    fmt.Sprintf("/actual-user-phone-code-by-user-id/%d", userID),
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return models.UserPhoneCode{}, err
	}
	return resp, nil
}

func (h HttpDbManager) CreateUserPhoneCode(ctx context.Context, userPhoneCode *models.UserPhoneCode) error {
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:     ctx,
		Method:  "POST",
		Url:     "/user-phone-code",
		Payload: *userPhoneCode,
		Result:  userPhoneCode,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpDbManager) UpdateUserPhoneCode(ctx context.Context, userPhoneCode *models.UserPhoneCode) error {
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:     ctx,
		Method:  "POST",
		Url:     fmt.Sprintf("/user-phone-code/%d", userPhoneCode.ID),
		Payload: *userPhoneCode,
		Result:  userPhoneCode,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpDbManager) ActivateUserPhone(ctx context.Context, userPhoneCodeID int) error {
	payload := schema.ReqActivateUserPhone{UserPhoneCodeID: userPhoneCodeID}
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:     ctx,
		Method:  "POST",
		Url:     "/activate-user-phone",
		Payload: payload,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpDbManager) GetUsersFeedOfUserID(ctx context.Context, userID int, lastUserFeedID int, limit int) ([]models.UserFeed, error) {
	var resp []models.UserFeed
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    fmt.Sprintf("/feed-by-user-id/%d", userID),
		UrlParams: map[string]string{
			"last_user_feed_id": strconv.Itoa(lastUserFeedID),
			"limit":             strconv.Itoa(limit),
		},
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) AddUsersFeed(ctx context.Context, usersFeed []models.UserFeed) error {
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:     ctx,
		Method:  "POST",
		Url:     fmt.Sprintf("/users-feeds"),
		Payload: usersFeed,
		Result:  &usersFeed,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpDbManager) DeleteUsersFeeds(ctx context.Context, userFeeds models.UserFeed) error {
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:     ctx,
		Method:  "DELETE",
		Url:     fmt.Sprintf("/users-feeds"),
		Payload: userFeeds,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpDbManager) AddFeedAdvertQueue(ctx context.Context, feedAdvertQueue *models.FeedAdvertQueue) error {
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:     ctx,
		Method:  "POST",
		Url:     "/add-feed-advert-queue",
		Payload: *feedAdvertQueue,
		Result:  feedAdvertQueue,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpDbManager) PollFeedAdvertQueue(ctx context.Context) (models.FeedAdvertQueue, error) {
	var resp models.FeedAdvertQueue
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "POST",
		Url:    "/poll-feed-advert-queue",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return models.FeedAdvertQueue{}, err
	}
	return resp, nil
}

func (h HttpDbManager) CompleteFeedAdvertQueue(ctx context.Context, advertID int) error {
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "POST",
		Url:    fmt.Sprintf("/complete-feed-advert-queue/%d", advertID),
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpDbManager) AddFeedReviewQueue(ctx context.Context, feedReviewQueue *models.FeedReviewQueue) error {
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:     ctx,
		Method:  "POST",
		Url:     "/add-feed-review-queue",
		Payload: *feedReviewQueue,
		Result:  feedReviewQueue,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpDbManager) PollFeedReviewQueue(ctx context.Context) (models.FeedReviewQueue, error) {
	var resp models.FeedReviewQueue
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "POST",
		Url:    "/poll-feed-review-queue",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return models.FeedReviewQueue{}, err
	}
	return resp, nil
}

func (h HttpDbManager) CompleteFeedReviewQueue(ctx context.Context, reviewID int) error {
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "POST",
		Url:    fmt.Sprintf("/complete-feed-review-queue/%d", reviewID),
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpDbManager) AddFeedUserSubscribeQueue(ctx context.Context, feedUserSubscribeQueue *models.FeedUserSubscribeQueue) error {
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:     ctx,
		Method:  "POST",
		Url:     "/add-feed-user-subscribe-queue",
		Payload: *feedUserSubscribeQueue,
		Result:  feedUserSubscribeQueue,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpDbManager) PollFeedUserSubscribeQueue(ctx context.Context) (models.FeedUserSubscribeQueue, error) {
	var resp models.FeedUserSubscribeQueue
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "POST",
		Url:    "/poll-feed-user-subscribe-queue",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return models.FeedUserSubscribeQueue{}, err
	}
	return resp, nil
}

func (h HttpDbManager) CompleteFeedUserSubscribeQueue(ctx context.Context, followerUserID int, followedUserID int) error {
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "POST",
		Url:    fmt.Sprintf("/complete-feed-user-subscribe-queue/%d/%d", followerUserID, followedUserID),
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpDbManager) GetFullPlaceMenu(ctx context.Context, placeID int) (models.PlaceMenu, error) {
	var resp models.PlaceMenu
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    fmt.Sprintf("/places/%d/menu", placeID),
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return models.PlaceMenu{}, err
	}
	return resp, nil
}
