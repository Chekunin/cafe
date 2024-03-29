package models

import (
	"time"
)

//MediaType type
type MediaType string

//MediaType enum
const (
	MediaTypePhoto MediaType = "photo"
	MediaTypeVideo MediaType = "video"
)

//Mark type
type Mark string

//Mark enum
const (
	MarkExcellent Mark = "excellent"
	MarkGood      Mark = "good"
	MarkBad       Mark = "bad"
)

type Place struct {
	tableName struct{} `pg:"main.places,discard_unknown_columns"`

	ID          int     `pg:"place_id,pk" json:"place_id" api:"place_id"`
	Name        string  `pg:"name" json:"name" api:"name"`
	Description string  `pg:"description" json:"description" api:"description"`
	Lat         float64 `pg:"lat,notnull" json:"lat" api:"lat"`
	Lng         float64 `pg:"lng,notnull"  json:"lng" api:"lng"`
	Address     string  `pg:"address" json:"address" api:"address"`
	Website     string  `pg:"website" json:"website" api:"website"`
	Rating      float64 `pg:"rating" json:"rating" api:"rating"`
	MarkAmount  int     `pg:"mark_amount" json:"mark_amount" api:"mark_amount"`

	Categories        []Category        `pg:"-" json:"categories" api:"categories"`
	KitchenCategories []KitchenCategory `pg:"-" json:"kitchen_categories" api:"kitchen_categories"`
	PlaceMedias       []PlaceMedia      `pg:"-" json:"place_medias" api:"place_medias"`
	PlaceSchedules    []PlaceSchedule   `pg:"-" json:"place_schedules" api:"place_schedules"`
}

type PlaceMedia struct {
	tableName struct{} `pg:"main.place_media,discard_unknown_columns"`

	ID              int       `pg:"place_media_id,pk" json:"place_media_id" api:"place_media_id"`
	PlaceID         int       `pg:"place_id" json:"place_id" api:"place_id"`
	MediaPath       string    `pg:"media_path" json:"media_path" api:"media_path"`
	MediaType       MediaType `pg:"media_type" json:"media_type" api:"media_type" enums:"photo,video"`
	Comment         string    `pg:"comment" json:"comment" api:"comment"`
	PublishDateTime time.Time `pg:"publish_datetime" json:"publish_datetime" api:"publish_datetime"`
}

type Category struct {
	tableName struct{} `pg:"main.categories,discard_unknown_columns"`

	ID   int    `pg:"category_id,pk" json:"category_id" api:"category_id"`
	Name string `pg:"name" json:"name" api:"name"`
}

type PlaceCategory struct {
	tableName struct{} `pg:"main.places_categories,discard_unknown_columns"`

	PlaceID    int `pg:"place_id,pk" json:"place_id" api:"place_id"`
	CategoryID int `pg:"category_id,pk" json:"category_id" api:"category_id"`
}

type KitchenCategory struct {
	tableName struct{} `pg:"main.kitchen_categories,discard_unknown_columns"`

	ID   int    `pg:"kitchen_category_id" json:"kitchen_category_id" api:"kitchen_category_id"`
	Name string `pg:"name" json:"name" api:"name"`
}

type PlaceKitchenCategory struct {
	tableName struct{} `pg:"main.places_kitchen_categories,discard_unknown_columns"`

	PlaceID           int `pg:"place_id,pk" json:"place_id" api:"place_id"`
	KitchenCategoryID int `pg:"kitchen_category_id,pk" json:"kitchen_category_id" api:"kitchen_category_id"`
}

type PlaceSchedule struct {
	tableName struct{} `pg:"main.places_schedules,discard_unknown_columns"`

	ID        int       `pg:"place_schedule_id,pk" json:"place_schedule_id" api:"place_schedule_id"`
	PlaceID   int       `pg:"place_id" json:"place_id" api:"place_id"`
	DayOfWeek int       `pg:"day_of_week" json:"day_of_week" api:"day_of_week"`
	StartTime time.Time `pg:"start_time" json:"start_time" api:"start_time"`
	EndTime   time.Time `pg:"end_time" json:"end_time" api:"end_time"`
	DateStart time.Time `pg:"date_start" json:"date_start" api:"date_start"`
	DateStop  time.Time `pg:"date_stop" json:"date_stop" api:"date_stop"`
}

