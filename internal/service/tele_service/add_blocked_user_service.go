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
	log.Debugf("userID: %d,groupID: %d", userID, groupID)
	// 忽略非管理员消息
	if !osenv.IsBotManagerID(userID) && !IsChatAdmin(c, userID) {
		log.Debugf("忽略非管理员(%d)指令", userID)
		return
	}

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
				if IsChatAdmin(c, req.UserID) {
					log.Debugf("IsChatAdmin(%d): Yes", req.UserID)
					c.Reply("神仙打架，凡人躲在一旁看热闹，结果还是被波及了，真是个‘不想当旁观者的’命运！")
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

				req.UserID = item.ID
				req.UserName = item.Username
				req.FirstName = item.FirstName
				req.LastName = item.LastName
				// req.BotID = botID
				// req.GroupID = groupID
				if userID == item.ID {
					return
				}
				if IsChatAdmin(c, req.UserID) {
					log.Debugf("IsChatAdmin(%d): Yes", req.UserID)
					c.Reply("神仙打架，凡人躲在一旁看热闹，结果还是被波及了，真是个‘不想当旁观者的’命运！")
					return
				}
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
