package main

import "github.com/sunho/engbreaker/pkg/models"

func main() {
	user := models.User{Username: "hello, world"}
	models.DB.NewRecord(user)
}
