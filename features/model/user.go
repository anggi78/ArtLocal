package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     		string `gorm:"not null" valid:"required~your name is required"`
	Email    		string `gorm:"not null;unique" valid:"required~your email is required, email~invalid email format"`
	Password 		string `gorm:"not null" valid:"required~your password is required,minstringlength(6)~password has to have a minimum length of 6 characters"`
	Follow_event    []Follow_event
	Artwork			[]Artwork
}

type Event struct {
	gorm.Model 
	Title       	string `gorm:"not null" valid:"required~your title is required" json:"title" form:"title"`
	Date        	string `gorm:"not null" valid:"required~your date  is required"`
	Description 	string `gorm:"not null;unique" valid:"required~your description is required"`
	Location    	string `gorm:"not null" valid:"required~your location is required"`
	AdminID     	uint
	Follow_event    []Follow_event
}

type Artwork struct {
	gorm.Model
	Title       	string `gorm:"not null" valid:"required~your title is required" json:"title" form:"title"`
	Description 	string `gorm:"not null;unique" valid:"required~your description is required"`
	//Image
	UserID      	uint 
}

type Follow_event struct {
	gorm.Model
	UserID      	uint 
	EventID     	uint 
}