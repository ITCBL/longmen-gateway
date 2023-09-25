package model

type App struct {
	BaseModel
	AppID    string `json:"app_id" gorm:"column:app_id;index:idx_app_id;unique;type:varchar(100);not null" description:"租户id"`
	Name     string `json:"name" gorm:"column:name;type:varchar(20);not null" description:"租户名称"`
	Secret   string `json:"secret" gorm:"column:secret;type:varchar(20);not null" description:"密钥"`
	WhiteIPS string `json:"white_ips" gorm:"column:white_ips;varchar(1000)" description:"ip白名单，支持前缀匹配"`
	Qpd      int64  `json:"qpd" gorm:"column:qpd;type:bigint;default:0" description:"每日请求量限制"`
	Qps      int64  `json:"qps" gorm:"column:qps;type:bigint;default:0" description:"每秒请求量限制"`
}

func (t *App) TableName() string {
	return "app"
}
