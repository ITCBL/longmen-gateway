package model

type ServiceGrpcRule struct {
	BaseModel
	ServiceID      int64  `json:"service_id" gorm:"type:bigint;not null" description:"服务id"`
	Port           int    `json:"port" gorm:"type:int;not null" description:"端口"`
	HeaderTransfor string `json:"header_transfor" gorm:"type:varchar(1000);default:''" description:"header转换支持增加(add)、删除(del)、修改(edit) 命令格式: add headname headvalue 命令示例：add token xxxxxx"`
}

func (t *ServiceGrpcRule) TableName() string {
	return "service_grpc_rule"
}
