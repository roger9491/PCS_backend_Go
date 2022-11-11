package sqlinit

import (
	"PCS_BACKEND_GO/global/database"
	"PCS_BACKEND_GO/model/calendar"
	"database/sql"
	"fmt"
	"log"
	"os/user"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	creatDataBase()
	database.DB = initMySQL()
}

// 建立資料庫
func creatDataBase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		database.Username,
		database.Password,
		database.Host,
		database.Port)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("無法打開 mysql", err.Error())
	}
	_, err = db.Exec("CREATE DATABASE "+ database.DBName)
	if err != nil {
		return
	}

}

func initMySQL() *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		database.Username, database.Password, database.Host, database.Port, database.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("連接數據庫失敗, err: ", err.Error())
	}

	db.AutoMigrate(&calendar.Calendar{}, &user.User{})

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
