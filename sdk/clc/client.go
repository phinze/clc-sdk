package clc

import (
	"github.com/mikebeyer/clc-sdk/sdk/api"
	"github.com/mikebeyer/clc-sdk/sdk/server"
	"github.com/mikebeyer/clc-sdk/sdk/status"
)

type Client struct {
	client *api.Client

	Server *server.Service
	Status *status.Service
}

func New(config api.Config) *Client {
	c := &Client{
		client: api.New(config),
	}

	c.Server = server.New(c.client)
	c.Status = status.New(c.client)
	return c
}