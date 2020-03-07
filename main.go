package main

import (
	"log"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api"
	flag "github.com/spf13/pflag"
)

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func main() {

	bot, err := botapi.NewBotAPI("YourAPIToken")
	check(err)

	var message = flag.String("send-message", "check", "Send message to default Chat ID")
	flag.Parse()

	// Change your target account aka Telegram Chat ID
	var chatID int64 = 154321022

	log.Printf("Authorized as %s", bot.Self.UserName)

	msg := botapi.NewMessage(chatID, *message)

	bot.Send(msg)

}
