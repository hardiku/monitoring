package handlers

import (
	"github.com/bartholdbos/golegram"
	"github.com/mdeheij/monitoring/configuration"
	"github.com/mdeheij/monitoring/log"
)

type instanceHolder struct {
	bot *golegram.Bot
}

var instance instanceHolder

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
	if instance.bot == nil {
		log.Info("Instance bot initialized! Using token:", configuration.Config.TelegramBotToken)
		instance.bot, err = golegram.NewBot(configuration.Config.TelegramBotToken)
	}

	if err == nil {

		if len(targets) >= 1 {
			//If targets have been set up in the specific service, use them
			for _, target := range targets {
				result, err := instance.bot.SendMessage(target, message)
				checkMessageError(result, err)
			}
		} else {
			//If not, then use default one
			target := configuration.Config.TelegramNotificationTarget
			result, err := instance.bot.SendMessage(target, message)
			checkMessageError(result, err)
		}

	} else {
		log.Error("Error sending Telegram message!", err.Error())
	}
}
