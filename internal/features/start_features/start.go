package start_features

import (
	"github.com/umfaka/umfaka_core/internal/features"
	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/start", OnStart)
}
func OnStart(c tele.Context) error {
	return nil
}
