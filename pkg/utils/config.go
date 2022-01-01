package utils

import (
	"flag"
	"os"
	"time"
)

var (
	// BotToken define bot token
	BotToken string
	// TickTime define tick time
	TickTime = 10 * time.Second
)

func InitFlag() {
	// init args
	flag.StringVar(&BotToken, "botToken", os.Getenv("BOT_TOKEN"), "BotToken define bot token")
}
