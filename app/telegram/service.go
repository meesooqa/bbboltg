package telegram

import (
	"context"
	"fmt"
	"math/rand"
	"sort"
	"time"

	td "github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
	"go.uber.org/zap"

	"github.com/meesooqa/bbboltg/app/config"
)

type Service interface {
	ProcessChannel(ctx context.Context, log *zap.Logger, ch config.ConfChannel) error
}

type TgService struct {
	conf *config.Conf
}

func NewTgService(conf *config.Conf) *TgService {
	return &TgService{conf: conf}
}

func (s *TgService) ProcessChannel(ctx context.Context, log *zap.Logger, ch config.ConfChannel) error {
	flow := auth.NewFlow(AuthFlow{}, auth.SendCodeOptions{})

	client, err := td.ClientFromEnvironment(td.Options{
		Logger: log,
	})
	if err != nil {
		return err
	}

	return client.Run(ctx, func(ctx context.Context) error {
		if err := client.Auth().IfNecessary(ctx, flow); err != nil {
			return err
		}

		api := tg.NewClient(client)
		// collect data
		channel, err := s.getChannel(ch.From, ctx, api)
		if err != nil {
			return err
		}
		messages, err := s.getMessages(channel, ctx, api)
		if err != nil {
			return err
		}
		toChannel, err := s.getChannel(ch.To, ctx, api)
		if err != nil {
			return err
		}
		randomIDs := s.generateRandomIDs(len(messages))
		// send
		for i, msg := range messages {
			m, ok := msg.(*tg.Message)
			if !ok {
				continue
			}

			_, err := api.MessagesForwardMessages(ctx, &tg.MessagesForwardMessagesRequest{
				ID: []int{m.ID},
				FromPeer: &tg.InputPeerChannel{
					ChannelID:  channel.ID,
					AccessHash: channel.AccessHash,
				},
				ToPeer: &tg.InputPeerChannel{
					ChannelID:  toChannel.ID,
					AccessHash: toChannel.AccessHash,
				},
				RandomID:   []int64{randomIDs[i]},
				DropAuthor: false,
			})

			if err != nil {
				return err
			} else {
				fmt.Printf("✅ %d\n", m.ID)
			}
		}

		return nil
	})
}

func (s *TgService) generateRandomIDs(n int) []int64 {
	// TODO deprecated rand.Seed(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	randomIDs := make([]int64, n)
	for i := range randomIDs {
		randomIDs[i] = rand.Int63()
	}
	return randomIDs
}

func (s *TgService) getChannel(name string, ctx context.Context, api *tg.Client) (*tg.Channel, error) {
	resolved, err := api.ContactsResolveUsername(ctx, &tg.ContactsResolveUsernameRequest{
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

func (s *TgService) getMessages(channel *tg.Channel, ctx context.Context, api *tg.Client) ([]tg.MessageClass, error) {
	//OffsetID: lastMessageID,
	messages, err := api.MessagesGetHistory(ctx, &tg.MessagesGetHistoryRequest{
		Peer: &tg.InputPeerChannel{
			ChannelID:  channel.ID,
			AccessHash: channel.AccessHash,
		},
		Limit: 2, // Количество сообщений
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка получения сообщений: %w", err)
	}

	channelMessages := messages.(*tg.MessagesChannelMessages)
	// Если новых сообщений нет — выходим
	if len(channelMessages.Messages) == 0 {
		fmt.Print("новых сообщений нет")
		return nil, nil
	}

	// Сортируем сообщения по ID (по возрастанию)
	sort.Slice(channelMessages.Messages, func(i, j int) bool {
		msgI, okI := channelMessages.Messages[i].(*tg.Message)
		msgJ, okJ := channelMessages.Messages[j].(*tg.Message)
		if !okI || !okJ {
			return false
		}
		return msgI.ID < msgJ.ID
	})
	return channelMessages.Messages, nil
}

type EmptyService struct{}

func NewEmptyService() *EmptyService {
	return &EmptyService{}
}

func (s *EmptyService) ProcessChannel(ctx context.Context, log *zap.Logger, ch config.ConfChannel) error {
	log.Debug("ProcessChannel", zap.Any("channel", ch))
	return nil
}
