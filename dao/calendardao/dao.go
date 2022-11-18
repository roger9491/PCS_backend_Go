package calendardao

import (
	"PCS_BACKEND_GO/model/calendar"
	"fmt"

	"gorm.io/gorm"
)

// GetAll 獲取指定id所有日曆資料
func GetAllByUserId(userId int64, tx *gorm.DB) (calendarList []calendar.Calendar, err error) {
	sqlStr := "SELECT * FROM calendar WHERE userid = ?"

	err = tx.Raw(sqlStr, userId).Scan(&calendarList).Error
	if tx.Error != nil {
		err = tx.Error
	}

	return
}

// InsertData 新增任務
func InsertData(data calendar.Calendar, tx *gorm.DB) (err error){
	sqlStr := "INSERT INTO `calendar` (date, title, content, userid) VALUES ( ?, ?, ?, ? );"
	fmt.Println(sqlStr, data.UserId)
	err = tx.Exec(sqlStr, data.Date, data.Title, data.Content, data.UserId).Error
	if tx.Error != nil {
		err = tx.Error
		fmt.Println(err)
	}
	return
}


// DeleteDataById 刪除指定id任務
func DeleteDataById(id int64, tx *gorm.DB) (err error){
	sqlStr := "DELETE FROM `calendar` WHERE id = ? "

	err = tx.Exec(sqlStr, id).Error
	if tx.Error != nil {
		err = tx.Error
	}
	return
}
