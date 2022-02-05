package bogdabot

import (
	"bogdabot/store/commands"
	"log"

	"go.uber.org/fx"
)

var Module = fx.Invoke(NewFX)

type Params struct {
	fx.In

	Store commands.Store
	Lifecycle fx.Lifecycle
}

type Result struct {
	fx.Out

	Service Service
}

func NewFX(p Params) Result {
	service, err := New(p.Lifecycle, p.Store)
	if err != nil {
		log.Panic(err)
	}

	return Result{
		Service: service,
	}
}
