package verification

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ServiceVerificationCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	itemID, err := strconv.Atoi(args)

	if err != nil {
		c.handleError("Please enter a number as an argument.",
			fmt.Sprintf("wrong args %v", args),
			inputMessage)
		return
	}

	status, err := c.verificationService.Remove(uint64(itemID))
	if err != nil {
		c.handleError(fmt.Sprintf("Deleting item ID: %d Error: %v", itemID, err),
			fmt.Sprintf("Can't delete item ID: %d, maybe it doesn't exist", itemID),
			inputMessage)
		return
	}

	msgValue := fmt.Sprintf("Item ID %d is deleted", itemID)
	if !status {
		msgValue = fmt.Sprintf("Can't delete item ID: %d, maybe it doesn't exist", itemID)
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		msgValue,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ServiceVerificationCommander.Remove: error sending reply message to chat - %v", err)
	}
}
