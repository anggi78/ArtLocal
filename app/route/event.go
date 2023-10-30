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

func EventRoute(e *echo.Echo, db *gorm.DB) {
	repo := repositories.NewEventRepositories(db)
	service := services.NewEventService(repo)
	handler := handler.NewEventHandler(service)

	app := e.Group("")
	app.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	// admin
	app.POST("/admin/event", handler.CreateEvent, helpers.AuthMiddleware("admin"))
	app.DELETE("/admin/event/:id", handler.DeleteEvent, helpers.AuthMiddleware("admin"))
	app.PUT("/admin/event/:id", handler.UpdateEvent, helpers.AuthMiddleware("admin"))
	app.GET("/admin/event/:id", handler.GetByIdEvent, helpers.AuthMiddleware("admin"))
	app.GET("/admin/event", handler.GetAllEvent, helpers.AuthMiddleware("admin"))

	// user
	app.GET("/users/event/:id", handler.GetByIdEvent, helpers.AuthMiddleware("user"))
	app.GET("/users/event", handler.GetAllEvent, helpers.AuthMiddleware("user"))
	app.POST("/users/event", handler.CreateEvent, helpers.AuthMiddleware("user"))
	app.DELETE("/users/event/:id", handler.DeleteEvent, helpers.AuthMiddleware("user"))
	app.PUT("/users/event/:id", handler.UpdateEvent, helpers.AuthMiddleware("user"))

	// user follow event
	app.GET("users/follow",  handler.GetAllUserEvent, helpers.AuthMiddleware("user"))
	app.POST("users/follow",  handler.CreateFollowEvent, helpers.AuthMiddleware("user"))
}