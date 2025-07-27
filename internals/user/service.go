package user

import (
	user_public "github.com/mjarmoc/modularized-go/internals/user/public"
	"go.uber.org/fx"
)

type UserService struct {
	TableName string
}

type UserServiceParams struct {
	fx.In
	TableName string `name:"TableName"`
	//Config    config.UserConfig `group:"config"`
}

type UserServiceResult struct {
	fx.Out
	UserService *UserService
}

func NewUserService(p UserServiceParams) UserServiceResult {
	service := &UserService{
		TableName: p.TableName,
	}
	return UserServiceResult{UserService: service}
}

func AsUserFacade(s *UserService) user_public.UserFacade {
	return s
}

func (s *UserService) GetUser() {
	println("I AM WORKING!")
}

func (s *UserService) GetUserRoles() {}
