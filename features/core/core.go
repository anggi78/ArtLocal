package core

type UserCore struct {
	ID			uint
	Name		string
	Email		string
	Password	string
}

type EventCore struct {
	Title		string
	Date		string
	Description	string
	Location	string
	AdminID		uint
}

type ArtworkCore struct {
	Title		string
	Description	string
	UserID		uint
}

type Follow_eventCore struct {
	UserID 	uint
	EventID uint
}