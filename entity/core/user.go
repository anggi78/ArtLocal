package core

import (
	"art-local/entity/model"
	"art-local/entity/request"
	"art-local/entity/response"
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
		ID: uint(core.ID),
		Name: core.Name,
		Email: core.Email,
		Password: core.Password,
	}
	return user
}

func FromModelToUser(model model.User) User {
	user := User{
		ID: model.ID,
		Name: model.Name,
		Email: model.Email,
		Password: model.Password,
	}
	return user
}

func FromCoreToUserResponse(users User, ID uint) response.UserResponse {
	user := response.UserResponse{
		ID: users.ID,
		Name: users.Name,
		Email: users.Email,
	}
	return user
}