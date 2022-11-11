package configinit

import (
	"PCS_BACKEND_GO/global/database"
	"PCS_BACKEND_GO/service/basic"
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

	basic.Secret = []byte(cfg.Section("auth").Key("SECRET").String())
}
