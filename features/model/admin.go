package model

import (
	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model  
	Name     		string `gorm:"not null" valid:"required~your name is required"`
	Email    		string `gorm:"not null;unique" valid:"required~your email is required, email~invalid email format"`
	Password 		string `gorm:"not null" valid:"required~your password is required,minstringlength(6)~password has to have a minimum length of 6 characters"`
	Event	[]Event 
}