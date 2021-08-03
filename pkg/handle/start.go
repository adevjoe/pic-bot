package handle

import tb "gopkg.in/tucnak/telebot.v2"

func Start(bot *tb.Bot, message *tb.Message) {
	_, _ = bot.Send(message.Sender, "欢迎使用图片分享 Bot！您可以直接发送图片来投稿，也可以发送社交媒体的分享链接（目前支持 Twitter、Instagram、Weibo），如：https://twitter.com/alracoco/status/1421829535532687365", &tb.SendOptions{
		DisableWebPagePreview: true,
	})
}
