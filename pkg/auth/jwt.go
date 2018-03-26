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
	_, err := model.GetUser(user.Provider, user.UserID)
	if err == nil {
		return CreateToken(user)
	}

	err = models.AddUser(model.User{
		AuthType: user.Provider,
		AuthID:   user.UserID,
		Nickname: user.NickName,
		Email:    user.Email,
	})

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
		return models.User{}, err
	}
	if !token.Valid {
		return models.User{}, fmt.Errorf("not valid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return models.User{}, fmt.Errorf("not valid token")
	}

	email, ok := claims["email"]
	if !ok {
		return models.User{}, fmt.Errorf("not valid token")
	}

	user, err := models.GetUser(email.(string))
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
