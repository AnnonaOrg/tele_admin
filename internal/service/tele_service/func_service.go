package tele_service

import (
	tele "gopkg.in/telebot.v3"
)

// 检查是否具有管理权限,出错返回false
func IsChatAdmin(c tele.Context, userID int64) bool {
	if userID > 0 {
		list, err := c.Bot().AdminsOf(c.Chat())
		if err != nil {
			return false
		}
		// log.Debugf("AdminsOf(%d): %+v", c.Chat().ID, list)
		for _, v := range list {
			if user := v.User; user != nil {
				// log.Debugf("AdminsOf(%d) user: %+v", c.Chat().ID, user)
				if user.ID == userID {
					return true
				}
			}

		}
	}
	return false
}
