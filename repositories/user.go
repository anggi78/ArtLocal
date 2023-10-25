package repositories

import (
	//"art-local/app/config"
	"art-local/features/core"
	"art-local/features/model"
	"log"

	"gorm.io/gorm"
)

type UserRepoInterface interface {
	GetAll() ([]core.User, error)
	CreateUser(core.User) (core.User, error)
	Login(email string, password string) (core.User, error)
	Update(userID int, user core.User) error
	FindByID(userID int) (*core.User, error)
	Delete(id int) (bool, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(DB *gorm.DB) *userRepo {
	return &userRepo{db: DB} 
}

func (u *userRepo) GetAll() ([]core.User, error) {
	var users []model.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	User := []core.User{}
	for _, user := range users {
		v := core.FromModelToUser(user)
		User = append(User, v)
	}
	return User, nil
}

func (u *userRepo) CreateUser(user core.User) (core.User, error) {
	insert := core.FromCoreToUserModel(user)
	err := u.db.Create(&insert).Error
	if err != nil {
		return user, err
	}
	data := core.FromModelToUser(insert)
	return data, nil
}

func (u *userRepo) Login(email string, password string) (core.User, error) {
	var user model.User
	var data core.User

	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return data, err
	}

    data = core.FromModelToUser(user)
	log.Println(data)
	return data, nil
}

func (u *userRepo) FindByID(userID int) (*core.User, error) {
	user := core.User{}

	result := u.db.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *userRepo) Update(userID int, user core.User) error {
    var existingUser model.User

    if err := u.db.First(&existingUser, userID).Error; err != nil {
        return err
    }

    existingUser.Name 		= user.Name
    existingUser.Email 		= user.Email
    existingUser.Password 	= user.Password

    if err := u.db.Save(&existingUser).Error; err != nil {
        return err
    }
    return nil
}

func (u *userRepo) Delete(id int) (bool, error) {
	users, err := u.FindByID(id)

	dataUsers := core.FromCoreToUserModel(*users)
	if err != nil {
		return false, err
	}

	err = u.db.Where("id = ?", id).Delete(&dataUsers).Error
	if err != nil {
		return false, err
	}
	return true, nil
}