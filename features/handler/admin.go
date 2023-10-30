package handler

import (
	"art-local/entity/core"
	"art-local/entity/request"
	"art-local/entity/response"
	"art-local/features/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type adminHandler struct {
	adminService services.AdminServiceInterface
}

func NewAdminHandler(adminService services.AdminServiceInterface) *adminHandler {
	return &adminHandler{adminService}
}

func (u *adminHandler) RegisterAdmin(e echo.Context) error {
    adminReq := request.AdminRequest{}
    err := e.Bind(&adminReq)
    if err != nil {
        return response.ResponseJSON(e, 400, err.Error(), nil)
    }

    insert := core.FromRequestToAdmin(adminReq)
    adminData, err := u.adminService.CreateAdmin(insert)
    if err != nil {
        return response.ResponseJSON(e, 400, err.Error(), nil)
    }

    adminResponse := core.FromCoreToAdminResponse(adminData)
    return response.ResponseJSON(e, 200, "success", adminResponse)
}

func (u *adminHandler) LoginAdmin(e echo.Context) error {
	adminReq := request.AdminRequest{}
	err := e.Bind(&adminReq)
	if err != nil {
		return response.ResponseJSON(e, 401, err.Error(), nil)
	}

	email := adminReq.Email
	password := adminReq.Password

	adminData, token, err := u.adminService.LoginAdmin(email, password)
	if err != nil {
		return response.ResponseJSON(e, 401, err.Error(), nil)
	}

	respon := core.FromCoreToAdminResponse(adminData)
	return e.JSON(200, echo.Map{
		"message": "success",
		"data": respon,
		"token": token,
	})
}

func (u *adminHandler) CreateEvent(e echo.Context) error {
	adminReq := request.EventRequest{}
	err := e.Bind(&adminReq)
	if err != nil {
		return response.ResponseJSON(e, 400, err.Error(), nil)
	}

	insert := core.EventRequestToEventCore(adminReq)
	adminData, err := u.adminService.CreateEvent(insert)
	if err != nil {
		return response.ResponseJSON(e, 400, err.Error(), nil)
	}

	respon := core.EventCoreToEventRespon(adminData, uint(adminData.ID))
	return response.ResponseJSON(e, 200, "success", respon)
}

func (u *adminHandler) Update(e echo.Context) error {
    idStr := e.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return response.ResponseJSON(e, 401, "Invalid user ID", nil)
    }
    newAdmin := request.AdminRequest{}

    err = e.Bind(&newAdmin)
    if err != nil {
        return response.ResponseJSON(e, 400, err.Error(), nil)
    }

    newAdmins := core.FromRequestToAdmin(newAdmin)
    updatedUser, _, updateErr := u.adminService.Update(id, &newAdmins)
    if updateErr != nil {
        return response.ResponseJSON(e, 401, updateErr.Error(), nil)
    }
    return response.ResponseJSON(e, 200, "success", updatedUser)
}