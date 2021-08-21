package handle

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
)

func Start(bot *tb.Bot, message *tb.Message) {
	_, _ = bot.Send(message.Sender, os.Getenv("WELCOME_MESSAGE"), &tb.SendOptions{
		DisableWebPagePreview: true,
	})
}
