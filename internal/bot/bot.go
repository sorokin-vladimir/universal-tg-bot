package bot

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api *tgbotapi.BotAPI
}

func New(token string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{api: api}, nil
}

func (b *Bot) SendMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = tgbotapi.ModeMarkdown
	_, err := b.api.Send(msg)
	return err
}

// Listen starts polling for incoming updates and handles commands.
// Blocks until the bot API returns a terminal error.
func (b *Bot) Listen() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		text := strings.TrimSpace(update.Message.Text)

		var replyText string
		switch text {
		case "/health":
			replyText = "✅ OK"
		default:
			replyText = ">>>"
		}

		reply := tgbotapi.NewMessage(update.Message.Chat.ID, replyText)
		reply.ReplyToMessageID = update.Message.MessageID
		if _, err := b.api.Send(reply); err != nil {
			log.Printf("command reply error: %v", err)
		}
	}
}
