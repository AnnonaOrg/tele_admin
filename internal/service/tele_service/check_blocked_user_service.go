package tele_service

import (
	"fmt"

	"github.com/AnnonaOrg/osenv"
	"github.com/umfaka/umfaka_core/internal/log"
	"github.com/umfaka/umfaka_core/internal/service"
	tele "gopkg.in/telebot.v3"
)

// 检查屏蔽名单 如果在黑名单或消息符合屏蔽条件，就返回错误，否则返回 nil
func CheckBlockedUser(c tele.Context) error {
	if isFromGroup := c.Message().FromGroup(); !isFromGroup {
		return nil
	}
	var userID, botID int64
	if sender := c.Message().Sender; sender != nil {
		userID = sender.ID
	}
	botID = c.Bot().Me.ID
	if osenv.IsBotManagerID(userID) {
		return nil
	}

	if isForwarded := c.Message().IsForwarded(); isForwarded {
		return fmt.Errorf("Message Sender(%d,%d) IsForwarded", userID, botID)
	}
	// if isBlocked := service.CheckBlockedUser(userID, botID); isBlocked {
	// 	return fmt.Errorf("Message Sender(%d,%d) is Blocked", userID, botID)
	// }
	if count, err := service.GetCountBlockedUserByUserIDAndBotID(userID, botID); count > 0 {
		return fmt.Errorf("Message Sender(%d,%d) is Blocked", userID, botID)
	} else if err != nil {
		log.Errorf("GetCountBlockedUserByUserIDAndBotID(%d,%d): %v", userID, botID, err)
	}
	return nil
}
