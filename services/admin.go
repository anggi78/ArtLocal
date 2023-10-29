package services

import (
	"art-local/features/core"
	"art-local/helpers"
	"art-local/repositories"
	"fmt"
	"log"
)

type AdminServiceInterface interface {
	CreateAdmin(admin core.Admin) (core.Admin, error)
	LoginAdmin(email string, password string) (core.Admin,string, error)
	CreateEvent(event core.EventCore) (core.EventCore, error)
	Update(AdminID int, admin *core.Admin) (*core.Admin, string, error)
}

type adminService struct {
	repo repositories.AdminRepoInterface
}

func NewAdminService(re repositories.AdminRepoInterface) *adminService {
	return &adminService{repo: re}
}

func (u *adminService) CreateAdmin(admin core.Admin) (core.Admin, error) {
    admin.Password = helpers.HashPassword(admin.Password)
    admins, err := u.repo.CreateAdmin(admin)
    if err != nil {
		fmt.Println(err)
		log.Println(admins)
        return admins, err
    }
    return admins, nil
}

func (u *adminService) LoginAdmin(email string, password string) (core.Admin, string, error) {
	adminData, err := u.repo.LoginAdmin(email, password)
	if err != nil {
		return adminData, "", err
	}

	token, _ := helpers.GenerateTokenAdmin(uint(adminData.ID))
	if token != "" {
		log.Printf("Token : %s\n", token)
	}
	return adminData, token, nil
}

func (u *adminService) CreateEvent(event core.EventCore) (core.EventCore, error) {
	events, err := u.repo.CreateEvent(event)
	if err != nil {
		return events, err
	}
	return events, nil
}

func (u *adminService) Update(AdminID int, admin *core.Admin) (*core.Admin, string, error) {
    admin.Password = helpers.HashPassword(admin.Password)
    existingUser, _ := u.repo.FindByID(AdminID)

    existingUser.Name = admin.Name
    existingUser.Email = admin.Email
    existingUser.Password = admin.Password

    updatedUser, updateErr := u.repo.Update(AdminID, existingUser)
    if updateErr != nil {
        return &core.Admin{},"", updateErr
    }
    return updatedUser,"", nil
}