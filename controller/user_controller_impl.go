package controller

import (
	"net/http"
	"strconv"

	"github.com/faridlan/nostra-api/helper"
	"github.com/faridlan/nostra-api/model/web"
	"github.com/faridlan/nostra-api/service"
	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateReq := web.UserCreateReq{}
	helper.RequestFromBody(request, &userCreateReq)

	user := controller.UserService.Create(request.Context(), userCreateReq)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   user,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateReq := web.UserUpdateReq{}
	helper.RequestFromBody(request, &userUpdateReq)

	userId := params.ByName("userId")

	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdateReq.Id = id

	user := controller.UserService.Update(request.Context(), userUpdateReq)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   user,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.UserService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	user := controller.UserService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   user,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	users := controller.UserService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   users,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
