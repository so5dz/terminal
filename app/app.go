package app

import (
	"log"

	"github.com/iskrapw/network/tcp"
	"github.com/iskrapw/terminal/correlator/ax25"
)

type TerminalApplication struct {
	dataClient tcp.Client
	correlator ax25.AX25Correlator // todo variability
}

func (app *TerminalApplication) Run() error {
	log.Println("Connecting to data server")
	err := app.dataClient.Connect()
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
	return nil
}
