package model

import (
	"time"

	"github.com/sunho/gorani-reader/pkg/config"
	"github.com/sunho/gorani-reader/pkg/dbs"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id           bson.ObjectId  `bson:"_id,omitempty"`
	Email        string         `bson:"email"`
	Nickname     string         `bson:"nickname"`
	AuthProvider string         `bson:"auth_provider"`
	AuthID       string         `bson:"auth_id"`
	Unkown       *bson.ObjectId `json:"-"`
}

func GetUser(provider string, id string) (User, error) {
	sess := dbs.MDB.Copy()
	defer sess.Close()

	out := User{}
	err := sess.DB("").C("users").Find(bson.M{
		"auth_provider": provider,
		"auth_id":       id,
	}).One(&out)

	return out, err
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
	out := []Wordbook{}
	_ = sess.DB("").C("wordbooks").
		Find(
			bson.M{
				"user_id": user.Id,
			}).
		Sort("-updated_at").
		Skip(maxentries * page).
		Limit(maxentries).
		All(&out)

	return out
}

func (user *User) GetWordbook(index int) (Wordbook, error) {
	sess := dbs.MDB.Copy()
	defer sess.Close()

	out := Wordbook{}
	err := sess.DB("").C("wordbooks").
		Find(bson.M{
			"user_id": user.Id,
			"index":   index,
		}).
		One(&out)

	return out, err
}

func (user *User) CreateWordbook(name string) error {
	sess := dbs.MDB.Copy()
	defer sess.Close()

	// indice are generated by counter.go
	wordbook := Wordbook{
		UpdatedAt: time.Now(),
		UserId:    user.Id,
		Index:     NextSeq(user.Id.String() + "wordbookseq"),
		Name:      name,
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
