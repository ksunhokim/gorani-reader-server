package db

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/inflection"
	"github.com/sunho/engbreaker/pkg/config"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open("mysql", config.GetString(config.MYSQLURL))
	if err != nil {
		panic(err)
	}
	DB = db
	DB.SetLogger(&gormLogger{})
	if config.GetString(config.DEBUG) == "true" {
		DB.LogMode(true)
	}
	DB.AutoMigrate(&User{}, &Book{}, &BookPicture{}, &Word{}, &Def{}, &Example{})
	many2ManyFIndex(&User{}, &Book{})
	DB.Model(&BookPicture{}).AddForeignKey("book_id", "books(id)", "CASCADE", "RESTRICT")
	DB.Model(&Def{}).AddForeignKey("word_id", "words(id)", "CASCADE", "RESTRICT")
	DB.Model(&Example{}).AddForeignKey("def_id", "defs(id)", "CASCADE", "RESTRICT")
}

func reduceModelToName(model interface{}) string {
	value := reflect.ValueOf(model)
	if value.Kind() != reflect.Ptr {
		return ""
	}

	elem := value.Elem()
	t := elem.Type()
	rawName := t.Name()
	return gorm.ToDBName(rawName)
}

func many2ManyFIndex(parentModel interface{}, childModel interface{}) {
	table1Accessor := reduceModelToName(parentModel)
	table2Accessor := reduceModelToName(childModel)

	table1Name := inflection.Plural(table1Accessor)
	table2Name := inflection.Plural(table2Accessor)

	joinTable := fmt.Sprintf("%s_%s", table1Accessor, table2Name)
	DB.Table(joinTable).AddForeignKey(table1Accessor+"_id", table1Name+"(id)", "CASCADE", "CASCADE")
	DB.Table(joinTable).AddForeignKey(table2Accessor+"_id", table2Name+"(id)", "CASCADE", "CASCADE")
	DB.Table(joinTable).AddUniqueIndex(joinTable+"_unique", table1Accessor+"_id", table2Accessor+"_id")
}
