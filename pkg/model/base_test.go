package model_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/engbreaker/pkg/dbs"
	"github.com/sunho/engbreaker/pkg/model"
	"gopkg.in/mgo.v2/bson"
)

func initDB() {
	os.Setenv("MONGO_DB", "test")
	dbs.Init()
	dbs.MDB.Collection("users").Delete(bson.M{})
	user := model.User{
		Nickname: "test",
	}
	dbs.MDB.Collection("users").Save(&user)
	user = model.User{
		Nickname: "test",
	}
	dbs.MDB.Collection("users").Save(&user)
	user = model.User{
		Nickname: "test2",
	}
	dbs.MDB.Collection("users").Save(&user)
}

func TestGetOneFail(t *testing.T) {
	initDB()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("This should panic")
		}
	}()

	user := model.User{}
	model.Get(user, bson.M{"nickname": "test2"})
}

func TestGetOne(t *testing.T) {
	initDB()
	a := assert.New(t)

	user := model.User{}
	err := model.Get(&user, bson.M{"nickname": "test2"})
	a.Equal(nil, err)
	a.Equal("test2", user.Nickname)
	err = model.Get(&user, bson.M{"nickname": "test"})
	a.Equal(nil, err)
	err = model.Get(&user, bson.M{"nickname": "sklasjgwr"})
	a.NotEqual(nil, err)
}

func TestGetSliceFail(t *testing.T) {
	initDB()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("This should panic")
		}
	}()

	users := []model.User{}
	model.Get(&users, bson.M{"nickname": "test"})
}

func TestGetSlice(t *testing.T) {
	initDB()
	a := assert.New(t)

	users := []*model.User{}
	model.Get(&users, bson.M{"nickname": "test"})
	a.Equal(len(users), 2)
	a.Equal("test", users[1].Nickname)
	model.Get(&users, bson.M{"nickname": "sklasjgwr"})
	a.Equal(len(users), 0)
}

func TestSave(t *testing.T) {
	initDB()
	a := assert.New(t)

	user := model.User{
		Nickname: "test3",
	}
	model.Save(&user)
	user2 := model.User{}
	err := model.Get(&user2, bson.M{"nickname": "test3"})
	a.Equal(nil, err)
	a.Equal("test3", user.Nickname)
}

func TestSaveSlice(t *testing.T) {
	initDB()
	a := assert.New(t)

	user := []*model.User{
		&model.User{
			Nickname: "test4",
		},
		&model.User{
			Nickname: "test4",
		},
	}
	err := model.Save(&user)
	a.Equal(nil, err)
	user2 := []*model.User{}
	model.Get(&user2, bson.M{"nickname": "test4"})
	a.Equal(2, len(user2))
	a.Equal("test4", user[0].Nickname)
}
