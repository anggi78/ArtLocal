package route

import (
	"art-local/features/handler"
	"art-local/helpers"
	"art-local/features/repositories"
	"art-local/features/services"
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

	// user
	app.POST("/users/art", handler.CreateArt, helpers.AuthMiddleware("user"))
	app.DELETE("/users/art/:id", handler.DeleteArt, helpers.AuthMiddleware("user"))
	app.PUT("/users/art/:id", handler.UpdateArt, helpers.AuthMiddleware("user"))
	app.GET("/users/art/:id", handler.GetByIdArt, helpers.AuthMiddleware("user"))
	app.GET("/users/art", handler.GetAllArt, helpers.AuthMiddleware("user"))

	// admin
	app.DELETE("/admin/art/:id", handler.DeleteArt, helpers.AuthMiddleware("admin"))
	app.GET("/admin/art/:id", handler.GetByIdArt, helpers.AuthMiddleware("admin"))
	app.GET("/admin/art", handler.GetAllArt, helpers.AuthMiddleware("admin"))
}