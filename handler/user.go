package handler

import (
	"art-local/features/core"
	"art-local/features/request"
	"art-local/features/response"
	"art-local/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService services.UserServiceInterface
}

func NewUserHandler(userService services.UserServiceInterface) *userHandler {
	return &userHandler{userService}
}

func (u *userHandler) Register(e echo.Context) error {
	userReq := request.UserRequest{}
	err := e.Bind(&userReq)
	if err != nil {
		return response.ResponseJSON(e, 400, err.Error(), nil)
	}

	insert := core.FromRequestToUser(userReq)
	userData, err := u.userService.CreateUser(insert)
	if err != nil {
		return response.ResponseJSON(e, 400, err.Error(), nil)
	}

	respon := core.FromCoreToUserResponse(userData)
	return response.ResponseJSON(e, 200, "success", respon)
}

func (u *userHandler) Login(e echo.Context) error {
	userReq := request.UserRequest{}
	err := e.Bind(&userReq)
	if err != nil {
		return response.ResponseJSON(e, 401, err.Error(), nil)
	}

	email := userReq.Email
	password := userReq.Password

	userData, token, err := u.userService.Login(email, password)
	if err != nil {
		return response.ResponseJSON(e, 401, err.Error(), nil)
	}

	respon := core.FromCoreToUserResponse(userData)
	return e.JSON(200, echo.Map{
		"message": "success",
		"data": respon,
		"token": token,
	})
}

func (u *userHandler) GetAllUsers(e echo.Context) error {
	users, err := u.userService.GetAll()
	if err != nil {
		return response.ResponseJSON(e, 500, err.Error(), nil)
	}

	respon := []response.UserResponse{}
	for _, user := range users {
		userRes := response.UserResponse{
			Name: user.Name,
			Email: user.Email,
		}
		respon = append(respon, userRes)
	}
	return response.ResponseJSON(e, 200, "success", respon)
}

func (u *userHandler) Update(e echo.Context) error {
    idStr := e.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return response.ResponseJSON(e, 401, "Invalid user ID", nil)
    }
    newUser := request.UserRequest{}

    err = e.Bind(&newUser)
    if err != nil {
        return response.ResponseJSON(e, 400, err.Error(), nil)
    }

    newUsers := core.FromRequestToUser(newUser)
    updatedUser, updateErr := u.userService.Update(id, newUsers)
    if updateErr != nil {
        return response.ResponseJSON(e, 401, updateErr.Error(), nil)
    }
    return response.ResponseJSON(e, 200, "success", updatedUser)
}