package core

import (
	"art-local/features/model"
	"art-local/features/request"
	"art-local/features/response"
)

func FromRequestToAdmin(Req request.AdminRequest) Admin {
	Admin := Admin{
		Name: Req.Name,
		Email: Req.Email,
		Password: Req.Password,
	}
	return Admin
}

func FromCoreToAdminModel(core Admin) model.Admin {
	Admin := model.Admin{
		Name: core.Name,
		Email: core.Email,
		Password: core.Password,
	}
	return Admin
}

func FromModelToAdmin(model model.Admin) Admin {
	Admin := Admin{
		Name: model.Name,
		Email: model.Email,
		Password: model.Password,
	}
	return Admin
}

func FromCoreToAdminResponse(Admins Admin) response.AdminResponse {
	Admin := response.AdminResponse{
		Name: Admins.Name,
		Email: Admins.Email,
	}
	return Admin
}