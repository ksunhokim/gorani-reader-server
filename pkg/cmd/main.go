package main

import (
	"github.com/sunho/engbreaker/pkg/config"
	"github.com/sunho/engbreaker/pkg/router"
)

func main() {
	addr := config.GetString("ADDR", "0.0.0.0:8080")
	app := router.New()
	app.Run(addr)
}
