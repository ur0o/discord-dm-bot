package request

import (
	"fmt"
	"encoding/json"
)

func SendMessage(m string) error {
	if dmId == "" {
		return fmt.Errorf("DMID is empty; ensure Init() has been execute.")
	}

	path := "/api/channels/" + dmId + "/messages"
	data := map[string]string{
		"content": m,
	}
	buf, _ := json.Marshal(&data)
	_, err := postJson(host + path, buf, default_header())
	return err
}