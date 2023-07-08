package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) New(inputMessage *tgbotapi.Message) {
	arg := inputMessage.CommandArguments()

	products, err := c.productService.New(arg)
	if err != nil {
		log.Printf("%v", err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("%v", err))
		c.bot.Send(msg)
	} else {
		outputMsg := "Here all products: \n\n"

		for _, p := range products {
			outputMsg += p.Title
			outputMsg += "\n"
		}

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)

		c.bot.Send(msg)
	}
}
