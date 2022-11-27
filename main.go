package main

import (
	"github.com/so5dz/terminal/app"
	terminalconfig "github.com/so5dz/terminal/config"
	"github.com/so5dz/utils/config"
	"github.com/so5dz/utils/misc"
)

const _ConfigLoadError = "unable to load/read config"
const _AppInitializationError = "unable to initialize application"
const _AppStartError = "unable to start application"

func main() {
	misc.WrapMain(mainWithError)()
}

func mainWithError() error {
	cfg, err := config.LoadConfigFromArgs[terminalconfig.Config]()
	if err != nil {
		return misc.WrapError(_ConfigLoadError, err)
	}

	var app app.TerminalApplication

	err = app.Initialize(cfg)
	if err != nil {
		return misc.WrapError(_AppInitializationError, err)
	}

	err = app.Run()
	if err != nil {
		return misc.WrapError(_AppStartError, err)
	}

	misc.BlockUntilInterrupted()

	return app.Close()
}
