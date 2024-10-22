package tele_service

import (
	"strconv"
	"strings"

	"github.com/AnnonaOrg/osenv"
	"github.com/umfaka/umfaka_core/internal/log"
	"github.com/umfaka/umfaka_core/internal/service"
	tele "gopkg.in/telebot.v3"
)

// 删除屏蔽名单
func DeleteBlockedUser(c tele.Context) {
	if isFromGroup := c.Message().FromGroup(); !isFromGroup {
		return
	}
	var userID, botID int64
	if sender := c.Message().Sender; sender != nil {
		userID = sender.ID
	}
	botID = c.Bot().Me.ID
	if !osenv.IsBotManagerID(userID) {
		// if !IsChatAdmin(c) {
		// 	return
		// }
		return
	}
	var targetUser int64
	// payload := c.Message().Payload
	if c.Message().IsReply() {
		if replyTo := c.Message().ReplyTo; replyTo != nil {
			if sender := replyTo.Sender; sender != nil {
				targetUser = sender.ID
			}
		}
	} else if payload := c.Message().Payload; len(payload) > 0 {
		if strings.HasPrefix(payload, "@") {
			username := strings.TrimPrefix(payload, "@")
			item, err := service.GetBlockedUserByUsername(username)
			if err != nil {
				log.Errorf("GetBlockedUserByUsername(%s): %v", username, err)
				return
			}
			targetUser = item.UserID
		} else if id, err := strconv.ParseInt(payload, 10, 64); err != nil {
			log.Errorf("ParseInt(%s): %v", payload, err)
			c.Reply("指令格式: /unban @用户名")
			return
		} else if id > 0 {
			targetUser = id
		}
	} else {
		c.Reply("指令格式: /unban @用户名")
		return
	}
	if userID == targetUser {
		return
	}

	c.Bot().Unban(c.Chat(), &tele.User{ID: targetUser}, true)
	chatMember := &tele.ChatMember{
		User: &tele.User{
			ID: targetUser,
		},
	}
	chatMember.CanSendMessages = true
	c.Bot().Restrict(
		c.Chat(),
		chatMember,
	)

	if err := service.DeleteBlockedUser(targetUser, botID); err != nil {
		log.Errorf("DeleteBlockedUser(%d,%d): %v", targetUser, botID, err)
		return
	}
	if err := c.Reply("🟢操作成功"); err != nil {
		log.Errorf("Reply(🟢操作成功！): %v", err)
	}

	return
}
