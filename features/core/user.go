package core

import (
	"art-local/features/model"
	"art-local/features/request"
	"art-local/features/response"
)

func FromRequestToUserCore(Req request.UserRequest) UserCore {
	user := UserCore{
		Name: Req.Name,
		Email: Req.Email,
		Password: Req.Password,
	}
	return user
}

func FromCoreToUserModel(core UserCore) model.User {
	user := model.User{
		Name: core.Name,
		Email: core.Email,
		Password: core.Password,
	}
	return user
}

func FromModelToUserCore(model model.User) UserCore {
	user := UserCore{
		Name: model.Name,
		Email: model.Email,
		Password: model.Password,
	}
	return user
}

func FromCoreToUserResponse(users UserCore) response.UserResponse {
	user := response.UserResponse{
		Name: users.Name,
		Email: users.Email,
	}
	return user
}