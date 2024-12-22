package usecase

import (
	authDto "shopping-service-be/modules/auth/authDto"
	userRepository "shopping-service-be/modules/user/userRepository"
	"shopping-service-be/pkg/config"
)

type (
	AuthUsecaseService interface {
		Authenticate(authenticationReq *authDto.AuthenticationReq) (*authDto.AuthenticationRes, error)
	}

	authUsecase struct {
		userRepo userRepository.UserRepositoryService
		cfg      *config.Config
	}
)

func NewAuthUsecase(userRepo userRepository.UserRepositoryService, cfg *config.Config) AuthUsecaseService {
	return &authUsecase{
		userRepo: userRepo,
		cfg:      cfg,
	}
}
