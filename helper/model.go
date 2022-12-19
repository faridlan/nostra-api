package helper

import (
	"github.com/faridlan/nostra-api/model/domain"
	"github.com/faridlan/nostra-api/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {

	var userResponses = []web.UserResponse{}
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}
