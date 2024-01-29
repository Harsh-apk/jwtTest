package utils

import (
	"time"

	"github.com/Harsh-apk/jwtTest/types"
	"github.com/golang-jwt/jwt/v5"
)

const JWTSECRET = "HarshKaWebAppHai"

func CreateToken(user *types.User) (*string, error) {
	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(JWTSECRET))
	if err != nil {
		return nil, err
	}
	return &t, nil
}
