package main

import (
	"fmt"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}
	// 初始化本地缓存
	service.InitLocalCache()

	// 创建Gin引擎
	router := gin.Default()

	router.POST("/send_msg", service.SendMsgHandler)
	// 启动HTTP服务器
	router.Run(":8080")
	// http.HandleFunc("/", service.IndexHandler)
	// http.HandleFunc("/api/count", service.CounterHandler)
	// http.HandleFunc("/send_msg", service.SendMsgHandler)

	// log.Fatal(http.ListenAndServe(":80", nil))
}
