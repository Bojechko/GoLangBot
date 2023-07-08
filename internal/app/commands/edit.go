package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

func (c *Commander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	words := strings.Split(args, " ")
	id, err := strconv.Atoi(words[0])
	if err != nil {
		log.Println("wrong args", args)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("id should be integer"))
		c.bot.Send(msg)
		return
	}

	title := strings.Join(words[1:], " ")

	product, err := c.productService.Edit(id, title)
	if err != nil {
		log.Printf("%v", err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("%v", err))
		c.bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, product.Title)
		c.bot.Send(msg)
	}
}
