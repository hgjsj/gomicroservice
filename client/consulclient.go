package client

import (
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
)

func NewConsulClient(service string, port int, tags []string, logger log.Logger) *consul.Instancer {
	var client consul.Client

	consulConfig := api.DefaultConfig()

	consulConfig.Address = fmt.Sprintf(":%d", port)

	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}
	client = consul.NewClient(consulClient)

	return consul.NewInstancer(client, logger, service, tags, true)
}
