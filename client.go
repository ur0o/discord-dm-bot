package ddm

import (
	"encoding/json"
	"time"

	"github.com/tidwall/gjson"

	"github.com/ur0o/discord-dm-bot/request"
	"github.com/ur0o/discord-dm-bot/message"
)

type Client struct {
	Message *message.MessageService
}

func New(bt, ui string) (*Client, error) {
	cli := request.HTTPClient{BotToken: bt, UserId: ui}
	dmId, err := fetchDmId(cli)
	if err != nil {
		return nil, err
	}
	m := message.NewMessageService(&cli, dmId)
	return &Client{Message: m}, nil
}

func fetchDmId(cli request.HTTPClient) (string, error) {
	path := "/api/users/@me/channels"

	data := map[string]string{
		"recipient_id": cli.UserId,
	}
	d, _ := json.Marshal(data)
	res, err := cli.PostJson(path, d, map[string]string{})
	if err != nil {
		return "", err
	}
	return gjson.Get(string(res), "id").Str, nil
}

func MaxRetry(maxRetry uint) request.Option {
	return request.MaxRetry(maxRetry)
}

func RetryInterval(retryInterval time.Duration) request.Option {
	return request.RetryInterval(retryInterval)
}

func Timeout(timeout time.Duration) request.Option {
	return request.Timeout(timeout)
}