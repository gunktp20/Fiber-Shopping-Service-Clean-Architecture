package repository

import (
	userDto "shopping-service-be/modules/user/userDto"
	userEntity "shopping-service-be/modules/user/userEntity"
)

type (
	UserRepositoryService interface {
		IsUniqueUser(email string) bool
		CreateOneUser(createUserReq *userDto.CreateUserReq) (*userDto.CreateUserRes, error)
		GetOneUserByEmail(email string) (*userEntity.User, error)
	}
)
