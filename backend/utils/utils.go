package utils

import (
	"github.com/Harsh-apk/jwtTest/types"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 12

func EncryptPassword(password *string) (*string, error) {
	encPw, err := bcrypt.GenerateFromPassword([]byte(*password), bcryptCost)
	if err != nil {
		return nil, err
	}
	ret := string(encPw)
	return &ret, nil
}

func ComparePassword(password *string, encPw *string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(*encPw), []byte(*password))
	return err == nil
}
func CreateUserFromIncomingUser(inUser *types.IncomingUser) (*types.User, error) {
	encPw, err := EncryptPassword(&inUser.Password)
	if err != nil {
		return nil, err
	}
	return &types.User{
		UserName: inUser.UserName,
		Email:    inUser.Email,
		EncPw:    *encPw,
	}, nil
}
