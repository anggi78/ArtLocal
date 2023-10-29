package main

import (
	"art-local/app/config"
	"art-local/app/database"
	"art-local/app/route"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := envPortOr("3000")
	_, dbConfig := config.InitConfig()
	database.InitDBMysql(dbConfig)

	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	route.UserRoute(app, database.DB)
	route.AdminRoute(app, database.DB)
	route.ArtworkRoute(app, database.DB)
	route.EventRoute(app, database.DB)
	route.AIReccRoute(app)
	
	app.Logger.Fatal(app.Start(port))
}

func envPortOr(port string) string {
	// If `PORT` variable in the environment exists, return it
	if envPort := os.Getenv("APP_PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}