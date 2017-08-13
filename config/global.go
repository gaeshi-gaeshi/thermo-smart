package config

var App = &app{
	DbDirectory:  "/tmp/thermo-smart",
	DbName:       "thermo-smart",
	DbConnection: "127.0.0.1",
}

var Time = &time{
	RFC3339: "2006-01-02T15:04:05Z07:00",
}

type app struct {
	DbDirectory  string
	DbName       string
	DbConnection string
}

type time struct {
	RFC3339 string
}
