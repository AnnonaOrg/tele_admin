package tele_service

import (
	"github.com/umfaka/umfaka_core/internal/log"
	tele "gopkg.in/telebot.v3"
)

// 删除消息
func Delete(c tele.Context) {
	if err := c.Delete(); err != nil {
		log.Errorf("Delete(): %v", err)
	}
}
