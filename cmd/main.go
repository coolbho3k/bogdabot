package main

import (
	"bogdabot/service/bogdabot"
	"bogdabot/store/commands"
	"go.uber.org/fx"
)

func main() {
	options := fx.Options(
		// Commands storage
		commands.Module,
		// Bogdabot service
		bogdabot.Module,
	)

	fx.New(options).Run()
}
