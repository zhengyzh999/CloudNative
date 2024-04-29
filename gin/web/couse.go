package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCourse(context *gin.Context) {
	// 这里写获取数据或处理数据逻辑
	id := context.DefaultQuery("id", "0")
	context.JSON(http.StatusOK, gin.H{
		"message": "已通过get方法请求体传参，获取全部id为" + id + "的课程",
	})
}

func GetCourseParam(context *gin.Context) {
	// 这里写获取数据或处理数据逻辑
	id := context.Param("id")
	context.JSON(http.StatusOK, gin.H{
		"message": "已通过get方法url传参，获取全部id为" + id + "的课程",
	})
}

func AddCourse(context *gin.Context) {
	course := &Course{}
	// Bind：出错后直接响应请求。ShouldBind：返回一个错误，可以开发者根据返回的错误自定义响应数据
	err := context.BindJSON(course)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	course.Name = course.Name + "后端处理过了"
	course.Id = course.Id + 10
	context.JSON(http.StatusOK, course)
}

type Course struct {
	Id      int     `json:"id" binding:"required"`
	Name    string  `json:"name" binding:"required"`
	Teacher string  `json:"teacher"`
	Price   float64 `json:"price" binding:"number"`
}
