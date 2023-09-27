package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"longmen-gateway/global"
	"longmen-gateway/initialize"
)

func main() {
	initialize.Logger()

	initialize.InitConfig()

	initialize.InitMysql()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := fmt.Sprintf(":%s", global.ServerConfig.Port)
	r.Run(port) // 监听并在 0.0.0.0:8080 上启动服务
}
