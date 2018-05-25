package models_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader/server/api/models"
	"github.com/sunho/gorani-reader/server/api/util"
)

func TestGetWordbook(t *testing.T) {
	gorn := Setup()
	a := assert.New(t)
	id, _ := uuid.Parse(TestWordbookUUID)
	bytes := util.UUIDToBytes(id)
	wordbook, err := models.GetWordbook(gorn.Mysql, bytes)
	if err != nil {
		panic(err)
	}

	a.Equal(wordbook.Name, "test")
	a.Equal(wordbook.IsUnknown, false)
}
