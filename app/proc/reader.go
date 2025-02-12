package proc

import (
	"context"

	"github.com/gotd/td/tg"

	"github.com/meesooqa/bbboltg/app/telegram"
)

type Reader interface {
	Execute(ctx context.Context, channel *tg.Channel) ([]telegram.TelegramMessage, error)
}

type TelegramReader struct {
	TelegramService telegram.Service
}

func NewTelegramReader(telegramService telegram.Service) TelegramReader {
	return TelegramReader{
		TelegramService: telegramService,
	}
}

func (r TelegramReader) Execute(ctx context.Context, channel *tg.Channel) ([]telegram.TelegramMessage, error) {
	return r.TelegramService.GetMessages(ctx, channel)
}
