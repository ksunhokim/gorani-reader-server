package dbh_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader/server/pkg/dbh"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

func TestGenre(t *testing.T) {
	gorn := util.SetupTestGorani()
	a := assert.New(t)
	genres, err := dbh.GetGenres(gorn.Mysql)
	a.Nil(err)
	a.Equal(1, len(genres))

	g := dbh.Genre{Name: "hoi"}
	err = dbh.AddGenre(gorn.Mysql, &g)
	a.Nil(err)

	genres, err = dbh.GetGenres(gorn.Mysql)
	a.Nil(err)
	a.Equal(2, len(genres))

	genre, err := dbh.GetGenreByCode(gorn.Mysql, g.Code)
	a.Nil(err)
	a.Equal(g, genre)

	genre, err = dbh.GetGenreByName(gorn.Mysql, "hoi")
	a.Nil(err)
	a.Equal(g, genre)
}

func TestUserPreferGenre(t *testing.T) {
	gorn := util.SetupTestGorani()
	a := assert.New(t)
	u, err := dbh.GetUser(gorn.Mysql, util.TestUserId)
	a.Nil(err)

	genres, err := u.GetPreferGenres(gorn.Mysql)
	a.Nil(err)
	a.Equal(0, len(genres))

	genres = append(genres, dbh.Genre{
		Code: 1,
		Name: "test",
	})

	err = u.PutPreferGenres(gorn.Mysql, genres)
	a.Nil(err)

	genres2, err := u.GetPreferGenres(gorn.Mysql)
	a.Nil(err)
	a.Equal(1, len(genres))

	a.Equal(genres, genres2)
}
