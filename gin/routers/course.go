package routers

import (
	"CloudNative/gin/web"
	"github.com/gin-gonic/gin"
)

func InitCourse(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		// get方法请求体传参
		v1.GET("/course", web.GetCourse)
		// get方法url路径传参
		v1.GET("/course/:id", web.GetCourseParam)
		v1.POST("/course", web.AddCourse)
		v1.PUT("/course")
	}
}
