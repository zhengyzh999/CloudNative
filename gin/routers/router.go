package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouters(router *gin.Engine) {
	// 路由分组，可以根据模块或版本划分接口
	//course := router.Group("/course")
	//{
	//	course.GET("/", getCourse)
	//}
	//
	//user := router.Group("/user")
	//{
	//	user.GET("/", getUser)
	//}
	InitUser(router)
	InitCourse(router)
}

func getUser(context *gin.Context) {
	// 这里写获取数据或处理数据逻辑
	context.JSON(http.StatusOK, gin.H{
		"message": "get user",
	})
}
func getCourse(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "get course",
	})
}
