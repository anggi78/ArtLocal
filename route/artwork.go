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

func ArtworkRoute(e *echo.Echo, db *gorm.DB) {
	repo := repositories.NewArtRepositories(db)
	service := services.NewArtService(repo)
	handler := handler.NewArtHandler(service) 

	app := e.Group("")
	app.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))
	app.POST("/users/art", handler.CreateArt)
	app.DELETE("/users/art/:id", handler.DeleteArt)
	app.PUT("/users/art/:id", handler.UpdateArt)
	app.GET("/users/art/:id", handler.GetByIdArt)
	app.GET("/users/art", handler.GetAllArt)
}