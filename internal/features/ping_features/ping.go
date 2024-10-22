package ping_features

import (
	"github.com/umfaka/umfaka_core/internal/features"
	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/ping", OnPing)
}

func OnPing(c tele.Context) error {
	return c.Reply("Pong!")
}
