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

// Execute command on camera
func (c *Camera) Execute(cmd Command) (map[string]interface{}, error) {
	reader, err := c.client.Send(cmd)
	if err != nil {
		return nil, err
	}

	return c.read(reader)
}

func (c *Camera) read(reader io.ReadCloser) (map[string]interface{}, error) {
	defer reader.Close()

	var result map[string]interface{}
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&result)
	if err != nil {
		return nil, err
	}

	if err, ok := result["error"]; ok {
		return nil, fmt.Errorf("unexpected error: %v", err)
	}

	return result, nil
}

// NewCamera camera with a client and default API URL
func NewCamera() Camera {
	return Camera{Client{api}}
}

// Client for HTTP calls
type Client struct {
	URL string
}

// Send Command
func (c *Client) Send(cmd Command) (io.ReadCloser, error) {
	reader, err := reader(cmd)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, api, reader)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

func (c *Client) do(req *http.Request) (io.ReadCloser, error) {
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

func reader(c Command) (io.Reader, error) {
	body, err := json.Marshal(&c)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(body), nil
}
