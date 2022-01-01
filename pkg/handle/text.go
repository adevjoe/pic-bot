package handle

import (
	"fmt"
	"github.com/adevjoe/pic-bot/pkg/utils"
	"github.com/adevjoe/pic-bot/pkg/utils/log"
	tb "gopkg.in/tucnak/telebot.v2"
	url2 "net/url"
	"os/exec"
	"strings"
)

func Text(bot *tb.Bot, message *tb.Message) {

	log.Info(message.Sender, "receive text: %s", message.Text)
	// parse url
	url, err := url2.Parse(message.Text)
	if err != nil {
		errMsg := fmt.Sprintf("url %s is not valid", message.Text)
		log.Info(message.Sender, errMsg)
		_, err = bot.Send(message.Sender, errMsg)
		if err != nil {
			log.Info(message.Sender, "send error: %s", err)
		}
		return
	}
	url.RawQuery = ""

	web := ""
	id := ""
	// analyse source
	switch true {
	case strings.Contains(url.Host, "twitter"):
		web = "Twitter"
		break
	case strings.Contains(url.Host, "weibo"):
		web = "Weibo"
		break
	case strings.Contains(url.Host, "instagram"):
		web = "Instagram"
		break
	default:
		errMsg := fmt.Sprintf("url %s is not supported", message.Text)
		log.Info(message.Sender, errMsg)
		_, err = bot.Send(message.Sender, errMsg)
		if err != nil {
			log.Info(message.Sender, "send error: %s", err)
		}
		return
	}

	// exec gallery-dl
	cmd := exec.Command("gallery-dl", url.String())
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Info(message.Sender, "cmd.Run() failed with %s\n, cmd: %s", err, cmd.String())
		return
	}
	log.Info(message.Sender, "out: %s", out)

	// get media
	s := strings.Split(strings.TrimRight(string(out), "\n"), "\n")
	if len(s) == 0 {
		log.Info(message.Sender, "media is empty")
		return
	}

	// generate album
	var albums tb.Album
	for i, f := range s {
		f = strings.Replace(f, "# ", "", 1)
		if id == "" {
			filepathList := strings.Split(f, "/")
			if len(filepathList) > 2 {
				id = filepathList[len(filepathList)-2]
			}
		}
		if strings.Contains(f, ".mp4") {
			album := &tb.Video{
				File: tb.FromDisk(f),
			}
			if i == 0 {
				album.Caption = utils.GetMediaInfoMessage(url.String(), web, id)
			}
			albums = append(albums, album)
		} else {
			album := &tb.Photo{
				File: tb.FromDisk(f),
			}
			albums = append(albums, album)
			if i == 0 {
				album.Caption = utils.GetMediaInfoMessage(url.String(), web, id)
			}
		}
	}

	// send message
	_, err = bot.SendAlbum(message.Sender, albums, &tb.SendOptions{ReplyTo: message})
	if err != nil {
		log.Info(message.Sender, "sent error, %v", err)
		return
	}
	log.Info(message.Sender, "%s successful", message.Text)
}
