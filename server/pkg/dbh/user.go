package dbh

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader/server/pkg/auth"
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

func GetUser(db *gorm.DB, id int) (User, error) {
	out := User{}
	if err := db.
		First(&out, id).
		Error; err != nil {
		return out, err
	}
	return out, nil
}

func CreateOrGetUserWithOauth(db *gorm.DB, user auth.User) (_ User, err error) {
	passport := OauthPassport{}

	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	result := tx.
		Raw(`SELECT
				* 
			FROM
				oauth_passport
			WHERE
				oauth_service = ? AND
				oauth_user_id = ?
			LOCK IN SHARE MODE;`,
			user.Service, user.Id).
		Scan(&passport)

	if result.RecordNotFound() {
		return createUser(tx, user)
	}

	if err = result.Error; err != nil {
		return User{}, err
	}

	return GetUser(tx, passport.UserId)
}

func createUser(db *gorm.DB, user auth.User) (User, error) {
	newUser := User{
		Name: user.Username,
	}
	if err := db.Create(&newUser).Error; err != nil {
		return User{}, err
	}

	newUserDetail := UserDetail{
		Id:           newUser.Id,
		ProfileImage: user.Avator,
		AddedDate:    time.Now().UTC(),
	}
	if err := db.Create(&newUserDetail).Error; err != nil {
		return User{}, err
	}

	newPassport := OauthPassport{
		Service:     user.Service,
		UserId:      newUser.Id,
		OauthUserId: user.Id,
	}
	if err := db.Create(&newPassport).Error; err != nil {
		return User{}, err
	}

	return newUser, nil
}
