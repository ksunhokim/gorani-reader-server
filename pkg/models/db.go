package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/sunho/engbreaker/pkg/dbs"
)

var db *sqlx.DB

func init() {
	db = dbs.MDB
}
