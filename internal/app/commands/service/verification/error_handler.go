package verification

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *ServiceVerificationCommander) handleError(userMsg string, systemMsg string, inputMessage *tgbotapi.Message){
	log.Print(systemMsg)
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		userMsg,
	)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("handleError: error sending reply message: %v [%v] to chat - %v", userMsg, systemMsg, err)
	}
}
