package request

//import "mime/multipart"

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type EventRequest struct {
	Title       string `json:"title" form:"title"`
	Date        string `json:"date" form:"date"`
	Description string `json:"description" form:"description"`
	Location    string `json:"location" form:"location"`
}

type ArtworkRequest struct {
	Title       string                `json:"title" form:"title"`
	Image       string 				  `json:"image" form:"image"`
	Description string                `json:"description" form:"description"`
	UserID      uint                  `json:"user_id" form:"user_id"`
}

type FollowEventRequest struct {
	UserID  uint `json:"user_id" form:"user_id"`
	EventID uint `json:"event_id" form:"event_id"`
}

type AdminRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type AskAboutEvent struct {
	Category	string	`json:"category"`
	Location	string	`json:"location"`
}
