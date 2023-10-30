package route

import (
	"art-local/features/handler"
	"art-local/helpers"
	"art-local/features/services"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func AIReccRoute(e *echo.Echo) {
	evtService := services.NewReccomEvent()
	evtHandler := handler.NewEventRecc(evtService)

	app := e.Group("")
	app.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	app.POST("/users/ask-about-event", evtHandler.Reccomend, helpers.AuthMiddleware("user"))
}