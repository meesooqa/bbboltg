package telegram

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	td "github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
	"go.uber.org/zap"

	"github.com/meesooqa/bbboltg/app/config"
)

// github.com/gotd/td/tg.MessageClass
type TelegramMessage interface {
	// String implements fmt.Stringer
	String() string
}

type Service interface {
	//Init(ctx context.Context) error
	ProcessChannel(ctx context.Context, ch config.ConfChannel) error
	GetChannel(ctx context.Context, name string) (*tg.Channel, error)
	GetMessages(ctx context.Context, channel *tg.Channel) ([]TelegramMessage, error)
	ForwardMessage(ctx context.Context, msg TelegramMessage, channelFrom *tg.Channel, channelTo *tg.Channel) error
}

type TgService struct {
	conf *config.Conf
	log  *zap.Logger
	api  *tg.Client
}

func NewTgService(conf *config.Conf, log *zap.Logger) *TgService {
	return &TgService{
		conf: conf,
		log:  log,
	}
}

func (s *TgService) ProcessChannel(ctx context.Context, ch config.ConfChannel) error {
	s.log.Info("TgService.ProcessChannel")

	flow := auth.NewFlow(AuthFlow{}, auth.SendCodeOptions{})
	client, err := td.ClientFromEnvironment(td.Options{
		Logger: s.log,
	})
	if err != nil {
		return err
	}
	return client.Run(ctx, func(ctx context.Context) error {
		if err := client.Auth().IfNecessary(ctx, flow); err != nil {
			return err
		}
		s.api = tg.NewClient(client)

		channelFrom, err := s.GetChannel(ctx, ch.From)
		if err != nil {
			s.log.Error("channelFrom getting", zap.Error(err))
			return err
		}
		messages, err := s.GetMessages(ctx, channelFrom)
		if err != nil {
			s.log.Error("message reading", zap.Error(err))
			return err
		}
		if len(messages) == 0 {
			s.log.Debug("no messages")
			return nil
		}
		/*
			messages, err = s.filterMessages(messages, ch)
			if err != nil {
				s.log.Error("message filtering", zap.Error(err))
				return err
			}
			if len(messages) == 0 {
				s.log.Debug("no filtered messages")
				return nil
			}
		*/
		/*
			messages = s.sortMessages(messages)
		*/
		channelTo, err := s.GetChannel(ctx, ch.To)
		if err != nil {
			s.log.Error("channelTo getting", zap.Error(err))
			return err
		}
		for _, message := range messages {
			err := s.ForwardMessage(ctx, message, channelFrom, channelTo)
			if err != nil {
				s.log.Error("message writing", zap.Error(err))
				return err
			}
		}

		return nil
	})
}

func (s *TgService) Init(ctx context.Context) error {
	flow := auth.NewFlow(AuthFlow{}, auth.SendCodeOptions{})
	client, err := td.ClientFromEnvironment(td.Options{
		Logger: s.log,
	})
	if err != nil {
		return err
	}
	return client.Run(ctx, func(ctx context.Context) error {
		if err := client.Auth().IfNecessary(ctx, flow); err != nil {
			return err
		}
		s.api = tg.NewClient(client)
		return nil
	})
}

func (s *TgService) getRandomID() int64 {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rnd.Int63()
}

func (s *TgService) ForwardMessage(
	ctx context.Context,
	msg TelegramMessage,
	channelFrom *tg.Channel,
	channelTo *tg.Channel,
) error {
	m, ok := msg.(*tg.Message)
	if !ok {
		return nil
	}
	_, err := s.api.MessagesForwardMessages(ctx, &tg.MessagesForwardMessagesRequest{
		ID: []int{m.ID},
		FromPeer: &tg.InputPeerChannel{
			ChannelID:  channelFrom.ID,
			AccessHash: channelFrom.AccessHash,
		},
		ToPeer: &tg.InputPeerChannel{
			ChannelID:  channelTo.ID,
			AccessHash: channelTo.AccessHash,
		},
		RandomID:   []int64{s.getRandomID()},
		DropAuthor: false,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *TgService) GetChannel(ctx context.Context, name string) (*tg.Channel, error) {
	resolved, err := s.api.ContactsResolveUsername(ctx, &tg.ContactsResolveUsernameRequest{
		Username: name,
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка поиска канала %s: %w", name, err)
	}
	channel, ok := resolved.Chats[0].(*tg.Channel)
	if !ok {
		return nil, fmt.Errorf("канал не найден")
	}
	return channel, nil
}

func (s *TgService) GetMessages(ctx context.Context, channel *tg.Channel) ([]TelegramMessage, error) {
	limit := 2
	//OffsetID: lastMessageID,

	messages, err := s.api.MessagesGetHistory(ctx, &tg.MessagesGetHistoryRequest{
		Peer: &tg.InputPeerChannel{
			ChannelID:  channel.ID,
			AccessHash: channel.AccessHash,
		},
		Limit: limit,
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка получения сообщений: %w", err)
	}

	channelMessages := messages.(*tg.MessagesChannelMessages)
	if len(channelMessages.Messages) == 0 {
		s.log.Debug("новых сообщений нет")
		return nil, nil
	}

	result := make([]TelegramMessage, len(channelMessages.Messages))
	for i, msg := range channelMessages.Messages {
		result[i] = msg
	}
	return result, nil
}

type EmptyService struct{}

func NewEmptyService() *EmptyService {
	return &EmptyService{}
}

func (s *EmptyService) Init(ctx context.Context) error {
	return nil
}

func (s *EmptyService) ProcessChannel(ctx context.Context, ch config.ConfChannel) error {
	return nil
}

func (s *EmptyService) GetChannel(ctx context.Context, name string) (*tg.Channel, error) {
	return nil, nil
}

func (s *EmptyService) GetMessages(ctx context.Context, channel *tg.Channel) ([]TelegramMessage, error) {
	return nil, nil
}

func (s *EmptyService) ForwardMessage(ctx context.Context, msg TelegramMessage, channelFrom *tg.Channel, channelTo *tg.Channel) error {
	return nil
}
