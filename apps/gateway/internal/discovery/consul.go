package discovery

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type ConsulClient struct {
	client *api.Client
}

func NewConsul(addr string) (*ConsulClient, error) {
	cfg := api.DefaultConfig()
	cfg.Address = addr
	c, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &ConsulClient{client: c}, nil
}

// 获取服务所有实例的地址列表
func (c *ConsulClient) GetServiceAddresses(serviceName string) ([]string, error) {
	services, _, err := c.client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var addrs []string
	for _, s := range services {
		addr := fmt.Sprintf("%s:%d", s.Service.Address, s.Service.Port)
		addrs = append(addrs, addr)
	}

	// 如果没有找到指定协议的服务，返回错误
	if len(addrs) == 0 {
		return nil, fmt.Errorf("no healthy service: %s", serviceName)
	}

	return addrs, nil
}
