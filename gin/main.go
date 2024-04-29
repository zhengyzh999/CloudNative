package main

import (
	"CloudNative/gin/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	initDB("tcp://localhost:3306")
	LogMiddleware(initDB)("tcp://localhost:3306")
	//r := gin.Default()
	r := gin.New()
	// AuthCheck, AuthCheck2()区别是闭包,后者方便传参到中间件中。这里使用的中间件，会对所有请求进行执行
	r.Use(gin.Logger(), gin.Recovery(), AuthCheck, AuthCheck2("abcd1234"))
	routers.InitRouters(r)
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.GET("/api/user", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"message": "get user",
	//	})
	//})
	//r.POST("/api/user", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"message": "add user",
	//	})
	//})
	//r.PUT("/api/user", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"message": "update user",
	//	})
	//})
	r.Run(":8080")
}

// AuthCheck 自定义中间件
func AuthCheck(c *gin.Context) {
	fmt.Println("call AuthCheck func")
	c.Next()
}
func AuthCheck2(param string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("call AuthCheck2 func", param)
		c.Next()
	}
}

// 简单使用中间件实现功能扩展
func initDB(connStr string) {
	fmt.Println("init db completed", connStr)
}

func LogMiddleware(initFunc func(string)) func(string) {
	return func(connStr string) {
		log.Println("start...")
		initFunc(connStr)
		log.Println("end...")
	}
}
