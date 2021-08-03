package utils

import (
	"flag"
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
	"strconv"
	"time"
)

var (
	// BotToken define bot token
	BotToken string
	// ChatID define chat id
	ChatID int64
	// Receiver define group or channel
	Receiver tb.ChatID
	// TickTime define tick time
	TickTime = 10 * time.Second
)

func InitFlag() {
	chatDefault, _ := strconv.Atoi(os.Getenv("CHAT_ID"))
	// init args
	flag.StringVar(&BotToken, "botToken", os.Getenv("BOT_TOKEN"), "BotToken define bot token")
	flag.Int64Var(&ChatID, "chatID", int64(chatDefault), "ChatID define chat id")

	Receiver = tb.ChatID(ChatID)
}
