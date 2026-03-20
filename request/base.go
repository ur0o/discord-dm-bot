package request

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const host = "https://discordapp.com"

func default_header() map[string]string {
	return map[string]string{
		"authorization": "Bot " + botToken,
	}
}

func post(url string, request_body []byte, headers map[string]string) ([]byte, error) {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(request_body))
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	byteArray, err := io.ReadAll(res.Body)
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf(string(byteArray))
	}
	return byteArray, nil
}

func postJson(url string, request_body []byte, headers map[string]string) ([]byte, error) {
	headers["Content-Type"] = "application/json"
	return post(url, request_body, headers)
}