package config

import "fmt"

type Redis struct {
	IP       string `json:"ip" yaml:"ip"`
	Port     int    `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"`
	PoolSize int    `json:"poolSize" yaml:"pool_size"`
}

func (r Redis) Addr() string {
	return fmt.Sprintf("%s:%d", r.IP, r.Port)
}
