package usecase

import (
	"errors"
	"fmt"
	authDto "shopping-service-be/modules/auth/authDto"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (u *authUsecase) Authenticate(authenticationReq *authDto.AuthenticationReq) (*authDto.AuthenticationRes, error) {

	user, err := u.userRepo.GetOneUserByEmail(authenticationReq.Email)
	if err != nil {
		return &authDto.AuthenticationRes{}, err
	}

	if user.Password == "" {
		return nil, errors.New("stored password is invalid or empty")
	}

	if authenticationReq.Password == "" {
		return nil, errors.New("provided password is empty")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authenticationReq.Password)); err != nil {
		fmt.Printf("failed compare hash and password %s", err.Error())
		return &authDto.AuthenticationRes{}, errors.New("email or password is invalid")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	accessToken, err := token.SignedString([]byte(u.cfg.Jwt.AccessSecretKey))

	if err != nil {
		return &authDto.AuthenticationRes{}, err
	}

	return &authDto.AuthenticationRes{
		AccessToken: accessToken,
	}, nil
}
