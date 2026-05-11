package request

import (
	"bytes"
	"fmt"
	"io"
	"maps"
	"net/http"
	"time"
)

type HTTPClient struct {
	BotToken string
	UserId	 string
}

const host = "https://discordapp.com"

func (c *HTTPClient) Post(path string, request_body []byte, headers map[string]string, opts ...Option) ([]byte, error) {
	options := options{
		MaxRetry: 3,
		RetryInterval: 500 * time.Millisecond,
		Timeout:	10 * time.Second,
	}
	for _, reqOpt := range(opts) { reqOpt(&options) }

	client := &http.Client{
		Timeout: options.Timeout,
	}
	newHeaders := map[string]string{}
	req, _ := http.NewRequest("POST", host + path, bytes.NewBuffer(request_body))
	maps.Copy(newHeaders, headers)
	maps.Copy(newHeaders, c.authHeader())
	for key, value := range newHeaders {
		req.Header.Set(key, value)
	}

	var res *http.Response
	var err error
	for i := 0; i < options.MaxRetry; i++ {
		res, err = client.Do(req)
		if err != nil { return nil, err }
		defer res.Body.Close()

		if res.StatusCode < http.StatusInternalServerError {
			break
		}
		time.Sleep(options.RetryInterval)
	}

	var byteArray []byte
	byteArray, err = io.ReadAll(res.Body)
	if err != nil { return nil, err }
	if res.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("%s", string(byteArray))
	}

	return byteArray, nil
}

func (c *HTTPClient) PostJson(url string, request_body []byte, headers map[string]string, opts ...Option) ([]byte, error) {
	headers["Content-Type"] = "application/json"
	return c.Post(url, request_body, headers, opts...)
}

func (c *HTTPClient) authHeader() map[string]string {
	return map[string]string{
		"authorization": "Bot " + c.BotToken,
	}
}
