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
	//Image       *multipart.FileHeader `json:"image" form:"image"`
	Description string                `json:"description" form:"description"`
	UserID      uint                  `json:"user_id"`
}

type FollowEventRequest struct {
	UserID  uint
	EventID uint
}

type AdminRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}