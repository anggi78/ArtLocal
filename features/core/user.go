package core

import (
	"art-local/features/model"
	"art-local/features/request"
	"art-local/features/response"
)

func FromRequestToUser(Req request.UserRequest) User {
	user := User{
		Name: Req.Name,
		Email: Req.Email,
		Password: Req.Password,
	}
	return user
}

func FromCoreToUserModel(core User) model.User {
	user := model.User{
		Name: core.Name,
		Email: core.Email,
		Password: core.Password,
	}
	return user
}

func FromModelToUser(model model.User) User {
	user := User{
		Name: model.Name,
		Email: model.Email,
		Password: model.Password,
	}
	return user
}

func FromCoreToUserResponse(users User) response.UserResponse {
	user := response.UserResponse{
		Name: users.Name,
		Email: users.Email,
	}
	return user
}