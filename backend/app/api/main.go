package main

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"order_satang/app/api/factory"
	"order_satang/configs"
	"order_satang/database"
)

func main() {

	var (
		conf = configs.Init()
		db   = database.NewPostgresConn(conf.PostgresConfig)
	)

	e := echo.New()

	factory.DependencyResolve(e, db)

	// Start the server
	e.Start(fmt.Sprintf(":%s", conf.App.Port))
}
