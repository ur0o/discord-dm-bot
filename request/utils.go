package request

import (
	"fmt"
	"encoding/json"

	"github.com/tidwall/gjson"
)

var botToken, dmId string

func SetBotToken(token string) {
	botToken = token
}

func SetDmId(ui string) error {
	if botToken == "" {
		return fmt.Errorf("botToken is empty; please set valid botToken.")
	}

	path := "/api/users/@me/channels"

	data := map[string]string{
		"recipient_id": ui,
	}
	d, _ := json.Marshal(data)
	res, err := postJson(host + path, d, default_header())
	if err != nil {
		return err
	}
	dmId = gjson.Get(string(res), "id").Str
	return nil
}