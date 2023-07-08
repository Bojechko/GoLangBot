package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list entities \n"+
			"/get - get element by id\n"+
			"/edit - change element title by id\n"+
			"/new - add new element\n"+
			"/help - help")
	c.bot.Send(msg)
}
