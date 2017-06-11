package camera

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const api = "http://192.168.122.1:8080/sony/camera"

// Camera Remote API
type Camera struct {
	client Client
}

// Action sends Request to Camera
func (c *Camera) Action(req Request) (Response, error) {
	var res Response
	bs, err := json.Marshal(&req)
	if err != nil {
		return res, err
	}

	reader, err := c.client.Send(bytes.NewReader(bs))
	if err != nil {
		return res, err
	}

	defer reader.Close()
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&res)
	if err != nil {
		return res, err
	}

	if len(res.Error) == 2 {
		return res, fmt.Errorf("response: %.f %s", res.Error[0], res.Error[1])
	}

	return res, nil
}

// NewCamera with a client and default API URL
func NewCamera() Camera {
	return Camera{Client{api}}
}

// Response from Camera
type Response struct {
	ID     int           `json:"id"`
	Result []interface{} `json:"result"`
	Error  []interface{} `json:"error"`
}

// Params of Request
type Params []interface{}

// Request to Camera
type Request struct {
	ID      int         `json:"id"`
	Version string      `json:"version"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

// NewRequest ...
func NewRequest(method string) Request {
	params := make(Params, 0)
	return NewParamsRequest(method, params)
}

// NewParamsRequest ...
func NewParamsRequest(method string, params interface{}) Request {
	return Request{id, version, method, params}
}

// Client for HTTP calls
type Client struct {
	URL string
}

// Send request
func (c *Client) Send(reader io.Reader) (io.ReadCloser, error) {
	req, err := http.NewRequest(http.MethodPost, api, reader)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		defer res.Body.Close()
		return nil, fmt.Errorf("unexpected HTTP response: status %s code %d", res.Status, res.StatusCode)
	}

	return res.Body, nil
}
