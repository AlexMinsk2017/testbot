package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()
	arg, err := strconv.Atoi(args)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	//msg := tgbotapi.NewMessage(inputMessage.Chat.ID, inputMessage.Text)
	//msg.ReplyToMessageID = inputMessage.MessageID

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Hello, succesfully, %v", arg))

	c.bot.Send(msg)

}
