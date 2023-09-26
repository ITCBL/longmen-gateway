package initialize

import (
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
		configFileName := "config/config-dev.yaml"
		v := viper.New()
		//文件的路径如何设置
		v.SetConfigFile(configFileName)
		if err := v.ReadInConfig(); err != nil {
			panic(err)
		}
		//这个对象如何在其他文件中使用 - 全局变量
		if err := v.Unmarshal(global.ServerConfig); err != nil {
			panic(err)
		}
	} else {
		// 从环境变量中获取
		global.ServerConfig.Port = os.Getenv("app_port")
		global.ServerConfig.MysqlConfig.Host = os.Getenv("mysql_host")
		global.ServerConfig.MysqlConfig.Port = os.Getenv("mysql_port")
		global.ServerConfig.MysqlConfig.Name = os.Getenv("mysql_name")
		global.ServerConfig.MysqlConfig.User = os.Getenv("mysql_user")
		global.ServerConfig.MysqlConfig.Password = os.Getenv("mysql_password")
	}
	zap.S().Infof("配置信息: %v", global.ServerConfig)
}
