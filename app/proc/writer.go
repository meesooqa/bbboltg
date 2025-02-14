package proc

import (
	"context"

	"github.com/gotd/td/tg"

	"github.com/meesooqa/bbboltg/app/telegram"
)

type Writer interface {
	Execute(
		ctx context.Context,
		messages []telegram.TelegramMessage,
		src *tg.Channel,
		dest *tg.Channel,
	) error
}

type TelegramWriter struct {
	TelegramService telegram.Service
}

func NewTelegramWriter(telegramService telegram.Service) *TelegramWriter {
	return &TelegramWriter{
		TelegramService: telegramService,
	}
}

func (r TelegramWriter) Execute(ctx context.Context, messages []telegram.TelegramMessage, src *tg.Channel, dest *tg.Channel) error {
	for _, message := range messages {
		err := r.TelegramService.ForwardMessage(ctx, message, src, dest)
		if err != nil {
			return err
		}
	}
	return nil
}
