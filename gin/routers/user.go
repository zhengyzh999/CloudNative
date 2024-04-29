package routers

import (
	"CloudNative/gin/middleware"
	"CloudNative/gin/web"
	"github.com/gin-gonic/gin"
)

func InitUser(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		// 放在分组上的中间件，只针对分组内的请求执行
		v1.Use(middleware.Auth())
		v1.GET("/user", web.GetUser)
		v1.POST("/user", web.AddUser)
	}
}
