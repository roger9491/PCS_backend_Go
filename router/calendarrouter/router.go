package calendarrouter

import (
	"PCS_BACKEND_GO/authentication"
	"PCS_BACKEND_GO/global/database"
	"PCS_BACKEND_GO/model/calendar"
	"PCS_BACKEND_GO/service/calendarservice"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CalendarApi(e *gin.Engine) {
	calendarGroup := e.Group("/api")
	calendarGroup.Use(authentication.AuthRequired)
	{
		calendarGroup.GET("/calendar", getCalendar)
		calendarGroup.POST("/calendar", postCalendar)
		calendarGroup.DELETE("/calendar", deleteCalendar)
	}

}

func getCalendar(c *gin.Context) {
	userId := c.Request.Header.Get("UserID")
	calendarData := calendarservice.GetCalendarData(userId, database.DB)
	c.JSON(http.StatusOK, calendarData)
}

func postCalendar(c *gin.Context) {
	userId := c.Request.Header.Get("UserID")
	fmt.Println(userId)
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusForbidden, err)
		return
	}

	var calendarData calendar.Calendar
	err = json.Unmarshal(body, &calendarData)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userIdInt64, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	calendarData.UserId = userIdInt64
	err = calendarservice.CreatCalendarData(calendarData, database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func deleteCalendar(c *gin.Context) {
	idStr := c.Query("id") 
	calendarservice.DeleteCalendarData(idStr, database.DB)
	c.JSON(http.StatusCreated, nil)
}
