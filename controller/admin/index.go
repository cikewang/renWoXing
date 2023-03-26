package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {

	fmt.Println("labelsService index")

	c.HTML(http.StatusOK, "adminIndex.html", gin.H{
		"title": "labelsService website",
	})

}
