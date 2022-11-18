package userdao

import (
	"PCS_BACKEND_GO/model/user"

	"gorm.io/gorm"
)

// InsertUser 新增使用者
func InsertUser(user user.User, tx *gorm.DB) (userId int64, err error) {
	sqlStr := "INSERT INTO `user` (username, password) VALUES ( ?, ? ) RETURNING id;"

    err = tx.Raw(sqlStr).Scan(&userId).Error
	if tx.Error != nil {
		err = tx.Error
	}
	return
}

// GetUserByUsername
func GetUserByUsername(user user.User, tx *gorm.DB) (userArr []user.User, err error) {
	sqlStr := "SELECT * FROM user WHERE username = ?"

	err = tx.Raw(sqlStr, user.UserName).Scan(&userArr).Error
	if tx.Error != nil {
		err = tx.Error
	}

	return
}

// GetUserbyUsernameAnd 檢驗帳號密碼正確
func GetUserbyNameAndPassword(user user.User, tx *gorm.DB) (userInfo user.User, err error) {

	sqlStr := "SELECT * FROM user WHERE username = ? AND password = ?"
	err = tx.Raw(sqlStr, user.UserName, user.Password).Scan(&userInfo).Error
	if tx.Error != nil {
		err = tx.Error
	}

	return
}
