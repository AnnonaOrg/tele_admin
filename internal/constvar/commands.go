package constvar

import (
	tele "gopkg.in/telebot.v3"
)

var Commands []tele.Command

func init() {
	Commands = []tele.Command{
		{
			Text:        "/start",
			Description: "开始",
		},
		{
			Text:        "/ban",
			Description: "屏蔽用户",
		},
		{
			Text:        "/unban",
			Description: "解除屏蔽用户",
		},
		{
			Text:        "/version",
			Description: "查看版本",
		},
	}
}
