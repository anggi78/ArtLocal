package services

import (
	"art-local/features/core"
	"art-local/helpers"
	"art-local/repositories"
	"log"
)

type AdminServiceInterface interface {
	LoginAdmin(email string, password string) (core.Admin,string, error)
	//GetAll() ([]core.EventCore, error)
	CreateEvent(event core.EventCore) (core.EventCore, error)
	Update(AdminID int, admin core.Admin) (core.Admin, string, error)
}

type adminService struct {
	repo repositories.AdminRepoInterface
}

func NewAdminService(repo repositories.AdminRepoInterface) *adminService {
	return &adminService{}
}

func (u *adminService) LoginAdmin(email string, password string) (core.Admin, string, error) {
	adminData, err := u.repo.LoginAdmin(email, password)
	if err != nil {
		return adminData, "", err
	}

	log.Println("login : ", adminData.Email)
	token, _ := helpers.GenerateToken(adminData.ID)
	if token != "" {
		log.Printf("Token : %s\n", token)
	}
	return adminData, token, nil
}

// func (u *userService) GetAll() ([]core.EventCore, error) {
// 	users, err := u.repo.GetAll()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }

func (u *adminService) CreateEvent(event core.EventCore) (core.EventCore, error) {
	events, err := u.repo.CreateEvent(event)
	if err != nil {
		return events, err
	}
	return events, nil
}

func (u *adminService) Update(AdminID int, admin core.Admin) (core.Admin, string, error) {
	admin.Password = helpers.HashPassword(admin.Password)
    existingUser, err := u.repo.FindByID(AdminID)
    if err != nil {
        return core.Admin{}, "", err
    }

    existingUser.Name = admin.Name
    existingUser.Email = admin.Email
	existingUser.Password = admin.Password

    if err := u.repo.Update(AdminID, *existingUser); err != nil {
        return core.Admin{}, "", err
    }
    return *existingUser, "", nil
}