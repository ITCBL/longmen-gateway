package model

type ServiceInfo struct {
	BaseModel
	LoadType    int    `json:"load_type" gorm:"type:int;not null" description:"负载类型 0=http 1=tcp 2=grpc"`
	ServiceName string `json:"service_name" gorm:"type:varchar(20);not null" description:"服务名称"`
	ServiceDesc string `json:"service_desc" gorm:"type:varchar(255);not null" description:"服务描述"`
}

func (t *ServiceInfo) TableName() string {
	return "service_info"
}
