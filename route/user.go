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

func UserRoute(e *echo.Echo, db *gorm.DB) {
	repositories := repositories.NewUserRepo(db)
	service := services.NewUserService(repositories)
	handler := handler.NewUserHandler(service)

	// user register dan login
	e.POST("/users/register", handler.Register)
	e.POST("/users/login", handler.Login)

	app := e.Group("")
	app.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// user update profil
	app.PUT("/users/:id", handler.Update)

	
	//app.GET("/users", handler.GetAllUsers)
}
