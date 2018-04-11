package auth

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/markbates/goth"
	"github.com/sirupsen/logrus"
	"github.com/sunho/engbreaker/pkg/config"
	"github.com/sunho/engbreaker/pkg/model"
)

var jwtSecret []byte

func init() {
	jwtSecret = []byte(config.GetString("JWT_SECRET", "ASDF"))
}

func CreateToken(user goth.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().Add(time.Minute * 60).Unix(),
		"id":       user.UserID,
		"provider": user.Provider,
	})
	tokenstring, _ := token.SignedString(jwtSecret)

	return tokenstring
}

func GetTokenOrRegister(user goth.User) string {
	user_, err := model.GetUser(user.Provider, user.UserID)
	if err == nil {
		return CreateToken(user)
	}

	user_ = model.User{
		Nickname:     user.NickName,
		Email:        user.Email,
		AuthProvider: user.Provider,
		AuthID:       user.UserID,
	}
	_, err = model.CreateUser(user_)
	if err != nil {
		logrus.Error(err)
	}

	return CreateToken(user)
}

func ParseToken(tokenString string) (model.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return model.User{}, err
	}
	if !token.Valid {
		return model.User{}, fmt.Errorf("Not valid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return model.User{}, fmt.Errorf("Not valid token")
	}

	provider, ok := claims["provider"].(string)
	if !ok {
		return model.User{}, fmt.Errorf("Not valid token")
	}

	id, ok := claims["id"].(string)
	if !ok {
		return model.User{}, fmt.Errorf("Not valid token")
	}

	user, err := model.GetUser(provider, id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