type Advert struct {
	tableName struct{} `pg:"main.adverts,discard_unknown_columns"`

	ID              int       `pg:"advert_id,pk" json:"advert_id" api:"advert_id"`
	PlaceID         int       `pg:"place_id" json:"place_id" api:"place_id"`
	RestaurateurID  int       `pg:"restaurateur_id" json:"restaurateur_id"`
	Text            string    `pg:"text" json:"text" api:"text"`
	PublishDateTime time.Time `pg:"publish_datetime" json:"publish_datetime" api:"publish_datetime"`

	AdvertMedias []AdvertMedia `pg:"many2many:main.advert_advert_medias" json:"advert_medias" api:"advert_medias"`
}

type AdvertMedia struct {
	tableName struct{} `pg:"main.advert_medias,discard_unknown_columns"`

	ID             int       `pg:"advert_media_id,pk" json:"advert_media_id" api:"advert_media_id"`
	PlaceID        int       `pg:"place_id" json:"place_id" api:"place_id"`
	RestaurateurID int       `pg:"restaurateur_id" json:"restaurateur_id"`
	MediaType      MediaType `pg:"media_type" json:"media_type" api:"media_type" enums:"photo,video"`
	MediaPath      string    `pg:"media_path" json:"media_path" api:"media_path"`
}

type AdvertAdvertMedias struct {
	tableName struct{} `pg:"main.advert_advert_medias,discard_unknown_columns"`

	AdvertID      int `pg:"advert_id,pk" json:"advert_id"`
	AdvertMediaID int `pg:"advert_media_id,pk" json:"advert_media_id"`
	Order         int `pg:"order" json:"order"`
}

type EvaluationCriterion struct {
	tableName struct{} `pg:"main.evaluation_criterions,discard_unknown_columns"`

	ID   int    `pg:"evaluation_criterion_id,pk" json:"evaluation_criterion_id" api:"evaluation_criterion_id"`
	Name string `pg:"name" json:"name" api:"name"`
}

type PlaceEvaluation struct {
	tableName struct{} `pg:"main.places_evaluations,discard_unknown_columns"`

	ID       int       `pg:"place_evaluation_id,pk" json:"place_evaluation_id" api:"place_evaluation_id"`
	PlaceID  int       `pg:"place_id" json:"place_id" api:"place_id"`
	UserID   int       `pg:"user_id" json:"user_id" api:"user_id"`
	DateTime time.Time `pg:"datetime" json:"datetime" api:"datetime"`
	Comment  string    `pg:"comment" json:"comment" api:"comment"`

	PlaceEvaluationMarks []PlaceEvaluationMark `pg:"-" json:"place_evaluation_marks" api:"place_evaluation_marks"`
}

type PlaceEvaluationMark struct {
	tableName struct{} `pg:"main.place_evaluation_marks,discard_unknown_columns"`

	PlaceEvaluationID     int  `pg:"place_evaluation_id" json:"-" api:"place_evaluation_id"`
	EvaluationCriterionID int  `pg:"evaluation_criterion_id" json:"evaluation_criterion_id" api:"evaluation_criterion_id"`
	Mark                  Mark `pg:"mark" json:"mark" api:"mark"`
}

type Review struct {
	tableName struct{} `pg:"main.reviews,discard_unknown_columns"`

	ID              int       `pg:"review_id,pk" json:"review_id" api:"review_id"`
	UserID          int       `pg:"user_id" json:"user_id" api:"user_id"`
	PlaceID         int       `pg:"place_id" json:"place_id" api:"place_id"`
	Text            string    `pg:"text" json:"text" api:"text"`
	PublishDateTime time.Time `pg:"publish_datetime" json:"publish_datetime" api:"publish_datetime"`

	ReviewMedias []ReviewMedia `pg:"many2many:main.review_review_medias" json:"review_medias" api:"review_medias"`
}

type ReviewMedia struct {
	tableName struct{} `pg:"main.review_medias,discard_unknown_columns"`

	ID        int       `pg:"review_media_id,pk" json:"review_media_id" api:"review_media_id"`
	UserID    int       `pg:"user_id" json:"user_id"`
	MediaType MediaType `pg:"media_type" json:"media_type" api:"media_type" enums:"photo,video"`
	MediaPath string    `pg:"media_path" json:"media_path"`
}

