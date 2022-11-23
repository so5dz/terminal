package main

import (
	"github.com/iskrapw/terminal/app"
	terminalconfig "github.com/iskrapw/terminal/config"
	"github.com/iskrapw/utils/config"
	"github.com/iskrapw/utils/misc"
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
