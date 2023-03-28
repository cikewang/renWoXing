package front

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"renWoXing/service/actionService"
	"renWoXing/service/labelsService"
	"time"
)

type lab struct {
	Icon     string    `json:"icon"`
	Position []float32 `json:"position"`
}

func Index(c *gin.Context) {

	lab := labelsService.Labels{}
	labs, _ := lab.Get()

	act := actionService.Action{}
	actionInfo, _ := act.GetUpdateTime()
	c.HTML(200, "mapIndex.tmpl", gin.H{
		"title":         "任我行",
		"labs":          labs,
		"updateTimeStr": time.Unix(int64(actionInfo.UpdateTime), 0).Format("2006-01-02 15:04"),
	})
}

func GetLabs(c *gin.Context) {
	var labs []lab
	lab1 := lab{Icon: "//127.0.0.1:8080/static/images/01.png", Position: []float32{116.205467, 39.907761}}
	labs = append(labs, lab1)
	fmt.Println("labs = ", labs)
	c.JSON(200, gin.H{"labs": labs})
}
