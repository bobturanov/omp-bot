package verification

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

func (c *ServiceVerificationCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	argsParts := strings.Split(args, " ")
	itemId, err := strconv.Atoi(argsParts[0])

	if err != nil {
		c.handleError("Please enter ID a number and Name as an argument.\nExample:\n\n/edit__service__verification 4 four edited.",
			fmt.Sprintf("wrong args %v", args),
			inputMessage)
		return
	}

	name := strings.Trim(args, argsParts[0]+" ")
	item, err := c.verificationService.Update(uint64(itemId), name)

	if err != nil {
		c.handleError("Please enter existing id a number as an argument.",
			fmt.Sprintf("fail to get item with idx %d: %v", itemId, err),
			inputMessage)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Item ID: %d Name: %v edited.", item.ID, item.Name),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Edit: error sending reply message to chat - %v", err)
	}
}