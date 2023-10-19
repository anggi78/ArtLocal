package repositories

import (
	"art-local/features/model"
	"art-local/helpers"


	"gorm.io/gorm"
)

type UserRepoInterface interface {
	CreateUser(user *model.User) error
	Login(email string, password string) (*model.User, error)
	FindAll([]model.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(DB *gorm.DB) *userRepo {
	return &userRepo{db: DB}
}

func (u *userRepo) FindAll([]model.User) error {
	var users []model.User
	err := u.db.Find(&users).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) CreateUser(user *model.User) error {
	if err := u.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepo) Login(email string, password string) (*model.User, error) {
	var user model.User
	
	err := u.db.Where("email= ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	comparePass, err := helpers.Compare([]byte(user.Password), []byte(password))
	if err != nil || !comparePass {
		return nil, err
	}
	return &user, nil
}