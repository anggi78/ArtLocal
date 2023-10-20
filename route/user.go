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

func UserRoute(app *echo.Echo, db *gorm.DB) {
	repositories := repositories.NewUserRepo(db)
	service := services.NewUserService(repositories)
	handler := handler.NewUserHandler(service)

	app.POST("/users/register", handler.Register)
	app.POST("users/login", handler.Login)

	apps := app.Group("users")
	apps.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	app.GET("/users", handler.GetAllUsers)
}