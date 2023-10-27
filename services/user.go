package services

import (
	"art-local/features/core"
	"art-local/helpers"
	"art-local/repositories"
)

type UserServiceInterface interface {
	CreateUser(user core.User) (core.User, error)
	Login(email string, password string) (core.User, string, error)
	GetAll() ([]core.User, error)
	Update(ID uint, user core.User) (core.User, string, error)
	Delete(ID uint) (bool, error)
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
	token, _ := helpers.GenerateToken(uint(userData.ID))
	return userData, token, nil
}

func (u *userService) GetAll() ([]core.User, error) {
	users, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userService) Update(ID uint, user core.User) (core.User, string, error) {
	user.Password = helpers.HashPassword(user.Password)
    existingUser, err := u.repo.FindByID(ID)
    if err != nil {
        return core.User{}, "", err
    }

    existingUser.Name = user.Name
    existingUser.Email = user.Email
	existingUser.Password = user.Password

    if err := u.repo.Update(ID, *existingUser); err != nil {
        return core.User{}, "", err
    }
    return *existingUser, "", nil
}

func (u *userService) Delete(ID uint) (bool, error) {
	users, err := u.repo.Delete(ID)
	if err != nil {
		return false, err
	}
	return users, nil
}