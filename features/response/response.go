package response

import "github.com/labstack/echo/v4"

type UserResponse struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

type EventResponse struct {
	ID		uint	`json:"id"`
	Title       string 	`json:"title"`
	Date        string 	`json:"date"`
	Description string 	`json:"description"`
	Location    string 	`json:"location"`
}

type ArtworkResponse struct {
	ID			uint	`json:"id"`
	Title       string 	`json:"title"`
	Image       string 	`json:"image"`
	Description string 	`json:"description"`
	UserID      uint	`json:"user_id"`
}

type FollowEventResponse struct {
	UserID  uint `json:"user_id"`
	EventID uint `json:"event_id"`
}

type AdminResponse struct {
	ID 			uint	`json:"id"`
	Name   		string 	`json:"name"`
	Email  		string 	`json:"email"`
}

// type UserEventResponse struct {
//     Event EventResponse `json:"event"`
//     User  UserResponse  `json:"user"`
// }


func ResponseJSON(e echo.Context, status int, message string, data interface{}) error {
	response := echo.Map{
		"status": 	message,
		"data": 	data,
	}
	return e.JSON(status, response)
}