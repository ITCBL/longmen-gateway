package model

type ServiceLoadBalance struct {
	BaseModel
	ServiceID              int64  `json:"service_id" gorm:"type:bigint;not null" description:"服务id"`
	CheckMethod            bool   `json:"check_method" gorm:"type:bool;default:false" description:"检查方法 tcpchk=检测端口是否握手成功"`
	CheckTimeout           bool   `json:"check_timeout" gorm:"type:bool;default:false" description:"check超时时间"`
	CheckInterval          int    `json:"check_interval" gorm:"type:int;default:2" description:"检查间隔, 单位s"`
	RoundType              string `json:"round_type" gorm:"type:varchar(100);default:'round'" description:"轮询方式 round/weight_round/random/ip_hash"`
	IpList                 string `json:"ip_list" gorm:"type:varchar(1000);default:''" description:"ip列表"`
	WeightList             string `json:"weight_list" gorm:"type:varchar(1000);default:''" description:"权重列表"`
	ForbidList             string `json:"forbid_list" gorm:"type:varchar(1000);default:''" description:"禁用ip列表"`
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" gorm:"type:int;default:2" description:"下游建立连接超时, 单位s"`
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" gorm:"type:int;default:2" description:"下游获取header超时, 单位s"`
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" gorm:"type:int;default:2" description:"下游链接最大空闲时间, 单位s"`
	UpstreamMaxIdle        int    `json:"upstream_max_idle" gorm:"type:int;default:5" description:"下游最大空闲链接数"`
}

func (t *ServiceLoadBalance) TableName() string {
	return "service_load_balance"
}
