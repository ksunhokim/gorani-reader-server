package dbh

import "github.com/jinzhu/gorm"

type Genre struct {
	Code int    `gorm:"column:genre_code;primary_key"`
	Name string `gorm:"column:genre_name"`
}

func (Genre) TableName() string {
	return "genre"
}

func GetGenres(db *gorm.DB) (genres []Genre, err error) {
	err = db.Find(&genres).Error
	return
}

func GetGenreByCode(db *gorm.DB, code int) (genre Genre, err error) {
	err = db.First(&genre, code).Error
	return
}

func GetGenreByName(db *gorm.DB, name string) (genre Genre, err error) {
	err = db.Where("genre_name = ?", name).First(&genre).Error
	return
}

func AddGenre(db *gorm.DB, genre *Genre) error {
	err := db.Create(genre).Error
	return err
}
