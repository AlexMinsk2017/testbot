package commands

import (
	"github.com/AlexMinsk2017/testbot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Comander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewComander(
	bot *tgbotapi.BotAPI,
	productService *product.Service) *Comander {
	return &Comander{
		bot:            bot,
		productService: productService,
	}
}
