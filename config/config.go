package config

type Config struct {
	Connections Connections `json:"connections"`
}

type Connections struct {
	Modem   ModemConnection   `json:"modem"`
	Control ControlConnection `json:"control"`
}

type ModemConnection struct {
	Host      string `json:"host"`
	DataPort  int    `json:"dataPort"`
	ExtraPort int    `json:"extraPort"`
}

type ControlConnection struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}
