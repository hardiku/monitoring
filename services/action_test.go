package services

import (
	"flag"
	"fmt"
	"testing"

	"github.com/mdeheij/monitoring/configuration"
)

var target string
var token string

func init() {
	flag.StringVar(&target, "target", "", "Telegram Target")
	flag.StringVar(&token, "token", "", "Telegram Bot Token")
	flag.Parse()

}

func TestAction(t *testing.T) {
	//fmt.Println(ServicesConfig)
	//fmt.Println("[Testing]")

	if target == "" || token == "" {
		fmt.Println("Token and/or target is not set! TEST WILL NOT RUN")
		t.SkipNow()

	} else {
		configuration.Config.TelegramBotToken = token
		configuration.Config.TelegramNotificationTarget = target
	}

	tgSlice := []string{target}
	ac := ActionConfig{Name: "telegram", Telegramtarget: tgSlice}
	s := Service{Host: "go test", Identifier: "TestAction", Threshold: 3, Health: 1, Output: "OK", Action: ac}

	a := NewAction(s)
	a.Run()

	s = Service{Host: "dev", Identifier: "localhost", Threshold: 3, Health: 1, Output: "OK", Action: ActionConfig{Name: "rpe"}}

	a = NewAction(s)
	a.Run()

	fmt.Println("A test message should be sent. Please verify.")
}
