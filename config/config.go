package config

type ServerConfig struct {
	Port      string      `mapstructure:"port" json:"port"`
	MysqlInfo MysqlConfig `mapstructure:"mysql" json:"mysql"`
	RedisInfo RedisConfig `mapstructure:"redis" json:"redis"`
	JWTInfo   JWTConfig   `mapstructure:"jwt" json:"jwt"`
	LogInfo   LogConfig   `mapstructure:"log" json:"log"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type RedisConfig struct {
	Host   string `mapstructure:"host" json:"host"`
	Port   int    `mapstructure:"port" json:"port"`
	Secret string `mapstructure:"secret" json:"secret"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type LogConfig struct {
	EnableSuccessResponseLog bool `mapstructure:"enable_success_response_log" json:"enable_success_response_log"`
}
