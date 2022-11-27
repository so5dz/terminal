package app

import (
	"log"

	tcpc "github.com/so5dz/network/client/tcp"
	tcps "github.com/so5dz/network/server/tcp"
	"github.com/so5dz/terminal/correlator"
	"github.com/so5dz/utils/misc"
)

const _DataClientConnectError = "unable to connect to data server"
const _KISSServerStartError = "unable to start KISS server"

type TerminalApplication struct {
	dataClient tcpc.StreamClient
	kissServer tcps.StreamServer
	correlator correlator.Correlator
}

func (app *TerminalApplication) Run() error {
	log.Println("Connecting to data server")
	err := app.dataClient.Connect()
	if err != nil {
		return misc.WrapError(_DataClientConnectError, err)
	}

	log.Println("Starting KISS server")
	err = app.kissServer.Start()
	if err != nil {
		return misc.WrapError(_KISSServerStartError, err)
	}

	log.Println("MARDES-terminal started, interrupt to close")
	return nil
}

func (app *TerminalApplication) Close() error {
	err := app.dataClient.Disconnect()
	if err != nil {
		log.Println(err)
	}
	err = app.kissServer.Stop()
	if err != nil {
		log.Println(err)
	}
	return nil
}
