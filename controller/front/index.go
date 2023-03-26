package front

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"renWoXing/service/labelsService"
)

type lab struct {
	Icon     string    `json:"icon"`
	Position []float32 `json:"position"`
}

func Index(c *gin.Context) {

	lab := labelsService.Labels{}
	labs, _ := lab.Get()

	c.HTML(200, "mapIndex.tmpl", gin.H{
		"title": "任我行",
		"labs":  labs,
	})
}

func GetLabs(c *gin.Context) {
	var labs []lab
	lab1 := lab{Icon: "//127.0.0.1:8080/static/images/01.png", Position: []float32{116.205467, 39.907761}}
	labs = append(labs, lab1)
	fmt.Println("labs = ", labs)
	c.JSON(200, gin.H{"labs": labs})
}
