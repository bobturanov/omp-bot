package verification

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *ServiceVerificationCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	if len(args) <= 0 {
		c.handleError("Please enter a name as an argument.",
			fmt.Sprintf("empty args %v", args),
			inputMessage)
		return
	}

	item := c.verificationService.Create(args)
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Item ID: %d, Name: %v created.\nTry /list__service__verification last page", item.ID, item.Name),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("ServiceVerificationCommander.New: error sending reply message to chat - %v", err)
	}
	return
}
