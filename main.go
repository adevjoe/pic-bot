package main

import (
	"flag"
	"fmt"
	"log"
	url2 "net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	// BotToken define bot token
	BotToken string
	// ChatID define chat id
	ChatID int64
	// Receiver define group or channel
	Receiver tb.ChatID
	// tick time
	tickTime time.Duration = 10 * time.Second
)

func init() {
	chatDefault, _ := strconv.Atoi(os.Getenv("CHAT_ID"))
	// init args
	flag.StringVar(&BotToken, "botToken", os.Getenv("BOT_TOKEN"), "BotToken define bot token")
	flag.Int64Var(&ChatID, "chatID", int64(chatDefault), "ChatID define chat id")

	Receiver = tb.ChatID(ChatID)
}

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  BotToken,
		Poller: &tb.LongPoller{Timeout: tickTime},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(tb.OnText, func(m *tb.Message) {
		log.Printf("receive text: %s", m.Text)
		// parse url
		url, err := url2.Parse(m.Text)
		if err != nil {
			errMsg := fmt.Sprintf("url %s is not valid", m.Text)
			log.Print(errMsg)
			_, _ = b.Send(m.Sender, errMsg)
			return
		}

		// exec gallery-dl
		cmd := exec.Command("gallery-dl", url.String())
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("cmd.Run() failed with %s\n, cmd: %s", err, cmd.String())
			return
		}
		log.Printf("out: %s", out)

		// get media
		s := strings.Split(strings.TrimRight(string(out), "\n"), "\n")
		if len(s) == 0 {
			log.Printf("media is empty")
			return
		}

		// generate album
		var album tb.Album
		for _, f := range s {
			f = strings.Replace(f, "# ", "", 1)
			if strings.Contains(f, ".mp4") {
				album = append(album, &tb.Video{
					File: tb.FromDisk(f),
				})
			} else {
				album = append(album, &tb.Photo{
					File: tb.FromDisk(f),
				})
			}
		}

		// send message
		_, err = b.SendAlbum(Receiver, album)
		if err != nil {
			log.Printf("sent to receiver %s error, %v", Receiver.Recipient(), err)
			return
		}
		log.Printf("%s successful", m.Text)
		_, _ = b.Send(m.Sender, "投稿成功！", &tb.SendOptions{
			ReplyTo:             m,
			DisableNotification: true,
		})
	})
	b.Handle(tb.OnPhoto, func(m *tb.Message) {
		log.Printf("got photo %s from %d(@%s)", m.Photo.UniqueID, m.Sender.ID, m.Sender.Username)
		_, err := b.Send(Receiver, m.Photo)
		if err != nil {
			log.Printf("sent photo to receiver %s error, %v", Receiver.Recipient(), err)
			return
		}
		_, _ = b.Send(m.Sender, "投稿成功！", &tb.SendOptions{
			DisableNotification: true,
		})
	})

	b.Start()
}
