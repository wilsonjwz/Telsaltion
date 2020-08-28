package httpclinet

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const defaultTimeOut = 30 * time.Second

type Option struct {
	Header  map[string]string
	Timeout time.Duration
}

type HttpClient struct {
	Url    string
	Client *http.Client
	Option *Option
}

func NewHttpClient(url string, opt *Option) *HttpClient {
	client := &HttpClient{
		Url:    url,
		Client: &http.Client{},
	}
	if opt == nil {
		client.Option = &Option{
			Timeout: defaultTimeOut,
			Header: map[string]string{
				"Content-Type": "application/json",
			},
		}
	} else {
		client.Option = opt
	}
	return client
}

func (c *HttpClient) Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func (c *HttpClient) HttpGet(path string) ([]byte, error) {
	getUrl := fmt.Sprintf("%s/%s", c.Url, path)
	req, err := http.NewRequest(http.MethodGet, getUrl, nil)
	if err != nil {
		return nil, err
	}
	response, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	return body, err
}

func (c *HttpClient) HttpPost(path string, bytesData []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.Url, path)

	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}

	for headerKey, headerValue := range c.Option.Header {
		request.Header.Set(headerKey, headerValue)
	}

	resp, err := c.Client.Do(request)
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBytes, err
}
