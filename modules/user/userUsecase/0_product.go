package usecase

import (
	userDto "shopping-service-be/modules/user/userDto"
	repository "shopping-service-be/modules/user/userRepository"
)

type (
	UserUsecaseService interface {
		CreateUser(createUserReq *userDto.CreateUserReq) (*userDto.CreateUserRes, error)
	}

	userUsecase struct {
		userRepo repository.UserRepositoryService
	}
)

func NewUserUsecase(userRepo repository.UserRepositoryService) UserUsecaseService {
	return &userUsecase{userRepo: userRepo}
}
