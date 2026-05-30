package repository

import (
	"spotify-clone/models"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
	
}

func NewUserRepository (db *gorm.DB) UserRepositoryInterface {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepositoryImpl) GetByEmail(email string) (*models.User, error) { 
	var user models.User
	if err := r.db.Where("email = ? ", email).First(&user).Error; err != nil{
		return nil, err
	}
	return &user, nil
}