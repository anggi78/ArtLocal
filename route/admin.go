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

func AdminRoute(e *echo.Echo, db *gorm.DB) {
	repositories := repositories.NewAdminRepo(db)
	service := services.NewAdminService(repositories)
	handler := handler.NewAdminHandler(service)

	e.POST("/admin/login", handler.LoginAdmin)

	app := e.Group("")
	app.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	//app.GET("/users", handler.GetAllUsers)
	app.POST("/event", handler.CreateEvent)
}