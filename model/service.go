package model

// ServiceDetail 不映射数据库表结构，用来汇总信息。
type ServiceDetail struct {
	Info          *ServiceInfo          `json:"info" description:"基本信息"`
	HTTPRule      *ServiceHttpRule      `json:"http_rule" description:"http拦截规则"`
	TCPRule       *ServiceTcpRule       `json:"tcp_rule" description:"tcp拦截规则"`
	GRPCRule      *ServiceGrpcRule      `json:"grpc_rule" description:"grpc拦截规则"`
	LoadBalance   *ServiceLoadBalance   `json:"load_balance" description:"负载均衡策略"`
	AccessControl *ServiceAccessControl `json:"access_control" description:"权限控制策略"`
}
