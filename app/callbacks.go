package app

func (app *TerminalApplication) onDataReceived(data []byte) {
	for _, b := range data {
		app.correlator.RX(b)
	}
}
