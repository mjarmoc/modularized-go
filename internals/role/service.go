package role

import (
	roles_public "github.com/mjarmoc/modularized-go/internals/role/public"
	"go.uber.org/fx"
)

type RoleService struct {
}

type RoleServiceParams struct {
	fx.In
	//Config    config.UserConfig `group:"config"`
}

type RoleServiceResult struct {
	fx.Out
	RoleService *RoleService
}

func NewRoleService(p RoleServiceParams) RoleServiceResult {
	service := &RoleService{}
	return RoleServiceResult{RoleService: service}
}

func AsRoleFacade(s *RoleService) roles_public.RoleFacade {
	return s
}

func (s *RoleService) GetRolesByUser() {}
