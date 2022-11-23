package app

import (
	"log"

	"github.com/iskrapw/network/tcp"
	"github.com/iskrapw/terminal/config"
)

func (app *TerminalApplication) Initialize(cfg config.Config) error {
	app.correlator.Initialize() // todo
	app.initializeDataClient(cfg)
	return nil
}

func (app *TerminalApplication) initializeDataClient(cfg config.Config) {
	log.Println("Initializing modem data client")
	app.dataClient = tcp.NewClient(cfg.Connections.Modem.Host, cfg.Connections.Modem.DataPort, tcp.TCPConnectionMode_Stream)
	app.dataClient.OnReceive(app.onDataReceived)
}
