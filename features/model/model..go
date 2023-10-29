package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null" valid:"required~your name is required"`
	Email       string `gorm:"not null;unique" valid:"required~your email is required, email~invalid email format"`
	Password    string `gorm:"not null"`
	FollowEvent []FollowEvent
	Artwork     []Artwork
}

type Admin struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null" valid:"required~your name is required"`
	Email    string `gorm:"not null;unique" valid:"required~your email is required, email~invalid email format"`
	Password string `gorm:"not null"`
}

type Event struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null" valid:"required~your title is required" json:"title" form:"title"`
	Date        string `gorm:"not null" valid:"required~your date  is required"`
	Description string `gorm:"not null;unique" valid:"required~your description is required"`
	Location    string `gorm:"not null" valid:"required~your location is required"`
	FollowEvent []FollowEvent
}

type Artwork struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null" valid:"required~your title is required" json:"title" form:"title"`
	Description string `gorm:"not null;unique" valid:"required~your description is required" form:"description"`
	Image       string `gorm:"not null"`
	UserID      uint   `gorm:"constraint:OnDelete:CASCADE;"`
}

type FollowEvent struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	UserID  uint
	EventID uint `gorm:"index; constraint:OnDelete:CASCADE;OnUpdate:CASCADE;"`
}