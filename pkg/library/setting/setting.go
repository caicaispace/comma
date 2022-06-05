package setting

import (
	"time"
)

type app struct {
	RootPath     string
	TimeFormat   string
	LogPath      string
	LogPrefix    string
	LogExtension string
}

var App = &app{}

type server struct {
	RunMode      string
	Port         string
	Host         string
	Addr         string
	Env          string
	RootPath     string
	LogPath      string
	ConfigPath   string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var Server = &server{}

type database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
	AutoMigrate bool
}

var Database = &database{}