type ReviewReviewMedias struct {
	tableName struct{} `pg:"main.review_review_medias,discard_unknown_columns"`

	ReviewID      int `pg:"review_id,pk" json:"review_id"`
	ReviewMediaID int `pg:"review_media_id,pk" json:"review_media_id"`
	Order         int `pg:"order" json:"order"`
}

type User struct {
	tableName struct{} `pg:"main.users,discard_unknown_columns"`

	ID            int       `pg:"user_id,pk" json:"user_id" api:"user_id"`
	Name          string    `pg:"name" json:"name" api:"name"`
	Phone         string    `pg:"phone" json:"phone" api:"phone"`
	Email         string    `pg:"email" json:"email" api:"email"`
	Password      string    `pg:"password" json:"password"`
	PhoneVerified bool      `pg:"phone_verified,use_zero" json:"phone_verified" api:"phone_verified"`
	EmailVerified bool      `pg:"email_verified,use_zero" json:"email_verified" api:"email_verified"`
	RegDateTime   time.Time `pg:"reg_datetime" json:"reg_datetime" api:"reg_datetime"`
	PhotoPath     string    `pg:"photo_path" json:"photo_path" api:"photo_path"`
}

type UserSubscription struct {
	tableName struct{} `pg:"main.users_subscriptions,discard_unknown_columns"`

	FollowerUserID int `pg:"follower_user_id,pk" json:"follower_user_id" api:"follower_user_id"`
	FollowedUserID int `pg:"followed_user_id,pk" json:"followed_user_id" api:"followed_user_id"`
}

type UserPhoneCode struct {
	tableName struct{} `pg:"main.user_phone_codes,discard_unknown_columns"`

	ID             int       `pg:"user_phone_code_id,pk" json:"user_phone_code_id"`
	UserID         int       `pg:"user_id" json:"user_id"`
	Code           string    `pg:"code" json:"code"`
	CreateDatetime time.Time `pg:"create_datetime" json:"create_datetime"`
	Actual         bool      `pg:"actual,use_zero" json:"actual"`
	LeftAttempts   int       `pg:"left_attempts" json:"left_attempts"`
}

type Restaurateur struct {
	tableName struct{} `pg:"main.restaurateurs,discard_unknown_columns"`

	ID            int       `pg:"restaurateur_id,pk" json:"restaurateur_id"`
	Email         string    `pg:"email" json:"email"`
	Password      string    `pg:"password" json:"password"`
	EmailVerified bool      `pg:"email_verified,use_zero" json:"email_verified"`
	RegDateTime   time.Time `pg:"reg_datetime" json:"reg_datetime"`
}

type RestaurateurRole struct {
	tableName struct{} `pg:"main.restaurateur_roles,discard_unknown_columns"`

	ID   int    `pg:"restaurateur_role_id,pk" json:"restaurateur_role_id"`
	Name string `pg:"name" json:"name"`
}

type PlaceRestaurateur struct {
	tableName struct{} `pg:"main.places_restaurateurs,discard_unknown_columns"`

	RestaurateurID int `pg:"restaurateur_id,pk" json:"restaurateur_id"`
	PlaceID        int `pg:"place_id,pk" json:"place_id"`
	RoleID         int `pg:"restaurateur_role_id" json:"restaurateur_role_id"`
}

type UserFeed struct {
	tableName struct{} `pg:"main.users_feeds,discard_unknown_columns"`

	ID              int       `pg:"users_feed_id,pk" json:"restaurateur_id" api:"restaurateur_id"`
	UserID          int       `pg:"user_id" json:"user_id" api:"user_id"`
	AdvertID        int       `pg:"advert_id" json:"advert_id" api:"advert_id"`
	ReviewID        int       `pg:"review_id" json:"review_id" api:"review_id"`
	PublishDatetime time.Time `pg:"publish_datetime" json:"publish_datetime" api:"publish_datetime"`
	FollowedUserID  int       `pg:"followed_user_id" json:"followed_user_id"`
	PlaceID         int       `pg:"place_id" json:"place_id"`

	Advert *Advert `pg:"fk:advert_id" api:"advert"`
	Review *Review `pg:"fk:review_id" api:"review"`
}

type FeedAdvertQueue struct {
	tableName struct{} `pg:"main.feed_advert_queue,discard_unknown_columns"`

	AdvertID             int       `pg:"advert_id,pk" json:"advert_id"`
	Status               int       `pg:"status" json:"status"`
	ChangeStatusDatetime time.Time `pg:"change_status_datetime" json:"change_status_datetime"`
}

