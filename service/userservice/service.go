package userservice

import (
	"PCS_BACKEND_GO/authentication"
	"PCS_BACKEND_GO/dao/userdao"
	"PCS_BACKEND_GO/global"
	"PCS_BACKEND_GO/model/user"
	"log"

	"gorm.io/gorm"
)

const (
	loginError = "-1"
)

// CreatUser 建立使用者
func CreatUser(user user.User, db *gorm.DB) (userInfo user.UserInfo, err error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
		if err != nil {
			log.Println(err)
		}
	}()

	// 檢查使用者是否以創建
	userArr, err := userdao.GetUserByUsername(user, tx)
	if err != nil {
		panic(err)
	}
	if len(userArr) != 0 {
		userInfo.JwtToken = global.LoginError
		return
	}

	userInfo.UserID, err = userdao.InsertUser(user, tx)
	if err != nil {
		panic(err)
	}
	tx.Commit()
	// 產生令牌
	userInfo.JwtToken, err = authentication.GenerateToken(user)
	if err != nil {
		userInfo.JwtToken = global.LoginError
		return
	}
	return

}

// LoginUser 登入使用者
func LoginUser(user user.User, db *gorm.DB) (userInfo user.UserInfo, err error) {
	tx := db.Begin()
	userTmp, err := userdao.GetUserbyNameAndPassword(user, tx)
	if err != nil {
		log.Println(err)
		return
	}

	if userTmp.UserName == user.UserName && userTmp.Password == user.Password {
		// 帳號密碼符合

		// 產生 jwt
		userInfo.JwtToken, err = authentication.GenerateToken(user)
		if err != nil {
			log.Println(err)
			return
		}

	} else {
		userInfo.JwtToken = global.LoginError
	}
	userInfo.UserID = userTmp.ID
	return
}
