package main

import (
	"CloudNative/logrus/logs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var log *logs.Log

func init() {
	conf := logs.LogConf{
		Level:       logrus.InfoLevel,
		AdapterName: "fileRotate",
	}
	log = logs.InitLog(conf)
}

func main() {
	r := gin.New()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.Use(myLogger)
	//r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	r.Run(":8080")
}

func myLogger(ctx *gin.Context) {
	start := time.Now()
	path := ctx.Request.URL.Path
	raw := ctx.Request.URL.RawQuery
	ctx.Next()
	mp := make(map[string]interface{})
	mp["path"] = path
	mp["client_ip"] = ctx.ClientIP()
	mp["method"] = ctx.Request.Method
	mp["statusCode"] = ctx.Writer.Status
	mp["size"] = ctx.Writer.Size()
	if raw != "" {
		mp["path"] = path + "?" + raw
	}
	mp["latency"] = time.Since(start).String()
	log.WithFields(mp).Info()
}

func main1() {
	defer func() {
		log.Flush()
	}()

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 100,
		// A common pattern is to re-use fields between logging statements by re-using
		// the logrus.Entry returned from WithFields()
	}).Fatal("The ice breaks!")

	contextLogger := log.WithFields(logrus.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")

}
