package routers

import (
	"github.com/gin-gonic/gin"
	"renWoXing/controller/admin"
	"renWoXing/controller/front"
)

func Router() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("./templates/**/*")
	r.Static("/static", "./static")

	// 前端路由
	frontGroup := r.Group("/front")
	{
		frontGroup.GET("/index", front.Index)
		frontGroup.GET("/labs", front.GetLabs)
	}

	// 后台管理路由
	adminGroup := r.Group("/admin")
	{
		adminGroup.GET("/page", admin.Index)
	}

	return r
}
