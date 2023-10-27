package core

import (
	"art-local/features/model"
	"art-local/features/request"
	"art-local/features/response"
)

func FromRequestToAdmin(Req request.AdminRequest) Admin {
	admin := Admin{
		Name: Req.Name,
		Email: Req.Email,
		Password: Req.Password,
	}
	return admin
}

func FromCoreToAdminModel(core Admin) model.Admin {
	admin := model.Admin{
		ID: core.ID,
		Name: core.Name,
		Email: core.Email,
		Password: core.Password,
	}
	return admin
}

func FromModelToAdmin(model model.Admin) Admin {
	admin := Admin{
		ID: model.ID,
		Name: model.Name,
		Email: model.Email,
		Password: model.Password,
	}
	return admin
}

func FromCoreToAdminResponse(Admins Admin) response.AdminResponse {
	Admin := response.AdminResponse{
		ID: Admins.ID,
		Name: Admins.Name,
		Email: Admins.Email,
	}
	return Admin
}