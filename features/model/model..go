package model

import "github.com/jinzhu/gorm"


type User struct {
	gorm.Model
	Name         string `gorm:"not null" valid:"required~your name is required"`
	Email        string `gorm:"not null;unique" valid:"required~your email is required, email~invalid email format"`
	Password     string `gorm:"not null"`
	Role      	 string `gorm:"type:ENUM('user', 'admin');not null;default:'user'"`
	Follow_event []FollowEvent
	Artwork      []Artwork
}

type Admin struct {
	gorm.Model
	Name         string `gorm:"not null" valid:"required~your name is required"`
	Email        string `gorm:"not null;unique" valid:"required~your email is required, email~invalid email format"`
	Password     string `gorm:"not null"`	
	Event []Event
}

type Event struct {
	gorm.Model
	Title        string `gorm:"not null" valid:"required~your title is required" json:"title" form:"title"`
	Date         string `gorm:"not null" valid:"required~your date  is required"`
	Description  string `gorm:"not null;unique" valid:"required~your description is required"`
	Location     string `gorm:"not null" valid:"required~your location is required"`
	AdminID      uint	`gorm:"constraint:OnDelete:CASCADE;"`
	Follow_event []FollowEvent
}

type Artwork struct {
	gorm.Model
	Title       string `gorm:"not null" valid:"required~your title is required" json:"title" form:"title"`
	Description string `gorm:"not null;unique" valid:"required~your description is required"`
	//Image       string `gorm:"not null" valid:"required~your image is required" json:"image" form:"image"`
	UserID 		uint	`gorm:"constraint:OnDelete:CASCADE;"`
}

type FollowEvent struct {
	gorm.Model
	UserID  uint
	EventID uint
}