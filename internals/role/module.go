package role

import (
	roles_public "github.com/mjarmoc/modularized-go/internals/role/public"
	"go.uber.org/fx"
)

var Module = fx.Module("role",
	fx.Provide(
		NewRoleService,
		fx.Annotate(AsRoleFacade, fx.As(new(roles_public.RoleFacade))),
	),
)
