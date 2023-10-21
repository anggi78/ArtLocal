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
	Update(userID int, user *core.User) (*model.User, error)
	FindByID(userID int) (*core.User, error)
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

func (u *userRepo) Update(userID int, user *core.User) (*model.User, error) {
	// updateUser := u.db.Table("user").Where("id = ?", userID).Updates(core.User{Name: user.Name, Email: user.Email, Password: user.Password})
	// if updateUser.Error != nil {
	// 	return nil, updateUser.Error
	// }
	// return user, nil
	var existingUser model.User

	if err := u.db.First(&existingUser, userID).Error; err != nil {
		return nil, err
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Password = user.Password

	if err := u.db.Save(&existingUser).Error; err != nil {
		return nil, err
	}
	return &existingUser, nil
}