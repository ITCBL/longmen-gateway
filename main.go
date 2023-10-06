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

	initialize.InitTrans("zh") // 翻译器

	initialize.RegisterValidation() // 注册验证器

	port := fmt.Sprintf(":%s", global.ServerInfo.Port)
	routers.Run(port) // 监听并启动服务 todo 后续需要启动两个端口，区分manager和proxy
}
