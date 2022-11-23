package runprogramrouter

import (
	"PCS_BACKEND_GO/global/database"
	"PCS_BACKEND_GO/service/runprogramservice"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunprogramApi(e *gin.Engine) {
	RunprogramGroup := e.Group("/api")
	// RunprogramGroup.Use(authentication.AuthRequired)
	{
		RunprogramGroup.POST("/runprogram", runprogram)
	}

}

func runprogram(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusForbidden, err)
		return
	}
	fmt.Println(string(body))
	output, err := runprogramservice.RunProgram(string(body), database.DB)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.String(http.StatusCreated, output)
}
