package main

import (
	"io/ioutil"
	"log"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api"
	homedir "github.com/mitchellh/go-homedir"
	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

type Conf struct {
	Config struct {
		APIToken string `yaml:"api-token"`
	} `yaml:"config"`
}

var filePath string

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func main() {

	// Change your target account aka Telegram Chat ID
	var chatID int64 = 154321022

	var home, err = homedir.Dir()
	check(err)

	// Sets default path for go-telegram config file
	var defaultPath = home + "/.go-telegram/config"

	var config = flag.String("config", defaultPath, "Specify the path to config file")
	var id = flag.Int64("chat-id", chatID, "Specify the Telegram chat ID to send the message to")
	flag.Parse()

	var newMsg = flag.Args()[0]

	var conf = getConf(*config)
	bot, err := botapi.NewBotAPI(conf.Config.APIToken)
	check(err)
	log.Printf("Authorized as %s", bot.Self.UserName)

	msg := botapi.NewMessage(*id, newMsg)

	bot.Send(msg)

}

func getConf(configPath string) *Conf {

	c := &Conf{}
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = yaml.Unmarshal(yamlFile, c)

	return c

}
