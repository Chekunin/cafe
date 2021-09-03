package models

// A list of task types.
const (
	TypeNewAdvert      = "feed:advert"
	TypeNewReview      = "feed:review"
	TypeSubscribeUser  = "feed:subscribe-user"
	TypeSubscribePlace = "feed:subscribe-place"
)

type NewAdvertTaskPayload struct {
	AdvertID int
}

type NewReviewTaskPayload struct {
	ReviewID int
}

type SubscribeUserTaskPayload struct {
	FollowerUserID int
	FollowedUserID int
}

type SubscribePlaceTaskPayload struct {
	UserID  int
	PlaceID int
}
