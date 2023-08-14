package config

type Database struct {
	Type   string `mapstructure:"type" json:"type" yaml:"type"` // 数据库类型
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Sqlite Sqlite `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
}

type Mysql struct {
	Path     string `mapstructure:"path" json:"path" yaml:"path"`             // 服务器地址：端口
	Port     string `mapstructure:"port" json:"port" yaml:"port"`             //:端口
	Dbname   string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`    // 数据库名
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 数据库用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
}

type Sqlite struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"` // sqlite 路径
}
