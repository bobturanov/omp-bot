package verification

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"

)

func (c *ServiceVerificationCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	idx, err := strconv.Atoi(args)

	if err != nil {
		c.handleError("Please enter a number as an argument.",
			fmt.Sprintf("wrong args: %v", args),
			inputMessage)
		return
	}

	item, err := c.verificationService.Describe(uint64(idx))
	if err != nil {
		c.handleError(fmt.Sprintf("I can not find item with id: %d", idx),
			fmt.Sprintf("fail to get item with idx %d: %v", idx, err),
			inputMessage)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		item.Name,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Get: error sending reply message to chat - %v", err)
	}
}
