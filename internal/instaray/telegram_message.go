package instaray

import (
	"github.com/go-telegram/bot/models"
)

// TelegramUpdate is an alias for models.Update.
type TelegramUpdate = models.Update

// TelegramMessage represents a message received from the Telegram bot API.
type TelegramMessage models.Message

// Attrs returns a slice of logging attributes for the message.
func (tm TelegramMessage) Attrs() []any {
	attrs := []any{
		"scope", "telegram",
		"chat_id", tm.Chat.ID,
		"user_id", tm.From.ID,
		"username", tm.From.Username,
		"message_id", tm.ID,
	}

	if tm.IsTopicMessage {
		attrs = append(attrs, "message_thread_id", tm.MessageThreadID)
	}

	if tm.Text != "" {
		attrs = append(attrs, "message_text", tm.Text)
	}

	return attrs
}

// AttrsWithError returns a slice of logging attributes for the message with an
// error.
func (tm TelegramMessage) AttrsWithError(err error) []any {
	return append(tm.Attrs(), "error", err)
}
