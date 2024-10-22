package callback

import (
	"github.com/umfaka/umfaka_core/internal/features"
	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature(tele.OnCallback, OnCallback)
}

// 点击按钮回掉
func OnCallback(c tele.Context) error {
	return nil
}
