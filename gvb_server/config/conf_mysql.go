package config

import "strconv"

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Config   string `yaml:"config"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"logLevel"` //日志等级,debug就是输出全部
}

func (m *Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
}