type FeedReviewQueue struct {
	tableName struct{} `pg:"main.feed_review_queue,discard_unknown_columns"`

	ReviewID             int       `pg:"review_id,pk" json:"review_id"`
	Status               int       `pg:"status" json:"status"`
	ChangeStatusDatetime time.Time `pg:"change_status_datetime" json:"change_status_datetime"`
}

type FeedUserSubscribeQueue struct {
	tableName struct{} `pg:"main.feed_user_subscribe_queue,discard_unknown_columns"`

	FollowerUserID       int       `pg:"follower_user_id,pk" json:"follower_user_id"`
	FollowedUserID       int       `pg:"followed_user_id,pk" json:"followed_user_id"`
	Status               int       `pg:"status" json:"status"`
	ChangeStatusDatetime time.Time `pg:"change_status_datetime" json:"change_status_datetime"`
}

type UserPlaceSubscription struct {
	tableName struct{} `pg:"main.users_places_subscriptions,discard_unknown_columns"`

	UserID  int `pg:"user_id,pk" json:"user_id"`
	PlaceID int `pg:"place_id,pk" json:"place_id"`
}

type PlaceMenuCategory struct {
	tableName struct{} `pg:"main.place_menu_categories,discard_unknown_columns"`

	ID              int       `pg:"place_menu_category_id,pk" json:"place_menu_category_id" api:"place_menu_category_id"`
	PlaceID         int       `pg:"place_id" json:"place_id" api:"place_id"`
	CategoryName    string    `pg:"category_name" json:"category_name" api:"category_name"`
	Order           int       `pg:"order" json:"order" api:"order"`
	LoadDateTime    time.Time `pg:"load_datetime" json:"load_datetime" api:"load_datetime"`
	PublishDateTime time.Time `pg:"publish_datetime" json:"publish_datetime" api:"publish_datetime"`

	PlaceMenuItems []PlaceMenuItem `pg:"rel:has-many" json:"place_menu_items" api:"place_menu_items"`
}

type PlaceMenuItemMedia struct {
	tableName struct{} `pg:"main.place_menu_item_medias,discard_unknown_columns"`

	ID             int       `pg:"place_menu_item_media_id,pk" json:"place_menu_item_media_id" api:"place_menu_item_media_id"`
	RestaurateurID int       `pg:"restaurateur_id" json:"restaurateur_id" api:"restaurateur_id"`
	MediaPath      string    `pg:"media_path" json:"media_path" api:"media_path"`
	MediaType      MediaType `pg:"media_type" json:"media_type" api:"media_type" enums:"photo,video"`
	LoadDateTime   time.Time `pg:"load_datetime" json:"load_datetime" api:"load_datetime"`
}

type PlaceMenuItem struct {
	tableName struct{} `pg:"main.place_menu_items,discard_unknown_columns"`

	ID                   int       `pg:"place_menu_item_id,pk" json:"place_menu_item_id" api:"place_menu_item_id"`
	Name                 string    `pg:"name" json:"name" api:"name"`
	Description          string    `pg:"description" json:"description" api:"description"`
	Weight               int       `pg:"weight" json:"weight" api:"weight"`
	Price                int       `pg:"price" json:"price" api:"price"`
	PlaceMenuItemMediaID int       `pg:"place_menu_item_media_id" json:"place_menu_item_media_id" api:"place_menu_item_media_id"`
	PlaceMenuCategoryID  int       `pg:"place_menu_category_id" json:"place_menu_category_id" api:"place_menu_category_id"`
	Order                int       `pg:"order" json:"order" api:"order"`
	LoadDateTime         time.Time `pg:"load_datetime" json:"load_datetime" api:"load_datetime"`
	PublishDateTime      time.Time `pg:"publish_datetime" json:"publish_datetime" api:"publish_datetime"`

	PlaceMenuItemMedia *PlaceMenuItemMedia `pg:"fk:place_menu_item_media_id" json:"place_menu_item_media" api:"place_menu_item_media"`
}

type PlaceMenu struct {
	PlaceID             int                 `json:"place_id" api:"place_id"`
	PlaceMenuCategories []PlaceMenuCategory `json:"place_menu_categories" api:"place_menu_categories"`
}
