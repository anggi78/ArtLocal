package services

import (
	"art-local/features/model"
	"art-local/helpers"
	"art-local/repositories"
	"log"
)

type UserServiceInterface interface {
	CreateUser(user model.User) (*model.User, error)
	Login (email string, password string) (*model.User, string, error)
	FindAll() ([]model.User, error) 
}

type userService struct {
	repo repositories.UserRepoInterface
}

func NewUserService(repo repositories.UserRepoInterface) *userService {
	return &userService{repo}
}

func (u *userService) CreateUser(user model.User) (*model.User, error) {
	createdUser := &model.User{}
	err := u.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (u *userService) Login(email string, password string) (*model.User, string, error) {
	userData, err := u.repo.Login(email, password)
	if err != nil {
		return userData, "", err
	}
	//log.Println("login : ", userData.Name)
	token, _ := helpers.GenerateToken(userData.UserID)
	if token != "" {
		log.Printf("Token : %s\n", token)
	}
	return userData, token, nil
}

func (u *userService) FindAll() ([]model.User, error) {
	users, err := u.repo.FindAll()
	if err != nil{
		return nil, err
	}
	return users, nil
}