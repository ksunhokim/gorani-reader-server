package dbh_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader-server/pkg/dbh"
	"github.com/sunho/gorani-reader-server/pkg/util"
)

func TestKnownWords(t *testing.T) {
	gorn := util.SetupTestGorani()
	a := assert.New(t)

	u, err := dbh.GetUser(gorn.Mysql, util.TestUserId)
	a.Nil(err)
	err = u.AddKnownWord(gorn.Mysql, 1)
	a.Nil(err)

	words, err := u.GetKnownWords(gorn.Mysql, 1)
	a.Nil(err)
	sol := []dbh.KnownWord{
		dbh.KnownWord{
			UserId: util.TestUserId,
			WordId: 1,
			Number: 1,
		},
	}
	a.Equal(sol, words)

	err = u.AddKnownWord(gorn.Mysql, 1)
	a.Nil(err)

	words, err = u.GetKnownWords(gorn.Mysql, 1)
	a.Nil(err)

	sol = []dbh.KnownWord{
		dbh.KnownWord{
			UserId: util.TestUserId,
			WordId: 1,
			Number: 2,
		},
	}
	a.Equal(sol, words)
}
