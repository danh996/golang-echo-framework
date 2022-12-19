package main

import (
	"database/sql"
	"log"

	"github.com/danh996/golang-echo-framework/cmd/api/handlers"
	"github.com/danh996/golang-echo-framework/util"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"

)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config", err)
	}
	_, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to DB", err)
	}

	e := echo.New()
	e.GET("/health-check", handlers.HealthCheckHandler)
	e.GET("/posts", handlers.PostIndexHandler)
	e.GET("/post/:id", handlers.PostDetailHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
