package user

import (
	user_public "github.com/mjarmoc/modularized-go/internals/user/public"
	"go.uber.org/fx"
)

var Module = fx.Module("user",
	fx.Supply(fx.Annotate("abc", fx.ResultTags(`name:"TableName"`))),
	fx.Provide(
		NewUserController,
		NewUserService,
		fx.Annotate(AsUserFacade, fx.As(new(user_public.UserFacade))),
	),
	fx.Invoke(NewUserController),
)
