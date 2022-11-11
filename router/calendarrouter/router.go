package calendarrouter

import "github.com/gin-gonic/gin"

func CalendarApi(e *gin.Engine){
	calendarGroup := e.Group("/api")
	{
		calendarGroup.GET("/Calendar",getCalendar)
	}
	
}


func getCalendar(c *gin.Context) {
	

}