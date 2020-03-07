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

	// Change your target account aka Telegram Chat ID
	var chatID int64 = 154321022

	bot, err := botapi.NewBotAPI("YourAPIToken")
	check(err)

	var message = flag.String("send-message", "check", "Send message to default Chat ID")
	var id = flag.Int64("chat-id", chatID, "Specify the Telegram chat ID to send the message to")
	flag.Parse()

	log.Printf("Authorized as %s", bot.Self.UserName)

	msg := botapi.NewMessage(*id, *message)

	bot.Send(msg)

}
