package service

import (
	"context"

	"github.com/faridlan/nostra-api/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateReq) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateReq) web.UserResponse
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}
