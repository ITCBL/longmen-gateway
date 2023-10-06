package initialize

import (
	"encoding/json"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"

	"longmen-gateway/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	local := os.Getenv("local")
	if local != "false" {
		configFilePath := "config/config-dev.yaml"
		v := viper.New()
		v.SetConfigFile(configFilePath)
		if err := v.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := v.Unmarshal(global.ServerInfo); err != nil { // 通过全局变量使用配置信息
			panic(err)
		}
	} else {
		global.ServerInfo.Port = os.Getenv("app_port") // 从运行环境中获取配置信息
		global.ServerInfo.MysqlInfo.Host = os.Getenv("mysql_host")
		global.ServerInfo.MysqlInfo.Port = os.Getenv("mysql_port")
		global.ServerInfo.MysqlInfo.Name = os.Getenv("mysql_name")
		global.ServerInfo.MysqlInfo.User = os.Getenv("mysql_user")
		global.ServerInfo.MysqlInfo.Password = os.Getenv("mysql_password")
		// ......
	}

	marshal, err := json.Marshal(global.ServerInfo)
	if err != nil {
		panic(err.Error())
	}

	zap.S().Infof("[System configuration information]:\n %v", string(marshal))
}
