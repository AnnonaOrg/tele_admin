package about_features

import (
	"github.com/umfaka/umfaka_core/internal/constvar"
	"github.com/umfaka/umfaka_core/internal/features"
	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/version", OnVersion)
}

func OnVersion(c tele.Context) error {
	text := constvar.APPAbout()
	return c.Reply(text)
}
