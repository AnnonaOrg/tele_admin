package tele_service

import (
	"strings"
	"time"

	"github.com/AnnonaOrg/osenv"
	"github.com/umfaka/umfaka_core/internal/log"
	"github.com/umfaka/umfaka_core/internal/request"
	"github.com/umfaka/umfaka_core/internal/service"
	tele "gopkg.in/telebot.v3"
)

// 添加屏蔽名单
func AddBlockedUser(c tele.Context) {
	// 忽略非群消息
	if isFromGroup := c.Message().FromGroup(); !isFromGroup {
		return
	}
	var userID, botID, groupID int64
	if sender := c.Message().Sender; sender != nil {
		userID = sender.ID
	}
	botID = c.Bot().Me.ID
	groupID = c.Message().Chat.ID
	// 忽略非管理员消息
	if isBotManagerID := osenv.IsBotManagerID(userID); !isBotManagerID {
		if !IsChatAdmin(c) {
			return
		}
		return
	}

	// text := c.Message().Text
	// log.Debugf("Message().Text: %s", text)
	// if isFlag := strings.EqualFold(text, "标记"); !isFlag {
	// 	return
	// }

	var req request.BlockedUserRequest
	req.BotID = botID
	req.GroupID = groupID
	if c.Message().IsReply() {
		if replyTo := c.Message().ReplyTo; replyTo != nil {
			if sender := replyTo.Sender; sender != nil {
				req.UserID = sender.ID
				req.UserName = sender.Username
				req.FirstName = sender.FirstName
				req.LastName = sender.LastName
				// req.BotID = botID
				// req.GroupID = groupID
				if userID == sender.ID {
					return
				}
				// c.Bot().BanSenderChat(
				// 	c.Chat(),
				// 	sender,
				// )
				chatMember := &tele.ChatMember{
					User: &tele.User{
						ID: sender.ID,
					},
				}
				c.Bot().Restrict(
					c.Chat(),
					chatMember,
				)
			}

			go func() {
				time.Sleep(3 * time.Second)
				if err := c.Bot().Delete(replyTo); err != nil {
					log.Errorf("Delete(%+v): %v", replyTo, err)
				}
				Delete(c)
			}()
		}
	} else if payload := c.Message().Payload; len(payload) > 0 {
		if strings.HasPrefix(payload, "@") {
			if item, err := c.Bot().ChatByUsername(payload); err != nil {
				if err := c.Reply("未找到用户(%s)信息: %v", payload, err); err != nil {
					log.Errorf("Reply(未找到用户(%s)信息): %v", payload, err)
				}
				return
			} else {
				if userID == item.ID {
					return
				}
				req.UserID = item.ID
				req.UserName = item.Username
				req.FirstName = item.FirstName
				req.LastName = item.LastName
				// req.BotID = botID
				// req.GroupID = groupID
			}

		} else {
			c.Reply("指令格式: /ban @用户名")
			return
		}
	} else {
		c.Reply("指令格式: /ban @用户名")
		return
	}

	log.Debugf("req: %+v", req)
	if _, err := service.CreateBlockedUserEx(&req); err != nil {
		log.Errorf("CreateBlockedUserEx(%+v): %v", req, err)
		return
	}

	if msg, err := c.Bot().Send(&tele.Chat{ID: c.Chat().ID}, "🟢操作成功"); err != nil {
		log.Errorf("Send(成功标记): %v", err)
	} else {
		time.Sleep(10 * time.Second)
		if err := c.Bot().Delete(msg); err != nil {
			log.Errorf("Delete(%+v): %v", msg, err)
		}
	}
}
