package commands

import (
	"github.com/AlexMinsk2017/testbot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var registeredCommands = map[string]func(c *Commander, msg *tgbotapi.MessageConfig){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewComander(
	bot *tgbotapi.BotAPI,
	productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {

	if update.Message == nil {
		return
	}
	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	default:
		c.Default(update.Message)
	}

}
