package ddb

import (
	"github.com/ur0o/discord-dm-bot/request"
)

func Init(bt, ui string) error {
	request.SetBotToken(bt)
	return request.SetDmId(ui)
}

func SendMessage(m string) error {
	return request.SendMessage(m)
}