package global

import (
	ut "github.com/go-playground/universal-translator"
	"gorm.io/gorm"

	"longmen-gateway/config"
)

var (
	DB         *gorm.DB
	ServerInfo *config.ServerConfig = &config.ServerConfig{}
	Trans      ut.Translator
	//MongoClient  *qmgo.Database
)

const (
	UserSessionKey         = "UserSessionKey"
	Issuer                 = "longmen-gateway"
	JwtTokenExpirationTime = 60 * 60 * 24 * 30
)
