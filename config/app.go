package config

var App = &app{
	DbDirectory: "/tmp/thermo-smart",
	RFC112:      "Mon, 02 Jan 2006 15:04:05 MST",
}

type app struct {
	DbDirectory string
	RFC112      string
}
