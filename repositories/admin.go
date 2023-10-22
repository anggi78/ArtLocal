package repositories

import (
	"art-local/features/core"
	"art-local/features/model"
	"log"

	"gorm.io/gorm"
)

type AdminRepoInterface interface {
	LoginAdmin(email string, password string) (core.Admin, error)
	//GetAll() ([]core.EventCore, error)
	CreateEvent(event core.EventCore) (core.EventCore, error)
	Update(AdminID int, admin core.Admin) error
	DeleteEvent(AdminID int)
	FindByID(AdminID int) (*core.Admin, error)
}

type adminRepo struct {
	db *gorm.DB
}

// DeleteEvent implements AdminRepoInterface.
func (*adminRepo) DeleteEvent(AdminID int) {
	panic("unimplemented")
}

func NewAdminRepo(DB *gorm.DB) *adminRepo {
	return &adminRepo{db: DB}
}

func (u *adminRepo) LoginAdmin(email string, password string) (core.Admin, error) {
	var admin model.Admin
	var data core.Admin

	err := u.db.Where("email = ?", email).First(&admin).Error
	if err != nil {
		return data, err
	}

	data = core.FromModelToAdmin(admin)
	log.Println(data)
	return data, nil
}

// func (u *userRepo) GetAll() ([]core.EventCore, error) {
// 	var events []model.Event
// 	err := u.db.Find(&events).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	eventCores := []core.EventCore{}
// 	for _, event := range events {
// 		eventCore := core.EventModelToEventCore(event)
// 		eventCores = append(eventCores, eventCore)
// 	}
// 	return eventCores, nil
// }

func (u *adminRepo) CreateEvent(event core.EventCore) (core.EventCore, error) {
	insert := core.EventCoreToEventModel(event)
	err := u.db.Create(&insert).Error
	if err != nil {
		return event, err
	}
	data := core.EventModelToEventCore(insert)
	return data, nil
}

func (u *adminRepo) FindByID(AdminID int) (*core.Admin, error) {
	admin := core.Admin{}

	result := u.db.First(&admin, AdminID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &admin, nil
}

func (u *adminRepo) Update(AdminID int, admin core.Admin) error {
	var existingUser model.Admin

	if err := u.db.First(&existingUser, AdminID).Error; err != nil {
		return err
	}

	existingUser.Name = admin.Name
	existingUser.Email = admin.Email
	existingUser.Password = admin.Password

	if err := u.db.Save(&existingUser).Error; err != nil {
		return err
	}
	return nil
}
