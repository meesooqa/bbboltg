package proc

import (
	"context"
	"sort"
	"time"

	"github.com/gotd/td/tg"
	"go.uber.org/zap"

	"github.com/meesooqa/bbboltg/app/config"
	"github.com/meesooqa/bbboltg/app/filter"
	"github.com/meesooqa/bbboltg/app/telegram"
)

// Processor is a feed reader and store writer
type Processor struct {
	Conf            *config.Conf
	log             *zap.Logger
	telegramService telegram.Service
	reader          Reader
	writer          Writer
}

func NewProcessor(conf *config.Conf, log *zap.Logger) *Processor {
	tgs := telegram.NewTgService(conf, log)
	//tgs := telegram.NewEmptyService()
	telegramService := telegram.NewCachingChannelInfoProvider(tgs, time.Hour)
	reader := NewTelegramReader(telegramService)
	writer := NewTelegramWriter(telegramService)
	return &Processor{
		Conf:            conf,
		log:             log,
		telegramService: telegramService,
		reader:          reader,
		writer:          writer,
	}
}

func (p *Processor) Do(ctx context.Context) error {
	ticker := time.NewTicker(p.Conf.UpdateInterval)
	defer ticker.Stop()

	if err := p.processChannels(ctx); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if err := p.processChannels(ctx); err != nil {
				return err
			}
		}
	}
}

func (p *Processor) processChannels(ctx context.Context) error {
	//err := p.telegramService.Init(ctx)
	//if err != nil {
	//	return err
	//}

	p.log.Info("processing channels")
	for _, ch := range p.Conf.Channels {
		//if err := p.processChannel(ctx, ch); err != nil {
		if err := p.telegramService.ProcessChannel(ctx, ch); err != nil {
			return err
		}
	}
	return nil
}

func (p *Processor) processChannel(ctx context.Context, ch config.ConfChannel) error {
	channelFrom, err := p.telegramService.GetChannel(ctx, ch.From)
	if err != nil {
		p.log.Error("channelFrom getting", zap.Error(err))
		return err
	}
	messages, err := p.reader.Execute(ctx, channelFrom)
	if err != nil {
		p.log.Error("message reading", zap.Error(err))
		return err
	}
	if len(messages) == 0 {
		p.log.Debug("no messages")
		return nil
	}
	messages, err = p.filterMessages(messages, ch)
	if err != nil {
		p.log.Error("message filtering", zap.Error(err))
		return err
	}
	if len(messages) == 0 {
		p.log.Debug("no filtered messages")
		return nil
	}
	messages = p.sortMessages(messages)
	channelTo, err := p.telegramService.GetChannel(ctx, ch.To)
	if err != nil {
		p.log.Error("channelTo getting", zap.Error(err))
		return err
	}
	err = p.writer.Execute(ctx, messages, channelFrom, channelTo)
	if err != nil {
		p.log.Error("message writing", zap.Error(err))
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

func (p *Processor) sortMessages(messages []telegram.TelegramMessage) []telegram.TelegramMessage {
	sort.Slice(messages, func(i, j int) bool {
		msgI, okI := messages[i].(*tg.Message)
		msgJ, okJ := messages[j].(*tg.Message)
		if !okI || !okJ {
			return false
		}
		return msgI.ID < msgJ.ID
	})
	return messages
}
