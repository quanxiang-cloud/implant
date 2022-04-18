package v1beta1

import (
	"context"
	"net/http"

	"github.com/quanxiang-cloud/cabin/tailormade/client"
)

type Client struct {
	client *http.Client
	target string
}

type Function struct {
	Labels map[string]string `json:"labels"`
	State  string            `json:"state"`
}

func New(c *client.Config, target string) *Client {
	client := client.New(client.Config{
		Timeout:      c.Timeout,
		MaxIdleConns: c.MaxIdleConns,
	})

	return &Client{
		client: &client,
		target: target,
	}
}

type Resp struct {
}

func (c *Client) Send(ctx context.Context, fn Function) (*Resp, error) {
	resp := &Resp{}
	err := client.POST(ctx, c.client, c.target, fn, resp)
	return resp, err
}
