package router

import (
	"server/controller/api"

	"github.com/gin-gonic/gin"
)

func StartRouter() *gin.Engine {
	router := gin.Default()

	// 加载静态文件

	// router.LoadHTMLGlob("static/index.html")
	// router.Static("/css", "./static/css")
	// router.Static("/js", "./static/js")
	// router.StaticFile("/favicon.ico", "./static/favicon.ico")
	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 	})
	// })

	// 开启跨域请求

	// router.Use(Cors())

	router.GET("/api/GetMemberInfo", api.GetMemberInfo)
	// router.GET("/api/add", api.Add)
	// router.GET("/api/checkHealthy",api.CheckHealthy)
	// router.GET("/api/testFunc",api.TestFunc)

	// router.POST("/api/createOrder",api.CreateOrder)

	return router
}
