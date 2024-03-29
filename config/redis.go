package config

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"` // 服务器地址：端口
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
