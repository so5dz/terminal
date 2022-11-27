package main

import (
	"github.com/so5dz/terminal/app"
	terminalconfig "github.com/so5dz/terminal/config"
	"github.com/so5dz/utils/config"
	"github.com/so5dz/utils/misc"
)

func main() {
	misc.WrapMain(mainWithError)()
}

func mainWithError() error {
	cfg, err := config.LoadConfigFromArgs[terminalconfig.Config]()
	if err != nil {
		return err
	}

	var app app.TerminalApplication

	err = app.Initialize(cfg)
	if err != nil {
		return err
	}

	err = app.Run()
	if err != nil {
		return err
	}

	misc.BlockUntilInterrupted()

	return app.Close()
}
