package main

import (
	"fmt"
	"log"
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
	service.InitWenXin()
	// 创建Gin引擎
	router := gin.Default()

	router.POST("/send_msg", service.SendMsgHandler)
	// 启动HTTP服务器
	log.Fatal(router.Run(":80"))

}
