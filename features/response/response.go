package response

import "github.com/labstack/echo/v4"

type UserResponse struct {
	UserID uint
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type EventResponse struct {
	EventID     uint
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Location    string `json:"location"`
	AdminID     uint
}

type ArtworkResponse struct {
	ArtworkID   uint
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint
}

type Follow_eventResponse struct {
	UserID  uint
	EventID uint
}

func ResponseJSON(e echo.Context, status int, message string, data interface{}) error {
	response := echo.Map{
		"status": 	message,
		"data": 	data,
	}
	return e.JSON(status, response)
}