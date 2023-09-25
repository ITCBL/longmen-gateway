package model

type ServiceHttpRule struct {
	BaseModel
	ServiceID      int64  `json:"service_id" gorm:"type:bigint;not null" description:"服务id"`
	RuleType       int    `json:"rule_type" gorm:"type:int;not null" description:"匹配类型 domain=域名, url_prefix=url前缀"`
	Rule           string `json:"rule" gorm:"type:varchar(1000);not null" description:"type=domain表示域名，type=url_prefix时表示url前缀"`
	NeedHttps      bool   `json:"need_https" gorm:"type:bool;default:false" description:"true支持https，false不支持https"`
	NeedWebsocket  int    `json:"need_websocket" gorm:"type:int;not null" description:"启用websocket true=启用"`
	NeedStripUri   int    `json:"need_strip_uri" gorm:"type:int;not null" description:"启用strip_uri true=启用"`
	UrlRewrite     string `json:"url_rewrite" gorm:"type:varchar(1000);default:''" description:"url重写功能，每行一个"`
	HeaderTransfor string `json:"header_transfor" gorm:"type:varchar(1000);default:''" description:"header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue"`
}

func (t *ServiceHttpRule) TableName() string {
	return "service_http_rule"
}
