package configinit

import (
	"PCS_BACKEND_GO/authentication"
	"PCS_BACKEND_GO/global"
	"PCS_BACKEND_GO/global/database"
	"log"

	"github.com/go-ini/ini"
)

func init() {

	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err.Error())
	}

	database.DBName = cfg.Section("database").Key("DBNAME").String()
	database.Host = cfg.Section("database").Key("HOST").String()
	database.Port = cfg.Section("database").Key("PORT").String()
	database.Username = cfg.Section("database").Key("USERNAME").String()
	database.Password = cfg.Section("database").Key("PASSWORD").String()

	authentication.Secret = []byte(cfg.Section("auth").Key("SECRET").String())

	global.IP = cfg.Section("host").Key("IP").String()
	global.Port = cfg.Section("host").Key("PORT").String()

}
