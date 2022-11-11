package userdao

import (
	"PCS_BACKEND_GO/model/user"

	"gorm.io/gorm"
)

// InsertUser 新增使用者
func InsertUser(user user.User, tx *gorm.DB) {
	result := tx.Create(user)
	if result.Error != nil {
		panic(result.Error)
	}
}

// GetUserbyUsernameAnd 新增使用者
func GetUserbyUsernameAnd(user user.User, tx *gorm.DB) (userInfo user.User){

	sqlStr := "SELECT * FROM user WHERE username = ? AND password = ?"
	row := tx.Raw(sqlStr, user.UserName, user.Password).Row()
	if row.Err()!= nil {
        panic(row.Err())
	}	
	row.Scan(&userInfo)

	return	
}