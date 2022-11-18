package calendarservice

import (
	"PCS_BACKEND_GO/dao/calendardao"
	"PCS_BACKEND_GO/model/calendar"
	"log"
	"strconv"

	"gorm.io/gorm"
)

func GetCalendarData(userIdstr string, db *gorm.DB) (calendarArr []calendar.Calendar) {
	tx := db.Begin()
	userId, err := strconv.ParseInt(userIdstr, 10, 64)
	if err != nil {
		log.Println(err)
	}

	calendarArr, err = calendardao.GetAllByUserId(userId, tx)
	if err != nil {
		log.Println(err)
		return

	}
	return
}

// CreatCalendarData 建立日曆任務
func CreatCalendarData(calendarData calendar.Calendar, db *gorm.DB) (err error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
		if err != nil {
			log.Println(err)
		}
	}()
	err = calendardao.InsertData(calendarData, tx)
	if err != nil {
		panic(err)
	}
	tx.Commit()
	return
}

// DeleteCalendarData 刪除指定日曆任務
func DeleteCalendarData(idStr string, db *gorm.DB) {
	tx := db.Begin()
	var err error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
		if err != nil {
			log.Println(err)
		}
	}()
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println(err)
		return

	}
	err = calendardao.DeleteDataById(id, tx)
	if err != nil {
		panic(err)
	}
	tx.Commit()
}
