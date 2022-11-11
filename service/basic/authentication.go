package basic

import (
	"PCS_BACKEND_GO/model/user"
	"log"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
)

var Secret []byte

var authUserList *authUserListStu

type authUserListStu struct {
	authUserArr []authUser
	m           *sync.Mutex
}

type authUser struct {
	userid int64
	jwtStr string
}



func (a *authUserListStu) setAuthUser(userid int64, jwtStr string) {
	a.m.Lock()
    defer a.m.Unlock()
    a.authUserArr = append(a.authUserArr, authUser{userid: userid, jwtStr: jwtStr})
}


func init() {

	// init authUserList
	authUserList = new(authUserListStu)

}

// User Auth function

// CreateAuthUser 建立用戶與token
func CreateAuthUser(userInfo user.User) (err error) {
	tokeString, err := generateToken(userInfo)
	if err != nil{
		log.Println("generateToken is failed: ", err)
		return
	}

	authUserList.setAuthUser(userInfo.ID, tokeString)
	return	
}

// CheckAuthUser




















// JWT function

type authClaims struct {
    jwt.StandardClaims
    UserID int64 `json:"userId"`
}

// generateToken 產生令牌
func generateToken(userInfo user.User)  (tokenString string, err error) {
    expiresAt := time.Now().Add(24 * time.Hour).Unix()
    token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
        StandardClaims: jwt.StandardClaims{
            Subject:   userInfo.UserName,
            ExpiresAt: expiresAt,
        },
        UserID: userInfo.ID,
    })
    tokenString, err = token.SignedString(Secret)
    if err != nil {
        return  
    }
    return  
}
