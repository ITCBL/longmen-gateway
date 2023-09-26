package config

type ServerConfig struct {
	Port        string `mapstructure:"port" json:"port"`
	MysqlConfig Mysql  `mapstructure:"mysql" json:"mysql"`
	RedisConfig Redis  `mapstructure:"redis" json:"redis"`
}

type Mysql struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type Redis struct {
	Uri      string `mapstructure:"uri" json:"uri"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}
