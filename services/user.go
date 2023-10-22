package services

import (
	"art-local/features/core"
	"art-local/helpers"
	"art-local/repositories"
	"log"
)

type UserServiceInterface interface {
	CreateUser(user core.User) (core.User, error)
	Login(email string, password string) (core.User, string, error)
	GetAll() ([]core.User, error)
	Update(userID int, user core.User) (core.User, string, error)
}

type userService struct {
	repo repositories.UserRepoInterface
}

func NewUserService(repo repositories.UserRepoInterface) *userService {
	return &userService{repo}
}

func (u *userService) CreateUser(user core.User) (core.User, error) {
	user.Password = helpers.HashPassword(user.Password)
	users, err := u.repo.CreateUser(user)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *userService) Login(email string, password string) (core.User, string, error) {
	userData, err := s.repo.Login(email, password)
	if err != nil {
		return userData, "", err
	}
	log.Println("role in login service : ", userData.Role)
	token, _ := helpers.GenerateToken(userData.ID)
	return userData, token, nil
}

func (u *userService) GetAll() ([]core.User, error) {
	users, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userService) Update(userID int, user core.User) (core.User, string, error) {
	user.Password = helpers.HashPassword(user.Password)
    existingUser, err := u.repo.FindByID(userID)
    if err != nil {
        return core.User{}, "", err
    }

    existingUser.Name = user.Name
    existingUser.Email = user.Email
	existingUser.Password = user.Password

    if err := u.repo.Update(userID, *existingUser); err != nil {
        return core.User{}, "", err
    }
    return *existingUser, "", nil
}