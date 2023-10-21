package response

import "github.com/labstack/echo/v4"

type UserResponse struct {
	Name   		string 	`json:"name"`
	Email  		string 	`json:"email"`
	Password	string	`json:"password"`
}

type EventResponse struct {
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Location    string `json:"location"`
	AdminID     uint
}

type ArtworkResponse struct {
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