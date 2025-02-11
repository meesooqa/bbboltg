package telegram

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/meesooqa/bbboltg/app/config"
)

// Processor is a feed reader and store writer
type Processor struct {
	Conf            *config.Conf
	TelegramService Service
}

func NewProcessor(conf *config.Conf) *Processor {
	return &Processor{
		Conf: conf,
		//TelegramService: NewTgService(conf),
		TelegramService: NewEmptyService(),
	}
}

func (p *Processor) Do(ctx context.Context, log *zap.Logger) error {
	ticker := time.NewTicker(p.Conf.UpdateInterval)
	defer ticker.Stop()

	if err := p.processChannels(ctx, log); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if err := p.processChannels(ctx, log); err != nil {
				return err
			}
		}
	}
}

func (p *Processor) processChannels(ctx context.Context, log *zap.Logger) error {
	for _, ch := range p.Conf.Channels {
		if err := p.TelegramService.ProcessChannel(ctx, log, ch); err != nil {
			return err
		}
	}
	return nil
}
