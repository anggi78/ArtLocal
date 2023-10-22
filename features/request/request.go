package request

type UserRequest struct {
	Name		string	`json:"name" form:"name"`
	Email		string	`json:"email" form:"email"`
	Password	string	`json:"password" form:"password"`
}

type EventRequest struct {
	Title 		string	`json:"title" form:"title"`
	Date		string	`json:"date" form:"date"`
	Description	string	`json:"description" form:"description"`
	Location	string	`json:"location" form:"location"`
	AdminID		uint
}

type ArtworkRequest struct {
	Title 		string	`json:"title" form:"title"`
	Description	string	`json:"description" form:"description"`
	UserID		uint
}

type FollowEventRequest struct {
	UserID		uint
	EventID		uint	
}

type AdminRequest struct {
	Name		string	`json:"name" form:"name"`
	Email		string	`json:"email" form:"email"`
	Password	string	`json:"password" form:"password"`
}