package core

type User struct {
	ID			uint
	Name		string
	Email		string
	Password	string
	Role		string
}

type EventCore struct {
	ID			uint
	Title		string
	Date		string
	Description	string
	Location	string
	AdminID		uint
}

type ArtworkCore struct {
	ID			uint
	Title		string
	Image		string
	Description	string
	UserID		uint
}

type FollowEventCore struct {
	UserID 	uint
	EventID uint
}

type Admin struct {
	ID			uint
	Name		string
	Email		string
	Password	string
}