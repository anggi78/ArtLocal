package services

import (
	"art-local/features/core"
	"art-local/features/model"
	"art-local/helpers"
	"art-local/repositories"
	"log"
)

type UserServiceInterface interface {
	CreateUser(user core.User) (core.User, error)
	Login(email string, password string) (core.User, string, error)
	GetAll() ([]core.User, error)
	Update(userID int, user *core.User) (*model.User, error)
}

type userService struct {
	repo repositories.UserRepoInterface
}

func NewUserService(repo repositories.UserRepoInterface) *userService {
	return &userService{repo}
}

func (u *userService) CreateUser(user core.User) (core.User, error) {
	users, err := u.repo.CreateUser(user)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (u *userService) Login(email string, password string) (core.User, string, error) {
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

func (u *userService) GetAll() ([]core.User, error) {
	users, err := u.repo.GetAll()
	if err != nil{
		return nil, err
	}
	return users, nil
}

func (u *userService) Update(userID int, user *core.User) (*model.User, error) {
    existingUser, err := u.repo.FindByID(userID)
    if err != nil {
        return nil, err
    }

    existingUser.Name = user.Name
    existingUser.Email = user.Email
    existingUser.Password = user.Password

    if err != u.repo.Update(&existingUser).Error; err != nil {
        return nil, err
    }
    return &existingUser, nil
}
