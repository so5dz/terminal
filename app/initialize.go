package app

import (
	"log"
	"strings"

	"github.com/so5dz/terminal/config"
	"github.com/so5dz/terminal/correlator/ax25"
	"github.com/so5dz/utils/misc"
)

func (app *TerminalApplication) Initialize(cfg config.Config) error {
	app.initializeDataClient(cfg)
	app.initializeKissServer(cfg)
	return app.initializeCorrelator(cfg)
}

func (app *TerminalApplication) initializeDataClient(cfg config.Config) {
	log.Println("Initializing modem data client")
	app.dataClient.Initialize(cfg.Connections.Modem.Host, cfg.Connections.Modem.DataPort)
	app.dataClient.OnReceive(app.onDataReceived)
}

func (app *TerminalApplication) initializeKissServer(cfg config.Config) {
	log.Println("Initializing KISS server")
	app.kissServer.Initialize(cfg.KissPort)
	app.kissServer.OnReceive(app.onKissReceived)
}

func (app *TerminalApplication) initializeCorrelator(cfg config.Config) error {
	log.Println("Initializing correlator")
	correlatorName := strings.ToUpper(cfg.Correlator)
	switch correlatorName {
	case "AX25":
		app.correlator = &ax25.AX25Correlator{}
		app.correlator.Initialize()
		return nil
	}
	return misc.NewError("unknown correlator", correlatorName)
}
