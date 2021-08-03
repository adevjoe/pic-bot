package handle

import (
	"github.com/adevjoe/pic-bot/pkg/utils"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
)

func Photo(bot *tb.Bot, message *tb.Message) {
	log.Printf("got photo %s from %d(@%s)", message.Photo.UniqueID, message.Sender.ID, message.Sender.Username)
	_, err := bot.Send(utils.Receiver, message.Photo)
	if err != nil {
		log.Printf("sent photo to receiver %s error, %v", utils.Receiver.Recipient(), err)
		return
	}
	_, _ = bot.Send(message.Sender, "投稿成功！", &tb.SendOptions{
		DisableNotification: true,
	})
}
