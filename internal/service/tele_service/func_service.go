package tele_service

import (
	tele "gopkg.in/telebot.v3"
)

// 检查是否具有管理权限,出错返回false
func IsChatAdmin(c tele.Context) bool {
	var userID int64
	if sender := c.Message().Sender; sender != nil {
		userID = sender.ID
	}
	if userID > 0 {
		list, err := c.Bot().AdminsOf(c.Chat())
		if err != nil {
			return false
		}
		for _, v := range list {
			if v.User.ID == userID {
				return true
			}
		}
	}
	return false
}
