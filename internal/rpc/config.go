package rpc

import "fmt"

type RpcConfig struct {
	Host string ``
	Port string ``
}

func (c RpcConfig) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
