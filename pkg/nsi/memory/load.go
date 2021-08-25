package memory

import (
	log "cafe/pkg/common/logman"
	"cafe/pkg/models"
	"context"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/kyroy/kdtree"
	"github.com/kyroy/kdtree/points"
	"sort"
)

// todo: сделать load для всех сущностей

func (n *NSI) Load(ctx context.Context) error {
	n.context = &NsiContext{}

	loaders := []func(ctx context.Context) error{
		n.loadPlaces,
		n.loadPlaceMedias,
		n.loadCategories,
		n.loadPlaceCategories,
		n.loadKitchenCategories,
		n.loadPlaceKitchenCategories,
		n.loadPlaceSchedules,
		n.loadAdverts,
		n.loadAdvertMedias,
		n.loadEvaluationCriterions,
		n.loadPlaceEvaluations,
		n.loadPlaceEvaluationMarks,
		n.loadReviews,
		n.loadReviewMedias,
		n.loadUsers,
		n.loadUserSubscriptions,
	}
	for i, loader := range loaders {
		log.Debug("Loading", log.Fields{"count": i * 100 / len(loaders)})
		if err := loader(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (n *NSI) loadPlaces(ctx context.Context) error {
	places, err := n.dbManager.GetAllPlaces(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllPlaces"), err)
		return err
	}

	n.context.places = places
	n.context.placesByID = map[int]*models.Place{}
	for i, v := range n.context.places {
		n.context.placesByID[v.ID] = &n.context.places[i]
	}

	if err := n.buildPlacesKDTreeByCoordinates(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("buildPlacesKDTreeByCoordinates"), err)
		return err
	}

	return nil
}

func (n *NSI) buildPlacesKDTreeByCoordinates() error {
	tree := kdtree.New([]kdtree.Point{})

	for i, place := range n.context.places {
		tree.Insert(points.NewPoint([]float64{place.Lat, place.Lng}, &n.context.places[i]))
	}

	tree.Balance()

	n.context.placesKDTreeByCoordinates = tree

	return nil
}

func (n *NSI) loadPlaceMedias(ctx context.Context) error {
	placeMedias, err := n.dbManager.GetAllPlaceMedias(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllPlaceMedias"), err)
		return err
	}

	n.context.placeMedias = placeMedias
	n.context.placeMediasByID = map[int]*models.PlaceMedia{}
	n.context.placeMediasByPlaceID = map[int][]*models.PlaceMedia{}
	for i, v := range n.context.placeMedias {
		n.context.placeMediasByID[v.ID] = &n.context.placeMedias[i]

		if _, has := n.context.placeMediasByPlaceID[v.PlaceID]; !has {
			n.context.placeMediasByPlaceID[v.PlaceID] = make([]*models.PlaceMedia, 0)
		}
		n.context.placeMediasByPlaceID[v.PlaceID] = append(n.context.placeMediasByPlaceID[v.PlaceID], &n.context.placeMedias[i])
	}
	return nil
}

func (n *NSI) loadCategories(ctx context.Context) error {
	categories, err := n.dbManager.GetAllCategories(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllCategories"), err)
		return err
	}

	n.context.categories = categories
	n.context.categoriesByID = map[int]*models.Category{}
	for i, v := range n.context.categories {
		n.context.categoriesByID[v.ID] = &n.context.categories[i]
	}
	return nil
}

func (n *NSI) loadPlaceCategories(ctx context.Context) error {
	placeCategories, err := n.dbManager.GetAllPlaceCategories(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllPlaceCategories"), err)
		return err
	}

	n.context.placeCategories = placeCategories

	if err := n.fillCategoriesByPlaceID(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("fillCategoriesByPlaceID"), err)
		return err
	}

	return nil
}

func (n *NSI) fillCategoriesByPlaceID() error {
	n.context.categoriesByPlaceID = map[int][]*models.Category{}
	for _, v := range n.context.placeCategories {
		if _, has := n.context.categoriesByPlaceID[v.PlaceID]; !has {
			n.context.categoriesByPlaceID[v.PlaceID] = make([]*models.Category, 0)
		}
		category, has := n.context.categoriesByID[v.CategoryID]
		if !has {
			err := wrapErr.NewWrapErr(fmt.Errorf("category with id=%d not found", v.CategoryID), nil)
			return err
		}
		n.context.categoriesByPlaceID[v.PlaceID] = append(n.context.categoriesByPlaceID[v.PlaceID], category)
	}

	return nil
}

func (n *NSI) loadKitchenCategories(ctx context.Context) error {
	kitchenCategories, err := n.dbManager.GetAllKitchenCategories(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllKitchenCategories"), err)
		return err
	}

	n.context.kitchenCategories = kitchenCategories
	n.context.kitchenCategoriesByID = map[int]*models.KitchenCategory{}
	for i, v := range n.context.kitchenCategories {
		n.context.kitchenCategoriesByID[v.ID] = &n.context.kitchenCategories[i]
	}

	return nil
}

func (n *NSI) loadPlaceKitchenCategories(ctx context.Context) error {
	placeKitchenCategories, err := n.dbManager.GetAllPlaceKitchenCategories(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllPlaceCategories"), err)
		return err
	}

	n.context.placeKitchenCategories = placeKitchenCategories

	if err := n.fillKitchenCategoriesByPlaceID(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("fillKitchenCategoriesByPlaceID"), err)
		return err
	}

	return nil
}

