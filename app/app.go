package app

import (
	"log"

	"github.com/so5dz/network/tcp"
	"github.com/so5dz/terminal/correlator"
)

type TerminalApplication struct {
	dataClient tcp.Client
	kissServer tcp.Server
	correlator correlator.Correlator
}

func (app *TerminalApplication) Run() error {
	log.Println("Connecting to data server")
	err := app.dataClient.Connect()
	if err != nil {
		return err
	}

	log.Println("Starting KISS server")
	err = app.kissServer.Start()
	if err != nil {
		return err
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
