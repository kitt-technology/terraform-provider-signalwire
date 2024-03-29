package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Config struct {
	Token     string
	ProjectId string
}

func (c *Config) Client() (interface{}, error) {
	return &Client{
		Config:     c,
		httpClient: http.DefaultClient,
	}, nil
}

type Client struct {
	Config     *Config
	BaseUrl    string
	httpClient *http.Client
}

func (c *Client) Req(method, space, uri string, payload map[string]interface{}) (map[string]interface{}, error) {
	var reader io.Reader
	var jsonStr []byte
	var err error
	if payload != nil {
		jsonStr, err = json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewBuffer(jsonStr)
	}

	url := "https://"+space+".signalwire.com/api/relay/rest/"+uri
	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Config.ProjectId, c.Config.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var jsonResp map[string]interface{}

	if len(body) > 0 {
		err = json.Unmarshal(body, &jsonResp)
		if err != nil {
			return nil, fmt.Errorf(
				`could not unmarshall: %s response from %s with payload of %s`,
				string(body),
				url,
				string(jsonStr),
			)
		}
	}

	return jsonResp, err
}
