package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	StatusCode int
	Body       string
	Headers    http.Header
}

type APIClient struct {
	BaseURL string
	Client  *http.Client
}

type RequestOption struct {
	Endpoint string
	Headers  map[string]string
	Payload  []byte
}

func (c *APIClient) Send(method string, opts RequestOption) (*Response, error) {
	req, err := http.NewRequest(method, opts.Endpoint, bytes.NewBuffer(opts.Payload))
	if err != nil {
		return nil, err
	}

	// Add Headers
	for k, v := range opts.Headers {
		req.Header.Set(k, v)
	}

	// Send Request
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Read Response Body
	bodyBytes, _ := io.ReadAll(res.Body)

	fmt.Println("Response Status:", res.Status)
	return &Response{
		StatusCode: res.StatusCode,
		Body:       string(bodyBytes),
		Headers:    res.Header,
	}, nil
}

func main() {
	client := &APIClient{
		BaseURL: "https://swapi.dev/api",
		Client:  &http.Client{},
	}

	opts := RequestOption{
		Endpoint: client.BaseURL + "/people/1",
		Headers:  map[string]string{"Content-Type": "application/json"},
		// Payload:  []byte(``),
	}

	response, err := client.Send("GET", opts)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response Body:", response.Body)
}
