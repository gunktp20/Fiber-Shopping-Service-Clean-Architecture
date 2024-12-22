package usecase

import (
	"errors"
	userDto "shopping-service-be/modules/user/userDto"

	"golang.org/x/crypto/bcrypt"
)

func (r *userUsecase) CreateUser(createUserReq *userDto.CreateUserReq) (*userDto.CreateUserRes, error) {

	// Hashing Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUserReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	createUserReq.Password = string(hashedPassword)

	return r.userRepo.CreateOneUser(createUserReq)
}