func (n *NSI) fillKitchenCategoriesByPlaceID() error {
	n.context.kitchenCategoriesByPlaceID = map[int][]*models.KitchenCategory{}
	for _, v := range n.context.placeKitchenCategories {
		if _, has := n.context.kitchenCategoriesByPlaceID[v.PlaceID]; !has {
			n.context.kitchenCategoriesByPlaceID[v.PlaceID] = make([]*models.KitchenCategory, 0)
		}
		kitchenCategory, has := n.context.kitchenCategoriesByID[v.KitchenCategoryID]
		if !has {
			err := wrapErr.NewWrapErr(fmt.Errorf("kitchenCategoryID with id=%d not found", v.KitchenCategoryID), nil)
			return err
		}
		n.context.kitchenCategoriesByPlaceID[v.PlaceID] = append(n.context.kitchenCategoriesByPlaceID[v.PlaceID], kitchenCategory)
	}

	return nil
}

func (n *NSI) loadPlaceSchedules(ctx context.Context) error {
	placeSchedules, err := n.dbManager.GetAllPlaceSchedules(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllPlaceSchedules"), err)
		return err
	}

	n.context.placeSchedules = placeSchedules
	n.context.placeSchedulesByPlaceID = map[int][]*models.PlaceSchedule{}
	for i, v := range n.context.placeSchedules {
		if _, has := n.context.placeSchedulesByPlaceID[v.PlaceID]; !has {
			n.context.placeSchedulesByPlaceID[v.PlaceID] = make([]*models.PlaceSchedule, 0)
		}
		n.context.placeSchedulesByPlaceID[v.PlaceID] = append(n.context.placeSchedulesByPlaceID[v.PlaceID], &n.context.placeSchedules[i])
	}

	return nil
}

func (n *NSI) loadAdverts(ctx context.Context) error {
	adverts, err := n.dbManager.GetAllAdverts(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllAdverts"), err)
		return err
	}

	n.context.adverts = adverts
	n.context.advertsByID = map[int]*models.Advert{}
	n.context.advertsByPlaceID = map[int][]*models.Advert{}
	for i, v := range n.context.adverts {
		n.context.advertsByID[v.ID] = &n.context.adverts[i]

		if _, has := n.context.advertsByPlaceID[v.PlaceID]; !has {
			n.context.advertsByPlaceID[v.PlaceID] = make([]*models.Advert, 0)
		}
		n.context.advertsByPlaceID[v.PlaceID] = append(n.context.advertsByPlaceID[v.PlaceID], &n.context.adverts[i])
	}

	return nil
}

