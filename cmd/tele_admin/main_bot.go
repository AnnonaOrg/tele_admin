package main

import (
	"time"

	"github.com/AnnonaOrg/osenv"
	_ "github.com/umfaka/umfaka_core/cmd/tele_admin/distro/all"
	"github.com/umfaka/umfaka_core/common"
	"github.com/umfaka/umfaka_core/internal/constvar"
	_ "github.com/umfaka/umfaka_core/internal/dotenv"
	"github.com/umfaka/umfaka_core/internal/features"
	"github.com/umfaka/umfaka_core/internal/log"
	tele "gopkg.in/telebot.v3"
)

func mainBot() {
	botToken := osenv.GetBotTelegramToken()
	botAPIProxyURL := osenv.GetBotTelegramAPIProxyURL()
	log.Debugf("GetBotTelegramAPIProxyURL(): %s", botAPIProxyURL)
	bot, err := tele.NewBot(tele.Settings{
		URL:    botAPIProxyURL,
		Token:  botToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Errorf("mainBot:NewBot(%s): %v", botToken, err)
	}
	common.Must(err)
	log.Infof("Bot: @%s %d", bot.Me.Username, bot.Me.ID)

	features.Handle(bot)

	commands := constvar.Commands
	bot.SetCommands(commands)

	bot.Start()
}
