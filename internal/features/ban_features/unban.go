package ban_features

import (
	"github.com/umfaka/umfaka_core/internal/features"
	"github.com/umfaka/umfaka_core/internal/service/tele_service"
	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/ban", OnBan)
	features.RegisterFeature("/unban", OnUnban)
}

func OnBan(c tele.Context) error {
	tele_service.AddBlockedUser(c)
	return nil
}

func OnUnban(c tele.Context) error {
	tele_service.DeleteBlockedUser(c)
	return nil
}
