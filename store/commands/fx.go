package commands

import (
	"log"

	"go.uber.org/fx"
)

var Module = fx.Provide(NewFX)

type Params struct {
	fx.In
}

type Result struct {
	fx.Out

	Store Store
}

func NewFX(p Params) Result {
	store, err := New()
	if err != nil {
		log.Panic(err)
	}

	return Result{
		Store: store,
	}
}
