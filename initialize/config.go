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
		// Using configuration information through global variables
		if err := v.Unmarshal(global.ServerConfig); err != nil {
			panic(err)
		}
	} else {
		// Obtain configuration information from the running environment
		global.ServerConfig.Port = os.Getenv("app_port")
		global.ServerConfig.MysqlConfig.Host = os.Getenv("mysql_host")
		global.ServerConfig.MysqlConfig.Port = os.Getenv("mysql_port")
		global.ServerConfig.MysqlConfig.Name = os.Getenv("mysql_name")
		global.ServerConfig.MysqlConfig.User = os.Getenv("mysql_user")
		global.ServerConfig.MysqlConfig.Password = os.Getenv("mysql_password")
		// ......
	}

	marshal, err := json.Marshal(global.ServerConfig)
	if err != nil {
		panic(err.Error())
	}

	zap.S().Infof("[System configuration information]:\n %v", string(marshal))
}
