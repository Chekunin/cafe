package storage

import "cafe/pkg/models"

type Storage interface {
	GetUserByUserID(userID string) (models.User, error)
}
