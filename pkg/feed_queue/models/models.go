package models

// A list of task types.
const (
	TypeNewAdvert = "feed:advert"
	TypeNewReview = "feed:review"
)

type NewAdvertTaskPayload struct {
	AdvertID int
}

type NewReviewTaskPayload struct {
	ReviewID int
}
