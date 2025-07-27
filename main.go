package main

import (
	"github.com/mjarmoc/modularized-go/internals/role"
	"github.com/mjarmoc/modularized-go/internals/user"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		user.Module, role.Module).Run()
}
