package userrouter

import (
	"PCS_BACKEND_GO/global"
	"PCS_BACKEND_GO/global/database"
	"PCS_BACKEND_GO/model/user"
	"PCS_BACKEND_GO/service/userservice"
	"net/http"

	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func UserApi(e *gin.Engine) {
	UserGroup := e.Group("/api")
	{
		UserGroup.PUT("/user", putUser)
		UserGroup.POST("/login", login)
	}

}

// putUser 新增使用者
func putUser(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusForbidden, err)
		return
	}

	var user user.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userInfo, err := userservice.CreatUser(user, database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, userInfo)

}

func login(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusForbidden, err)
		return
	}

	var user user.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	userInfo, err := userservice.LoginUser(user, database.DB)
	if err != nil || userInfo.JwtToken == global.LoginError {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	c.JSON(http.StatusCreated, userInfo)
}
