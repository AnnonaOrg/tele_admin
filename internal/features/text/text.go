package text

import (
	"github.com/umfaka/umfaka_core/internal/features"
	"github.com/umfaka/umfaka_core/internal/service/tele_service"
	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature(tele.OnText, OnMessage)
	features.RegisterFeature(tele.OnPhoto, OnMessage)
	features.RegisterFeature(tele.OnAudio, OnMessage)
	features.RegisterFeature(tele.OnAnimation, OnMessage)
	features.RegisterFeature(tele.OnDocument, OnMessage)
	features.RegisterFeature(tele.OnSticker, OnMessage)
	features.RegisterFeature(tele.OnVideo, OnMessage)
	features.RegisterFeature(tele.OnVoice, OnMessage)
	features.RegisterFeature(tele.OnVideoNote, OnMessage)
	features.RegisterFeature(tele.OnContact, OnMessage)
	features.RegisterFeature(tele.OnLocation, OnMessage)
	features.RegisterFeature(tele.OnMedia, OnMessage)
}

// func OnText(c tele.Context) error {
// 	if err := tele_service.CheckBlockedUser(c); err != nil {
// 		tele_service.Delete(c)
// 		return nil
// 	}

// 	tele_service.AddBlockedUser(c)
// 	return nil
// }

func OnMessage(c tele.Context) error {
	if err := tele_service.CheckBlockedUser(c); err != nil {
		tele_service.Delete(c)
		return nil
	}

	return nil
}
