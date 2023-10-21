package services

import (
	"art-local/features/core"
	"art-local/helpers"
	"art-local/repositories"
	"log"
)

type UserServiceInterface interface {
	CreateUser(user core.UserCore) (core.UserCore, error)
	Login(email string, password string) (core.UserCore, string, error)
	GetAll() ([]core.UserCore, error)
}

type userService struct {
	repo repositories.UserRepoInterface
}

func NewUserService(repo repositories.UserRepoInterface) *userService {
	return &userService{repo}
}

func (u *userService) CreateUser(user core.UserCore) (core.UserCore, error) {
	users, err := u.repo.CreateUser(user)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (u *userService) Login(email string, password string) (core.UserCore, string, error) {
	userData, err := u.repo.Login(email, password)
	if err != nil {
		return userData, "", err
	}

	log.Println("login : ", userData.Email)
	token, _ := helpers.GenerateToken(userData.ID)
	if token != "" {
		log.Printf("Token : %s\n", token)
	}
	return userData, token, nil
}

func (u *userService) GetAll() ([]core.UserCore, error) {
	users, err := u.repo.GetAll()
	if err != nil{
		return nil, err
	}
	return users, nil
}