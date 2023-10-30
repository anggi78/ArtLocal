package services

import (
	"art-local/entity/core"
	"art-local/features/mocks"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateAdmin_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAdminRepo := mocks.NewMockAdminRepoInterface(ctrl)
    adminService := NewAdminService(mockAdminRepo)

    fakeAdmin := core.Admin{
        Name:     "anggi",
        Email:    "janggi@gmail.com",
        Password: "mysecret",
    }

    mockAdminRepo.EXPECT().CreateAdmin(gomock.Any()).Return(fakeAdmin, nil)

    createdAdmin, err := adminService.CreateAdmin(fakeAdmin)

    assert.NoError(t, err)
    assert.NotNil(t, createdAdmin)
    assert.Equal(t, fakeAdmin.Name, createdAdmin.Name)
    assert.Equal(t, fakeAdmin.Email, createdAdmin.Email)
}

func TestCreateAdmin_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAdminRepo := mocks.NewMockAdminRepoInterface(ctrl)
    adminService := NewAdminService(mockAdminRepo)

    fakeAdmin := core.Admin{
        Name:     "anggi",
        Email:    "janggi@gmail.com",
        Password: "mysecret",
    }

    expectedError := errors.New("Simulated error")

    mockAdminRepo.EXPECT().CreateAdmin(gomock.Any()).Return(core.Admin{}, expectedError)

    createdAdmin, err := adminService.CreateAdmin(fakeAdmin)

    assert.Error(t, err)
    assert.Equal(t, expectedError, err)
    assert.Empty(t, createdAdmin)
}

func TestCreateEvent_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAdminRepo := mocks.NewMockAdminRepoInterface(ctrl)
    adminService := NewAdminService(mockAdminRepo)

    event := core.EventCore{
        Title: "events local",
		Date: "20-10-2023",
		Description: "ini event yang sangat menarik",
		Location: "Palu",
    }

    mockAdminRepo.EXPECT().
        CreateEvent(event).
        Return(event, nil)

    createdEvent, err := adminService.CreateEvent(event)

    assert.NoError(t, err)
    assert.Equal(t, event, createdEvent)
}

func TestCreateEvent_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAdminRepo := mocks.NewMockAdminRepoInterface(ctrl)
    adminService := NewAdminService(mockAdminRepo)

    event := core.EventCore{
        Title: "events local",
		Date: "20-10-2023",
		Description: "ini event yang sangat menarik",
		Location: "Palu",
    }

    expectedError := errors.New("Simulated error")

    mockAdminRepo.EXPECT().
        CreateEvent(event).
        Return(core.EventCore{}, expectedError)

    createdEvent, err := adminService.CreateEvent(event)

    assert.Error(t, err)
    assert.Equal(t, expectedError, err)
    assert.Empty(t, createdEvent)
}

func TestLoginAdmin_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAdminRepoInterface(ctrl)
	service := NewAdminService(mockRepo)

	email := "aisyah@gmail.com.com"
	password := "password123"
	expectedUser := core.Admin{
		Name: "aisyah",
		Email:    "aisyah@gmail.com",
		Password: "hashed_password",
	}

	mockRepo.EXPECT().LoginAdmin(email, password).Return(expectedUser, nil)
	user, token, err := service.LoginAdmin(email, password)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	assert.NotEmpty(t, token)
}

func TestLoginAdmin_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAdminRepoInterface(ctrl)
	service := NewAdminService(mockRepo)

	email := "aisyah@gmail.com"
	password := "password123"
	expectedErr := errors.New("authentication failed")

	mockRepo.EXPECT().LoginAdmin(email, password).Return(core.Admin{}, expectedErr)
	user, token, err := service.LoginAdmin(email, password)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Empty(t, user)
	assert.Empty(t, token)
}

func TestUpdateAdmin_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAdminRepo := mocks.NewMockAdminRepoInterface(ctrl)
    adminService := NewAdminService(mockAdminRepo)

    adminID := 1
    admin := &core.Admin{ID: 1, Name: "pudu", Email: "pudu@gmail.com"}

    mockAdminRepo.EXPECT().
        FindByID(adminID).
        Return(admin, nil)

    mockAdminRepo.EXPECT().
        Update(adminID, gomock.Any()).
        Return(admin, nil)

    updatedAdmin, _, err := adminService.Update(adminID, admin)

    assert.NoError(t, err)
    assert.Equal(t, admin, updatedAdmin)
}

func TestUpdateAdmin_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAdminRepo := mocks.NewMockAdminRepoInterface(ctrl)
    adminService := NewAdminService(mockAdminRepo)

    adminID := 1
    admin := &core.Admin{ID: 1, Name: "pudu", Email: "pudu@gmail.com"}
    expectedError := errors.New("Simulated error")

    mockAdminRepo.EXPECT().
        FindByID(adminID).
        Return(admin, nil)

	mockAdminRepo.EXPECT().
        Update(adminID, gomock.Any()).
        Return(admin, expectedError)

    updatedAdmin, _, err := adminService.Update(adminID, admin)

    assert.Error(t, err)
    assert.Equal(t, expectedError, err)
    assert.Empty(t, updatedAdmin)
}