package model

type ServiceAccessControl struct {
	BaseModel
	ServiceID         int64  `json:"service_id" gorm:"type:bigint;not null" description:"服务id"`
	OpenAuth          bool   `json:"open_auth" gorm:"type:bool;default:false" description:"是否开启权限 true=开启"`
	BlackList         string `json:"black_list" gorm:"type:varchar(1000);default:''" description:"黑名单ip"`
	WhiteList         string `json:"white_list" gorm:"type:varchar(1000);default:''" description:"白名单ip"`
	WhiteHostName     string `json:"white_host_name" gorm:"type:varchar(255);default:''" description:"白名单主机名称"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" gorm:"type:int;default:0" description:"客户端ip限流"`
	ServiceFlowLimit  int    `json:"service_flow_limit" gorm:"type:int;default:0" description:"服务端限流"`
}

func (t *ServiceAccessControl) TableName() string {
	return "service_access_control"
}
