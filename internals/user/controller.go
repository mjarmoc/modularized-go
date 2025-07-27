package user

import (
	"go.uber.org/fx"
)

type UserController struct {
	service *UserService
}

func NewUserController(p UserControllerParams) UserControllerResult {
	controller := &UserController{service: p.UserService}
	return UserControllerResult{UserController: *controller}
}

type UserControllerParams struct {
	fx.In
	UserService *UserService
	//Config    config.UserConfig `group:"config"`
}

type UserControllerResult struct {
	fx.Out
	UserController UserController `group:"userController"`
}
