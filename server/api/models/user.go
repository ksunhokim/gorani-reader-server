package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader/server/api/auth"
)

type User struct {
	Id           int       `gorm:"column:user_id"`
	Name         string    `gorm:"column:name"`
	ProfileImage string    `gorm:"column:user_profile_image"`
	AddedDate    time.Time `gorm:"column:user_added_date"`
}

func (User) TableName() string {
	return "user_with_detail"
}

func GetUser(db *gorm.DB, id int) (User, error) {
	out := User{}
	if err := db.
		Where("user_id = ?", id).
		First(&out).
		Error; err != nil {
		return out, err
	}
	return out, nil
}

func CreateOrGetUserByOauth(db *gorm.DB, user auth.User) (_ User, err error) {
	passport := OauthPassport{}

	code, err := GetOauthServiceCodeByName(db, user.Service)
	if err != nil {
		return User{}, err
	}

	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	if result := tx.
		Raw(`SELECT
				* 
			FROM
			oauth_passport
			WHERE
				oauth_service_code = ? AND
				oauth_user_id = ?
			LOCK IN SHARE MODE;`,
			code, user.Id).
		First(&passport); result.RecordNotFound() {
		return createUser(tx, user)
	} else if err = result.Error; err != nil {
		return User{}, err
	} else {
		return GetUser(tx, passport.UserId)
	}
}

func createUser(db *gorm.DB, user auth.User) (User, error) {
	code, err := GetOauthServiceCodeByName(db, user.Service)
	if err != nil {
		return User{}, err
	}

	newUser := User{
		Name:         user.Username,
		ProfileImage: user.Avator,
		AddedDate:    time.Now().UTC(),
	}
	if err = db.Create(&newUser).Error; err != nil {
		return User{}, err
	}

	newPassport := OauthPassport{
		Code:        code,
		UserId:      newUser.Id,
		OauthUserId: user.Id,
	}
	if err = db.Create(&newPassport).Error; err != nil {
		return User{}, err
	}

	return newUser, nil
}
