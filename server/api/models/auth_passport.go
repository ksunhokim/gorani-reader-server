package models

import "github.com/jinzhu/gorm"

type OauthService struct {
	Code int    `gorm:"column:oauth_service_code"`
	Name string `gorm:"column:oauth_service_name"`
}

func GetOauthServiceNameByCode(db *gorm.DB, code int) (string, error) {
	service := OauthService{}
	if err := db.
		Where("oauth_service_code = ?", code).
		First(&service).
		Error; err != nil {
		return "", err
	}
	return service.Name, nil
}

func GetOauthServiceCodeByName(db *gorm.DB, name string) (int, error) {
	service := OauthService{}
	if err := db.
		Where("oauth_service_name = ?", name).
		First(&service).
		Error; err != nil {
		return 0, err
	}
	return service.Code, nil
}

type OauthPassport struct {
	Code        int    `gorm:"column:oauth_service_code"`
	UserId      int    `gorm:"column:user_id"`
	OauthUserId string `gorm:"column:oauth_user_id"`
}

func (OauthPassport) TableName() string {
	return "oauth_passport"
}
