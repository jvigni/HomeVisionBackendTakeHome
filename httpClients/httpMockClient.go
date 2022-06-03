package httpClients

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type FakeRequest struct {
	url          string
	statusCode   int
	responseBody string
}

type HttpMockClient struct {
	fakeRequestsByUrl map[string]FakeRequest
}

func NewHttpMockClient() *HttpMockClient {
	return &HttpMockClient{fakeRequestsByUrl: make(map[string]FakeRequest)}
}

func (c *HttpMockClient) Simulate(url string, statusCode int, responseBody string) {
	newFakeRequest := FakeRequest{ url, statusCode, responseBody }
	c.fakeRequestsByUrl[url] = newFakeRequest
}

func (c *HttpMockClient) Get(url string) (*http.Response, error) {
	val, ok := c.fakeRequestsByUrl[url]
	if !ok {
		return nil, fmt.Errorf("can't find mock request for %s", url)
	}
	return &http.Response{
		StatusCode: val.statusCode,
        Body: ioutil.NopCloser(bytes.NewBufferString(val.responseBody)),
    }, nil
}

