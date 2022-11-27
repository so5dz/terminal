package app

import (
	"log"

	"github.com/so5dz/network/server"
	"github.com/so5dz/terminal/kiss"
)

func (app *TerminalApplication) onDataReceived(data []byte) {
	for _, b := range data {
		packet := app.correlator.Feed(b)
		if len(packet) > 0 {
			kissEncodedPacket := kiss.Encode(0, kiss.DataFrame, packet)
			app.kissServer.Broadcast(kissEncodedPacket)
		}
	}
}

func (app *TerminalApplication) onKissReceived(remote server.Remote, data []byte) {
	log.Println(len(data))
}
