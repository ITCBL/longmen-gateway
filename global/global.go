package global

import (
	ut "github.com/go-playground/universal-translator"
	"gorm.io/gorm"

	"longmen-gateway/config"
)

var (
	DB           *gorm.DB
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	Trans        ut.Translator
	//MongoClient  *qmgo.Database
)
