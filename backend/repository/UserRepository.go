package repository

import (
	"spotify-clone/models"
)

type UserRepositoryInterface interface {
	GetByEmail(email string) (*models.User, error);
	GetByIdentifier(identifier string) (*models.User, error)
	Create(user *models.User) error
}

