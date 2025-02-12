package proc

import (
	"context"
	"github.com/meesooqa/bbboltg/app/filter"
	"time"

	"go.uber.org/zap"

	"github.com/meesooqa/bbboltg/app/config"
	"github.com/meesooqa/bbboltg/app/telegram"
)

// Processor is a feed reader and store writer
type Processor struct {
	Conf            *config.Conf
	TelegramService telegram.Service
	reader          Reader
}

func NewProcessor(conf *config.Conf) *Processor {
	//telegramService := telegram.NewTgService(conf);
	telegramService := telegram.NewEmptyService()
	reader := NewTelegramReader(telegramService)
	return &Processor{
		Conf:            conf,
		TelegramService: telegramService,
		reader:          reader,
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
		if err := p.processChannel(ctx, log, ch); err != nil {
			return err
		}
	}
	return nil
}

func (p *Processor) processChannel(ctx context.Context, log *zap.Logger, ch config.ConfChannel) error {
	channelFrom, err := p.TelegramService.GetChannel(ctx, ch.From)
	if err != nil {
		return err
	}

	messages, err := p.reader.Execute(ctx, channelFrom)
	if err != nil {
		return err
	}
	messages, err = p.filterMessages(messages, ch)
	if err != nil {
		return err
	}
	return nil
}

func (p *Processor) filterMessages(messages []telegram.TelegramMessage, ch config.ConfChannel) ([]telegram.TelegramMessage, error) {
	var result []telegram.TelegramMessage
	allFiltersAreEmpty := true

	for _, ccFilter := range ch.Filters {
		if ccFilter != nil && ccFilter.Prefix != "" {
			messagesFilter, err := filter.NewFilter(ccFilter)
			if err != nil {
				return nil, err
			}
			filteredMessages := filter.FilterTelegramMessages(messages, messagesFilter.IsValid)
			result = append(result, filteredMessages...)
			allFiltersAreEmpty = false
		}
	}

	if allFiltersAreEmpty {
		result = messages
	}

	return result, nil
}
