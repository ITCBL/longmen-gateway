package main

import (
	"fmt"
	"longmen-gateway/global"
	"longmen-gateway/initialize"
)

func main() {
	initialize.Logger()

	initialize.InitConfig()

	initialize.InitMysql()

	routers := initialize.InitRouters()

	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	port := fmt.Sprintf(":%s", global.ServerConfig.Port)
	routers.Run(port) // 监听并在 0.0.0.0:8080 上启动服务
}
