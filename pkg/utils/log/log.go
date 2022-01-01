package log

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
)

func Info(sender *tb.User, format string, v ...interface{}) {
	format = "[ID: " + sender.FirstName + "] [Username: " + sender.LastName + "] " + format
	log.Printf(format, v...)
}
