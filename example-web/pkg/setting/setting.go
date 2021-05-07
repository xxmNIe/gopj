package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort int

	ReadTimeout time.Duration
	WriteTimeout time.Duration

	PageSize int
	JwtSecret string
)

func init(){
	var err error
	fmt.Println("[init] start......")
	Cfg,err = ini.Load("conf/app.ini")
	if err !=nil{
		log.Fatalf("Fail to parse 'conf/app.ini' : %v",err)
	}
	loadBase()
	LoadServer()
	LoadApp()
	fmt.Println("[init] end......")
}

func loadBase(){
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

}

func LoadServer(){
	sec,err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)

	ReadTimeout = sec.Key("READ_TIMEOUT").MustDuration(60)
	WriteTimeout = sec.Key("WRITE_TIMEOUT").MustDuration(60)

}

func LoadApp(){
	sec,err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("23347$040412")
	PageSize = sec.Key("PAGE_SIZR").MustInt(10)
}
