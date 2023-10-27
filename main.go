package main

import (
	"art-local/app/config"
	"art-local/app/database"
	"art-local/route"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	appConfig, dbConfig := config.InitConfig()
	database.InitDBMysql(dbConfig)

	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	route.UserRoute(app, database.DB)
	route.AdminRoute(app, database.DB)
	route.ArtworkRoute(app, database.DB)
	route.EventRoute(app, database.DB)
	//route.FollowRoute(app, database.DB)
	
	app.Logger.Fatal(app.Start(fmt.Sprintf(":%d", appConfig.APP_PORT)))
}