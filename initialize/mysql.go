package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"longmen-gateway/global"
	"longmen-gateway/model"
	"os"
	"time"
)

func InitMysql() {
	mysqlConfig := global.ServerConfig.MysqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Name)
	zap.S().Info("dsn:  ", dsn)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true, // Disable color printing
		},
	)

	var err error
	// global
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Custom table name
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	global.DB.AutoMigrate(
		&model.App{},
		&model.ServiceAccessControl{},
		&model.ServiceGrpcRule{},
		&model.ServiceHttpRule{},
		&model.ServiceInfo{},
		&model.ServiceLoadBalance{},
		&model.ServiceTcpRule{},
		&model.User{},
	)
	zap.S().Info("mysql connect success:", global.DB)
}
