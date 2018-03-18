package dbs

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/sunho/engbreaker/pkg/config"
)

var MDB *sqlx.DB

func init() {
	url := config.GetString("MYSQL_URL", "engbreaker:engbreaker@/engbreaker")
	tdb, err := sqlx.Connect("mysql", fmt.Sprintf(`%s?parseTime=true`, url))
	if err != nil {
		logrus.Panic(err)
	}

	MDB = tdb
}
