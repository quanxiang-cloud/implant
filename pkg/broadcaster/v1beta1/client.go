package v1beta1

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/quanxiang-cloud/cabin/tailormade/client"
	"github.com/quanxiang-cloud/cabin/tailormade/header"
	"github.com/quanxiang-cloud/cabin/tailormade/resp"
)

type Client struct {
	client *http.Client
	target string
}

type Function struct {
	State       string `json:"state"`
	ResourceRef string `json:"resourceRef"`
	Topic       string `json:"topic"`
	Name        string `json:"name"`
}

type Pipeline struct {
	Name  string `json:"name"`
	Topic string `json:"topic"`
	State string `json:"state"`
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

type Resp = resp.Resp

func (c *Client) Send(ctx context.Context, params interface{}) (*Resp, error) {
	resp := &Resp{}
	if err := POST(ctx, c.client, c.target, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// POST http post
func POST(ctx context.Context, client *http.Client, uri string, params interface{}, entity interface{}) error {
	if reflect.ValueOf(entity).Kind() != reflect.Ptr {
		return errors.New("the entity type must be a pointer")
	}

	paramByte, err := json.Marshal(params)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(paramByte)
	req, err := http.NewRequest("POST", uri, reader)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(header.GetRequestIDKV(ctx).Wreck())
	req.Header.Add(header.GetTimezone(ctx).Wreck())
	req.Header.Add(header.GetTenantID(ctx).Wreck())

	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("expected state value is 200, actually %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, entity)
}
