package message

import (
	"encoding/json"

	"github.com/ur0o/discord-dm-bot/request"
)

type MessageService struct {
	client *request.HTTPClient
	dmId	 string
}

func NewMessageService(cli *request.HTTPClient, dmId string) *MessageService {
	return &MessageService{client: cli, dmId: dmId}
}

func (s *MessageService) Post(m string, options ...request.Option) error {
	path := "/api/channels/" + s.dmId + "/messages"
	data := map[string]string{
		"content": m,
	}
	buf, _ := json.Marshal(&data)
	_, err := s.client.PostJson(path, buf, map[string]string{}, options...)
	return err
}