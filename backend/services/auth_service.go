package services

import (
	"errors"
	"spotify-clone/repository"
	"spotify-clone/models"
	"spotify-clone/utils"
	)

type AuthService interface { 
	Login(email,password string) (*models.User, error)
	Register(email, password string) (*models.User, error)
}

type authService struct {
	userRepo repository.UserRepositoryInterface
}

func NewAuthService(userRepo repository.UserRepositoryInterface) AuthService{
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) Login(email, password string) (*models.User, error){
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("salah email")
	}

	isValidPassword := utils.CheckPasswordHash(password, user.HashedPassword)
	if !isValidPassword { 
		return nil, errors.New("salah password")
	}
	return user, nil
}

func (s *authService) Register(email, password string) (*models.User, error){
	//ngecek email udah dipake belom
	existingUser, err := s.userRepo.GetByEmail(email)
	if err == nil && existingUser != nil{
		return nil, errors.New("email udah ada")
	}

	// ngehash sebelum save ke database
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, errors.New("gagal hashing")
	}

	// buat user
	newUser := &models.User{
		Email: email,
		HashedPassword: hashedPassword,
	}
}



