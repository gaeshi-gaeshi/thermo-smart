package config

var App = &app{
	DbDirectory: "/tmp/thermo-smart",
}

var Time = &time{
	RFC3339: "2006-01-02T15:04:05Z07:00",
}

type app struct {
	DbDirectory string
}

type time struct {
	RFC3339 string
}
