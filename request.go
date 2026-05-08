package ddm

import (
	"bytes"
	"fmt"
	"io"
	"time"
	"net/http"
)

const host = "https://discordapp.com"

func post(url string, request_body []byte, headers map[string]string, opts ...RequestOption) ([]byte, error) {
	options := requestOptions{
		MaxRetry: 3,
		RetryInterval: 500 * time.Millisecond,
		Timeout:	10 * time.Second,
	}
	for _, reqOpt := range(opts) { reqOpt(&options) }

	client := &http.Client{
		Timeout: options.Timeout,
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(request_body))
	for key, value := range headers {
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

func postJson(url string, request_body []byte, headers map[string]string, opts ...RequestOption) ([]byte, error) {
	headers["Content-Type"] = "application/json"
	return post(url, request_body, headers, opts...)
}

func defaultHeader(botToken string) map[string]string {
	return map[string]string{
		"authorization": "Bot " + botToken,
	}
}