func (n *NSI) loadAdvertMedias(ctx context.Context) error {
	advertMedias, err := n.dbManager.GetAllAdvertMedias(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllAdvertMedias"), err)
		return err
	}

	n.context.advertMedias = advertMedias
	n.context.advertMediasByID = map[int]*models.AdvertMedia{}
	for i, v := range n.context.advertMedias {
		n.context.advertMediasByID[v.ID] = &n.context.advertMedias[i]
	}

	return nil
}

func (n *NSI) loadEvaluationCriterions(ctx context.Context) error {
	evaluationCriterions, err := n.dbManager.GetAllEvaluationCriterions(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllEvaluationCriterions"), err)
		return err
	}

	n.context.evaluationCriterions = evaluationCriterions
	n.context.evaluationCriterionsByID = map[int]*models.EvaluationCriterion{}
	for i, v := range n.context.evaluationCriterions {
		n.context.evaluationCriterionsByID[v.ID] = &n.context.evaluationCriterions[i]
	}

	return nil
}

func (n *NSI) loadPlaceEvaluations(ctx context.Context) error {
	placeEvaluations, err := n.dbManager.GetAllPlaceEvaluations(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllPlaceEvaluations"), err)
		return err
	}

	n.context.placeEvaluations = placeEvaluations
	n.context.placeEvaluationsByID = map[int]*models.PlaceEvaluation{}
	n.context.placeEvaluationsByPlaceID = map[int][]*models.PlaceEvaluation{}
	n.context.placeEvaluationsByUserIDByPlaceID = map[int]map[int]*models.PlaceEvaluation{}
	for i, v := range n.context.placeEvaluations {
		n.context.placeEvaluationsByID[v.ID] = &n.context.placeEvaluations[i]

		if _, has := n.context.placeEvaluationsByPlaceID[v.PlaceID]; !has {
			n.context.placeEvaluationsByPlaceID[v.PlaceID] = make([]*models.PlaceEvaluation, 0)
		}
		n.context.placeEvaluationsByPlaceID[v.PlaceID] = append(n.context.placeEvaluationsByPlaceID[v.PlaceID], &n.context.placeEvaluations[i])

		if _, has := n.context.placeEvaluationsByUserIDByPlaceID[v.UserID]; !has {
			n.context.placeEvaluationsByUserIDByPlaceID[v.UserID] = make(map[int]*models.PlaceEvaluation, 0)
		}
		n.context.placeEvaluationsByUserIDByPlaceID[v.UserID][v.PlaceID] = &n.context.placeEvaluations[i]
	}

	return nil
}

func (n *NSI) loadPlaceEvaluationMarks(ctx context.Context) error {
	placeEvaluationMarks, err := n.dbManager.GetAllPlaceEvaluationMarks(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllPlaceEvaluationMarks"), err)
		return err
	}

	n.context.placeEvaluationMarks = placeEvaluationMarks
	n.context.placeEvaluationMarksByPlaceEvaluationID = map[int][]*models.PlaceEvaluationMark{}
	for i, v := range n.context.placeEvaluationMarks {
		if _, has := n.context.placeEvaluationMarksByPlaceEvaluationID[v.PlaceEvaluationID]; !has {
			n.context.placeEvaluationMarksByPlaceEvaluationID[v.PlaceEvaluationID] = make([]*models.PlaceEvaluationMark, 0)
		}
		n.context.placeEvaluationMarksByPlaceEvaluationID[v.PlaceEvaluationID] = append(n.context.placeEvaluationMarksByPlaceEvaluationID[v.PlaceEvaluationID], &n.context.placeEvaluationMarks[i])
	}

	return nil
}

func (n *NSI) loadReviews(ctx context.Context) error {
	reviews, err := n.dbManager.GetAllReviews(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllReviews"), err)
		return err
	}

	n.context.reviews = reviews
	n.context.reviewsByID = map[int]*models.Review{}
	n.context.reviewsByUserID = map[int][]*models.Review{}
	for i, v := range n.context.reviews {
		n.context.reviewsByID[v.ID] = &n.context.reviews[i]

		if _, has := n.context.reviewsByUserID[v.UserID]; !has {
			n.context.reviewsByUserID[v.UserID] = make([]*models.Review, 0)
		}
		n.context.reviewsByUserID[v.UserID] = append(n.context.reviewsByUserID[v.UserID], &n.context.reviews[i])
	}

	return nil
}

