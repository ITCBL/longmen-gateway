package model

type ServiceTcpRule struct {
	BaseModel
	ServiceID int64 `json:"service_id" gorm:"type:bigint:not null" description:"服务id"`
	Port      int   `json:"port" gorm:"type:int:not null" description:"端口"`
}

func (t *ServiceTcpRule) TableName() string {
	return "service_tcp_rule"
}
