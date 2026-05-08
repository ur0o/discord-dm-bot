package ddm

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

type Client struct {
	BotToken string
	UserId	 string
	DmId		 string
}

func NewClient(bt, ui string) (*Client, error) {
	path := "/api/users/@me/channels"

	data := map[string]string{
		"recipient_id": ui,
	}
	d, _ := json.Marshal(data)
	res, err := postJson(host + path, d, defaultHeader(bt))
	if err != nil {
		return nil, err
	}
	dmId := gjson.Get(string(res), "id").Str
	return &Client{BotToken: bt, UserId: ui, DmId: dmId}, nil
}

func (c *Client)SendMessage(m string, options ...RequestOption) error {
	path := "/api/channels/" + c.DmId + "/messages"
	data := map[string]string{
		"content": m,
	}
	buf, _ := json.Marshal(&data)
	_, err := postJson(host + path, buf, c.defaultHeader(), options...)
	return err
}

func (c *Client)defaultHeader() map[string]string {
	return defaultHeader(c.BotToken)
}