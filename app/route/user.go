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

func UserRoute(e *echo.Echo, db *gorm.DB) {
	repositories := repositories.NewUserRepo(db)
	service := services.NewUserService(repositories)
	handler := handler.NewUserHandler(service)

	// user register dan login
	e.POST("/users/register", handler.RegisterUsers)
	e.POST("/users/login", handler.LoginUsers)

	app := e.Group("")
	app.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// user 
	app.PUT("/users/:id", handler.UpdateUsers, helpers.AuthMiddleware("user"))

	// admin
	app.GET("/users", handler.GetAllUsers, helpers.AuthMiddleware("admin"))
	app.DELETE("/users/:id", handler.DeleteUsers, helpers.AuthMiddleware("admin"))
}
