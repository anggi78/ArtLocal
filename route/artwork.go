package route

import (
	"art-local/handler"
	"art-local/repositories"
	"art-local/services"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ArtworkRoute(app *echo.Echo, db *gorm.DB) {
	repo := repositories.NewArtRepositories(db)
	service := services.NewArtService(repo)
	handler := handler.NewArtHandler(service) 

	app.GET("/users/art", handler.GetAll)

	e := app.Group("")
	app.GET("/users/art/:id", handler.GetById)
	e.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.POST("/users/art", handler.Create)
	e.DELETE("/users/art/:id", handler.Delete)
	e.PUT("/users/art/:id", handler.Update)
}