package memory

import (
	"cafe/pkg/models"
	"github.com/kyroy/kdtree"
)

type NsiContext struct {
	places                    []models.Place
	placesByID                map[int]*models.Place
	placesKDTreeByCoordinates *kdtree.KDTree

	placeMedias          []models.PlaceMedia
	placeMediasByID      map[int]*models.PlaceMedia
	placeMediasByPlaceID map[int][]*models.PlaceMedia

	categories          []models.Category
	categoriesByID      map[int]*models.Category
	categoriesByPlaceID map[int][]*models.Category

	placeCategories []models.PlaceCategory

	kitchenCategories          []models.KitchenCategory
	kitchenCategoriesByID      map[int]*models.KitchenCategory
	kitchenCategoriesByPlaceID map[int][]*models.KitchenCategory

	placeKitchenCategories []models.PlaceKitchenCategory

	placeSchedules          []models.PlaceSchedule
	placeSchedulesByPlaceID map[int][]*models.PlaceSchedule

	adverts          []models.Advert
	advertsByID      map[int]*models.Advert
	advertsByPlaceID map[int][]*models.Advert

	advertMedias     []models.AdvertMedia
	advertMediasByID map[int]*models.AdvertMedia

	evaluationCriterions     []models.EvaluationCriterion
	evaluationCriterionsByID map[int]*models.EvaluationCriterion

	placeEvaluations          []models.PlaceEvaluation
	placeEvaluationsByID      map[int]*models.PlaceEvaluation
	placeEvaluationsByPlaceID map[int][]*models.PlaceEvaluation
	//placeEvaluationsByUserID  map[int][]*models.PlaceEvaluation
	placeEvaluationsByUserIDByPlaceID map[int]map[int]*models.PlaceEvaluation

	placeEvaluationMarks                    []models.PlaceEvaluationMark
	placeEvaluationMarksByPlaceEvaluationID map[int][]*models.PlaceEvaluationMark

	reviews         []models.Review
	reviewsByID     map[int]*models.Review
	reviewsByUserID map[int][]*models.Review

	reviewMedias           []models.ReviewMedia
	reviewMediasByID       map[int]*models.ReviewMedia
	reviewMediasByReviewID map[int][]*models.ReviewMedia

	users     []models.User
	usersByID map[int]*models.User

	userSubscriptions                 []models.UserSubscription
	userSubscriptionsByFollowerUserID map[int][]*models.UserSubscription
	userSubscriptionsByFollowedUserID map[int][]*models.UserSubscription
}
