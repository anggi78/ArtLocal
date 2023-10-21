package repositories

import (
	"art-local/features/core"
	"art-local/features/model"
	"log"

	"gorm.io/gorm"
)

type UserRepoInterface interface {
	GetAll() ([]core.UserCore, error)
	CreateUser(core.UserCore) (core.UserCore, error)
	Login(email string, password string) (core.UserCore, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(DB *gorm.DB) *userRepo {
	return &userRepo{DB}
}

func (u *userRepo) GetAll() ([]core.UserCore, error) {
	var users []model.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	userCore := []core.UserCore{}
	for _, user := range users {
		v := core.FromModelToUserCore(user)
		userCore = append(userCore, v)
	}
	return userCore, nil
}

func (u *userRepo) CreateUser(user core.UserCore) (core.UserCore, error) {
	insert := core.FromCoreToUserModel(user)
	err := u.db.Create(&insert).Error
	if err != nil {
		return user, err
	}
	data := core.FromModelToUserCore(insert)
	return data, nil
}

func (u *userRepo) Login(email string, password string) (core.UserCore, error) {
	var user model.User
	var data core.UserCore

	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return data, err
	}

    data = core.FromModelToUserCore(user)
	log.Println(data)
	return data, nil
}