package model

import (
	"github.com/sunho/engbreaker/pkg/config"
	"github.com/sunho/engbreaker/pkg/dbs"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id           bson.ObjectId  `bson:"_id,omitempty"`
	Email        string         `json:"email" bson:"email"`
	Nickname     string         `json:"nickname" bson:"nickname"`
	AuthProvider string         `json:"auth_provider" bson:"auth_provider"`
	AuthID       string         `json:"auth_id" bson:"auth_id"`
	Unkown       *bson.ObjectId `json:"-"`
}

func GetUser(provider string, id string) (User, error) {
	sess := dbs.MDB.Copy()
	defer sess.Close()

	user_ := User{}
	err := sess.DB("").C("users").Find(bson.M{
		"auth_provider": provider,
		"auth_id":       id,
	}).One(&user_)

	return user_, err
}

func CreateUser(user User) (User, error) {
	sess := dbs.MDB.Copy()
	defer sess.Close()

	user.Id = bson.NewObjectId()
	err := sess.DB("").C("users").Insert(&user)

	return user, err
}

func (user *User) GetWordbooks(page int) []Wordbook {
	sess := dbs.MDB.Copy()
	defer sess.Close()

	maxentries := config.GetInt("USER_MAX", 10)
	iter := sess.DB("").C("wordbooks").
		Find(
			bson.M{
				"user_id": user.Id,
			}).
		Sort("updated_at").
		Skip(maxentries * page).
		Limit(maxentries).
		Iter()

	books := []Wordbook{}
	wordbook := Wordbook{}
	for iter.Next(&wordbook) {
		books = append(books, wordbook)
	}

	return books
}

func (user *User) GetWordbook(index int) (Wordbook, error) {
	sess := dbs.MDB.Copy()
	defer sess.Close()

	wordbook := Wordbook{}
	err := sess.DB("").C("wordbooks").
		Find(bson.M{
			"user_id": user.Id,
			"index":   index,
		}).
		One(&wordbook)

	return wordbook, err
}

func (user *User) CreateWordbook(name string) error {
	sess := dbs.MDB.Copy()
	defer sess.Close()

	wordbook := Wordbook{
		UserId: user.Id,
		Index:  NextSeq(user.Id.String() + "wordbookseq"),
		Name:   name,
	}
	err := sess.DB("").C("wordbooks").Insert(&wordbook)
	if err != nil {
		return err
	}

	return nil
}

func (user *User) DeleteWordbook(index int) error {
	sess := dbs.MDB.Copy()
	defer sess.Close()

	err := sess.DB("").C("wordbooks").Remove(bson.M{
		"user_id": user.Id,
		"index":   index,
	})

	return err
}
