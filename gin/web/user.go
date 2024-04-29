package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(context *gin.Context) {
	// 这里写获取数据或处理数据逻辑
	id := context.DefaultQuery("id", "0")
	context.JSON(http.StatusOK, gin.H{
		"message": "已获取id为" + id + "的user信息",
	})
}

func AddUser(context *gin.Context) {
	user := &User{}
	// Bind：出错后直接响应请求。ShouldBind：返回一个错误，可以开发者根据返回的错误自定义响应数据
	err := context.BindJSON(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	user.Name = user.Name + "后端处理过了"
	user.Id = user.Id + 10
	context.JSON(http.StatusOK, user)
}

type User struct {
	Id    int     `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Phone string  `json:"phone" binding:"required,e164"`
	Price float64 `json:"price" binding:"number"`
}
