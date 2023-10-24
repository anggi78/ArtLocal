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

	// admin register dan login
	e.POST("/admin/register", handler.RegisterAdmin)
	e.POST("/admin/login", handler.LoginAdmin)

	admin := e.Group("")
	admin.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	//app.GET("/users", handler.GetAllUsers)

	// update profil admin
	admin.PUT("/admin/:id", handler.Update)

	// admin buat event
	admin.POST("/admin/event", handler.CreateEvent)
}