func (n *NSI) loadReviewMedias(ctx context.Context) error {
	reviewMedias, err := n.dbManager.GetAllReviewMedias(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllReviewMedias"), err)
		return err
	}

	n.context.reviewMedias = reviewMedias
	n.context.reviewMediasByID = map[int]*models.ReviewMedia{}
	n.context.reviewMediasByReviewID = map[int][]*models.ReviewMedia{}
	for i, v := range n.context.reviewMedias {
		n.context.reviewMediasByID[v.ID] = &n.context.reviewMedias[i]
	}

	// ------
	// todo: перепроверить оптимальность работы данного метода
	reviewReviewMedias, err := n.dbManager.GetAllReviewReviewMedias(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllReviewReviewMedias"), err)
		return err
	}

	reviewReviewMediasByReviewID := map[int][]*models.ReviewReviewMedias{}
	for i, v := range reviewReviewMedias {
		if _, has := reviewReviewMediasByReviewID[v.ReviewID]; !has {
			reviewReviewMediasByReviewID[v.ReviewID] = make([]*models.ReviewReviewMedias, 0)
		}
		reviewReviewMediasByReviewID[v.ReviewID] = append(reviewReviewMediasByReviewID[v.ReviewID], &reviewReviewMedias[i])
	}
	for reviewID, v := range reviewReviewMediasByReviewID {
		sort.Slice(reviewReviewMediasByReviewID[reviewID], func(i, j int) bool {
			return reviewReviewMediasByReviewID[reviewID][i].Order < reviewReviewMediasByReviewID[reviewID][j].Order
		})

		if _, has := n.context.reviewMediasByReviewID[reviewID]; !has {
			n.context.reviewMediasByReviewID[reviewID] = make([]*models.ReviewMedia, 0)
		}
		for _, v2 := range v {
			n.context.reviewMediasByReviewID[reviewID] = append(n.context.reviewMediasByReviewID[reviewID], n.context.reviewMediasByID[v2.ReviewMediaID])
		}
	}

	return nil
}

func (n *NSI) loadUsers(ctx context.Context) error {
	users, err := n.dbManager.GetAllUsers(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllUsers"), err)
		return err
	}

	n.context.users = users
	n.context.usersByID = map[int]*models.User{}
	for i, v := range n.context.reviewMedias {
		n.context.usersByID[v.ID] = &n.context.users[i]
	}

	return nil
}

func (n *NSI) loadUserSubscriptions(ctx context.Context) error {
	userSubscriptions, err := n.dbManager.GetAllUserSubscriptions(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetAllUserSubscriptions"), err)
		return err
	}

	n.context.userSubscriptions = userSubscriptions
	n.context.userSubscriptionsByFollowerUserID = map[int][]*models.UserSubscription{}
	n.context.userSubscriptionsByFollowedUserID = map[int][]*models.UserSubscription{}
	for i, v := range n.context.userSubscriptions {
		if _, has := n.context.userSubscriptionsByFollowerUserID[v.FollowerUserID]; !has {
			n.context.userSubscriptionsByFollowerUserID[v.FollowerUserID] = make([]*models.UserSubscription, 0)
		}
		n.context.userSubscriptionsByFollowerUserID[v.FollowerUserID] = append(n.context.userSubscriptionsByFollowerUserID[v.FollowerUserID], &n.context.userSubscriptions[i])

		if _, has := n.context.userSubscriptionsByFollowedUserID[v.FollowedUserID]; !has {
			n.context.userSubscriptionsByFollowedUserID[v.FollowedUserID] = make([]*models.UserSubscription, 0)
		}
		n.context.userSubscriptionsByFollowedUserID[v.FollowedUserID] = append(n.context.userSubscriptionsByFollowedUserID[v.FollowedUserID], &n.context.userSubscriptions[i])
	}

	return nil
}
