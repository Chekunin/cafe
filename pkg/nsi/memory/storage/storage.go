package storage

import "cafe/pkg/models"

type Storage interface {
	GetAllPlaces() ([]models.Place, error)
}
