package main

import (
	// init
	"PCS_BACKEND_GO/global"
	_ "PCS_BACKEND_GO/init/configinit"
	_ "PCS_BACKEND_GO/init/sqlinit"
	"os"

	"PCS_BACKEND_GO/init/routerinit"
	"PCS_BACKEND_GO/router/calendarrouter"
	"PCS_BACKEND_GO/router/userrouter"
	"log"
)

func main() {

	// 設置日誌輸出配置

	f, err := os.OpenFile("mrmaster_log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		// test
		log.Fatal("OpenFile is failed")
	}

	defer f.Close()

	log.SetOutput(f)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println("test test ")

	// 加載路由
	routerinit.Include(calendarrouter.CalendarApi, userrouter.UserApi)
	r := routerinit.InitRouters()

	err = r.Run(global.IP + ":" + global.Port) // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Println("err ", err.Error())
	}
}
