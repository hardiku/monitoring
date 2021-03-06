package handlers

import (
	"github.com/bartholdbos/golegram"
	"github.com/mdeheij/monitoring/configuration"
	"github.com/mdeheij/monitoring/log"
)

var bot *golegram.Bot

func SetupTelegram() {
	var err error
	log.Error("Instance bot initialized! Using token:", configuration.Config.TelegramBotToken)
	bot, err = golegram.NewBot(configuration.Config.TelegramBotToken)
	if err != nil {
		log.Error("[Telegram] Init error")
		log.Error(err)
	}
}

func checkMessageError(result golegram.Message, err error) {
	if err != nil {
		log.Error("[Telegram] Message error")
		log.Error(err)
		log.Error(result)
	}
}

//Telegram sends a Telegram message to one or more users by their unique ID
func Telegram(targets []string, message string) {
	var err error

	//Check if a bot instance have been constructed before
	if bot == nil {
		SetupTelegram()
	}

	if err == nil {

		if len(targets) >= 1 {
			//If targets have been set up in the specific service, use them
			for _, target := range targets {
				//disable_web_page_preview bool, parse_mode string
				result, err := bot.SendMessage(target, message, true, "Markdown")
				checkMessageError(result, err)
			}
		} else {
			//If not, then use default one
			target := configuration.Config.TelegramNotificationTarget
			result, err := bot.SendMessage(target, message, true, "Markdown")
			checkMessageError(result, err)
		}

	} else {
		log.Error("Error sending Telegram message!", err.Error())
	}
}
