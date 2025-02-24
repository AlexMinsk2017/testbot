package main

import (
	"github.com/AlexMinsk2017/testbot/internal/app/commands"
	"github.com/AlexMinsk2017/testbot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// const token = "7655447204:AAEfYMaYTT3-Uh-1yOhpoy7R1kmUaoeRvmU"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()
	commander := commands.NewComander(bot, productService)

	for update := range updates {
		commander.HandleUpdate(update)
	}
}
