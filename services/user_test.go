package services

import (
	"art-local/features/core"
	"art-local/mocks"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepoInterface(ctrl)
	userService := NewUserService(mockUserRepo)

	fakeUser := core.User{
		Name: "anggi",
		Email: "anggi@gmail.com",
		Password: "123",
	}

	mockUserRepo.EXPECT().CreateUser(gomock.Any()).Return(fakeUser, nil)

	createdUserResult, err := userService.CreateUser(fakeUser)

	assert.NoError(t, err)
	assert.Equal(t, fakeUser, createdUserResult)
}

func TestCreateUser_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepoInterface(ctrl)
	userService := NewUserService(mockUserRepo)

	fakeUser := core.User{
		Name: "anggi",
		Email: "anggi@gmail.com",
		Password: "123",
	}

	expectedError := errors.New("Simulasi error saat membuat pengguna")
	mockUserRepo.EXPECT().CreateUser(gomock.Any()).Return(core.User{}, expectedError)
	createdUser, err := userService.CreateUser(fakeUser)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, core.User{}, createdUser)
}

func TestUserDelete_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepoInterface(ctrl)
	userService := NewUserService(mockUserRepo)

	userID := uint(1)

	mockUserRepo.EXPECT().Delete(userID).Return(true, nil)

	result, err := userService.Delete(userID)

	assert.NoError(t, err)
	assert.True(t, result)
}

func TestUserDelete_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepoInterface(ctrl)
	userService := NewUserService(mockUserRepo)

	userID := uint(1)

	expectedError := errors.New("Simulasi error saat menghapus pengguna")
	mockUserRepo.EXPECT().Delete(userID).Return(false, expectedError) 

	result, err := userService.Delete(userID)

	assert.Error(t, err)
	assert.False(t, result)
}

func TestGetAllUsers_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
    
    mockUserRepo := mocks.NewMockUserRepoInterface(ctrl)

    expectedUsers := []core.User{
		{Name: "anggi",
		Email: "anggi@gmail.com",
		Password: "123",},
		{Name: "aisyah",
		Email: "aisyah@gmail.com",
		Password: "321",},
	}
    mockUserRepo.EXPECT().GetAll().Return(expectedUsers, nil)

    userService := NewUserService(mockUserRepo)

    users, err := userService.GetAll()

    assert.Nil(t, err)
    assert.Equal(t, expectedUsers, users)
}

func TestGetAllUsers_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
    
    mockUserRepo := mocks.NewMockUserRepoInterface(ctrl)

    mockUserRepo.EXPECT().GetAll().Return(nil, errors.New("Gagal mengambil data pengguna"))

    userService := NewUserService(mockUserRepo)

    users, err := userService.GetAll()

    assert.NotNil(t, err)
    assert.Nil(t, users)
}


func TestLogin_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepoInterface(ctrl)
	service := NewUserService(mockRepo)

	email := "aisyah@gmail.com.com"
	password := "password123"
	expectedUser := core.User{
		Name: "aisyah",
		Email:    "aisyah@gmail.com",
		Password: "hashed_password",
	}

	mockRepo.EXPECT().Login(email, password).Return(expectedUser, nil)
	user, token, err := service.Login(email, password)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	assert.NotEmpty(t, token)
}

func TestLogin_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepoInterface(ctrl)
	service := NewUserService(mockRepo)

	email := "aisyah@gmail.com"
	password := "password123"
	expectedErr := errors.New("authentication failed")

	mockRepo.EXPECT().Login(email, password).Return(core.User{}, expectedErr)
	user, token, err := service.Login(email, password)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Empty(t, user)
	assert.Empty(t, token)
}

func TestUpdateUser_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepo := mocks.NewMockUserRepoInterface(ctrl)
    userService := NewUserService(mockUserRepo)

    userID := uint(1)
    user := core.User{ID: 1, Name: "pudu", Email: "pudu@gmail.com"}

    mockUserRepo.EXPECT().
        FindByID(userID).
        Return(&user, nil)

    mockUserRepo.EXPECT().
        Update(userID, gomock.Any()).
        Return(nil)

    updatedUser, _, err := userService.Update(userID, user)

    assert.NoError(t, err)
    assert.Equal(t, user, updatedUser)
}

func TestUpdateUser_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepo := mocks.NewMockUserRepoInterface(ctrl)
    userService := NewUserService(mockUserRepo)

    userID := uint(1)
    user := core.User{ID: 1, Name: "pudu", Email: "pudu@gmail.com"}
    expectedError := errors.New("Simulated error")

    mockUserRepo.EXPECT().
        FindByID(userID).
        Return(&user, nil)

    mockUserRepo.EXPECT().
        Update(userID, gomock.Any()).
        Return(expectedError)

    updatedUser, _, err := userService.Update(userID, user)

    assert.Error(t, err)
    assert.Equal(t, expectedError, err)
    assert.Empty(t, updatedUser)
}