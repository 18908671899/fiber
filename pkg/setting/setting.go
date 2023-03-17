package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini':%v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadApp() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("dev")
}

func LoadServer() {
	serverSec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	HTTPPort = serverSec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(serverSec.Key("READ_TIMEOUT").MustInt(60))
	WriteTimeout = time.Duration(serverSec.Key("WRITE_TIMEOUT").MustInt(60))
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("dev")
}
