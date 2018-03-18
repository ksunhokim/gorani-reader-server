package auth

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/markbates/goth"
	"github.com/sirupsen/logrus"
	"github.com/sunho/engbreaker/pkg/config"
	"github.com/sunho/engbreaker/pkg/models"
)

const CookieName = "_eng_token"

var jwtSecret []byte

func init() {
	jwtSecret = []byte(config.GetString(config.JWTSECRET))
}

func CreateToken(email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   time.Now().Add(time.Minute * 60).Unix(),
		"email": email,
	})
	tokenstring, _ := token.SignedString(jwtSecret)
	return tokenstring
}

func GetTokenOrRegister(user goth.User) string {
	_, err := models.GetUser(user.Email)
	if err == nil {
		return CreateToken(user.Email)
	}

	err = models.AddUser(models.User{
		Username: user.NickName,
		Email:    user.Email,
	})
	if err != nil {
		logrus.Error(err)
	}

	return CreateToken(user.Email)
}

func ParseToken(tokenString string) (models.User, error) {
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
