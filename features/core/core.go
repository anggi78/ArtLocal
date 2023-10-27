package core

type User struct {
	ID			uint
	Name		string
	Email		string
	Password	string
}

type EventCore struct {
	ID			uint
	Title		string
	Date		string
	Description	string
	Location	string
	FollowEvent []FollowEventCore
}

type EventAll struct {
	ID			uint
	Name		string	`json:"name"`
	EventID     uint	`json:"event_id"`
	Title		string	`json:"title"`
	Date		string	`json:"date"`
	Description	string	`json:"description"`
	Location	string	`json:"location"`
}

type ArtworkCore struct {
	ID			uint
	Title		string
	Image		string
	Description	string
	UserID		uint
}

type FollowEventCore struct {
	ID uint
	UserID 	uint
	EventID uint
}

type Admin struct {
	ID			uint
	Name		string
	Email		string
	Password	string
}