package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/sunho/engbreaker/pkg/api"
	"github.com/sunho/engbreaker/pkg/models"
)

func main() {
	db, err := sqlx.Connect("mysql", "eng_dev:eng_dev@(localhost)/eng_dev?parseTime=true")
	if err != nil {
		fmt.Println(err)
	}
	u, _ := models.GetUser(db, 1)
	books, _ := u.GetWordBooks(db)
	fmt.Println(books)
	book := models.WordBook{
		Name: "asfdsdf",
	}
	err = u.AddWordBook(db, book)
	if err != nil {
		fmt.Println(err)
	}
	server := api.NewHTTPServer()
	err = server.Start()
	if err != nil {
		logrus.Panic(err)
	}
}
