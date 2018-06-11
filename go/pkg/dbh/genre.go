package dbh

import "github.com/jinzhu/gorm"

type Genre struct {
	Code int    `gorm:"column:genre_code;primary_key"`
	Name string `gorm:"column:genre_name"`
}

func (Genre) TableName() string {
	return "genre"
}

type UserPreferGenre struct {
	UserId int `gorm:"column:user_id"`
	Code   int `gorm:"column:genre_code"`
}

func (UserPreferGenre) TableName() string {
	return "user_prefer_genre"
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

func (u *User) GetPreferGenres(db *gorm.DB) (genres []Genre, err error) {
	err = db.Raw(`
			SELECT genre.*
			FROM
				user_prefer_genre ug
			INNER JOIN
				genre
			ON
				genre.genre_code = ug.genre_code
			WHERE
				ug.user_id = ?;`, u.Id).
		Scan(&genres).Error
	return
}

func (u *User) PutPreferGenres(db *gorm.DB, genres []Genre) (err error) {
	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	err = tx.
		Where("user_id = ?", u.Id).
		Delete(&UserPreferGenre{}).Error
	if err != nil {
		return
	}

	for _, genre := range genres {
		ug := UserPreferGenre{
			UserId: u.Id,
			Code:   genre.Code,
		}
		err = tx.Create(&ug).Error
		if err != nil {
			return
		}
	}
	return
}
