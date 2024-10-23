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
	if IsChatAdmin(c) {
		return nil
	}
	var userID, groupID int64
	if sender := c.Message().Sender; sender != nil {
		userID = sender.ID
	}
	groupID = c.Message().Chat.ID
	// botID = c.Bot().Me.ID
	if osenv.IsBotManagerID(userID) {
		return nil
	}

	if isForwarded := c.Message().IsForwarded(); isForwarded {
		return fmt.Errorf("Message Sender(%d,%d) IsForwarded", userID, groupID)
	}
	// if isBlocked := service.CheckBlockedUser(userID, botID); isBlocked {
	// 	return fmt.Errorf("Message Sender(%d,%d) is Blocked", userID, botID)
	// }
	if count, err := service.GetCountBlockedUserByUserIDAndGroupID(userID, groupID); count > 0 {
		return fmt.Errorf("Message Sender(%d,%d) is Blocked", userID, groupID)
	} else if err != nil {
		log.Errorf("GetCountBlockedUserByUserIDAndGroupID(%d,%d): %v", userID, groupID, err)
	}
	return nil
}
