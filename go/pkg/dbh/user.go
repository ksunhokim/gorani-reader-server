package dbh

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader-server/go/pkg/auth"
)

type OauthPassport struct {
	Service     string `gorm:"column:oauth_service"`
	UserId      int    `gorm:"column:user_id"`
	OauthUserId string `gorm:"column:oauth_user_id"`
}

func (OauthPassport) TableName() string {
	return "oauth_passport"
}

type User struct {
	Id   int    `gorm:"column:user_id;primary_key"`
	Name string `gorm:"column:user_name"`
}

func (User) TableName() string {
	return "user"
}

type UserDetail struct {
	Id           int       `gorm:"column:user_id;primary_key"`
	ProfileImage string    `gorm:"column:user_profile_image"`
	AddedDate    time.Time `gorm:"column:user_added_date"`
}

func (UserDetail) TableName() string {
	return "user_detail"
}

func GetUser(db *gorm.DB, id int) (user User, err error) {
	err = db.First(&user, id).Error
	return
}

func CreateOrGetUserWithOauth(db *gorm.DB, user auth.User) (_ User, err error) {
	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	passport := OauthPassport{}
	result := tx.
		Where("oauth_service = ? AND oauth_user_id = ?").
		Set("gorm:query_option", "LOCK IN SHARE MODE").
		First(&passport)

	if result.RecordNotFound() {
		return createUser(tx, user)
	}

	if err = result.Error; err != nil {
		return User{}, err
	}

	return GetUser(tx, passport.UserId)
}

func createUser(db *gorm.DB, user auth.User) (newUser User, err error) {
	newUser.Name = user.Username

	err = db.Create(&newUser).Error
	if err != nil {
		return
	}

	newUserDetail := UserDetail{
		Id:           newUser.Id,
		ProfileImage: user.Avator,
		AddedDate:    time.Now().UTC(),
	}
	err = db.Create(&newUserDetail).Error
	if err != nil {
		return
	}

	newPassport := OauthPassport{
		Service:     user.Service,
		UserId:      newUser.Id,
		OauthUserId: user.Id,
	}
	err = db.Create(&newPassport).Error
	if err != nil {
		return
	}

	return
}
