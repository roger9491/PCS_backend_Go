package userservice

import (
	"PCS_BACKEND_GO/dao/userdao"
	"PCS_BACKEND_GO/model/user"

	"gorm.io/gorm"
)

func CreatUser(user user.User, db *gorm.DB) {
	tx := db.Begin()
	defer func ()  {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	userdao.InsertUser(user, db)
	tx.Commit()

}
