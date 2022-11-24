package app

import (
	"log"

	"github.com/iskrapw/network/tcp"
	"github.com/iskrapw/terminal/kiss"
)

func (app *TerminalApplication) onDataReceived(data []byte) {
	for _, b := range data {
		packet := app.correlator.RX(b)
		if len(packet) > 0 {
			app.kissServer.Broadcast(kiss.Encode(0, kiss.DataFrame, packet))
		}
	}
}

func (app *TerminalApplication) onKissReceived(remote tcp.Remote, data []byte) {
	log.Println(len(data))
}
