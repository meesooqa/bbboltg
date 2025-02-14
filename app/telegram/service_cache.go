package telegram

import (
	"context"
	"fmt"
	"time"

	"github.com/gotd/td/tg"
	"github.com/patrickmn/go-cache"

	"github.com/meesooqa/bbboltg/app/config"
)

// CachingService is Proxi/Decorator to cache some telegram.Service data
type CachingService struct {
	tgs   Service
	cache *cache.Cache
}

func NewCachingChannelInfoProvider(tgs Service, ttl time.Duration) *CachingService {
	return &CachingService{
		tgs:   tgs,
		cache: cache.New(ttl, 2*ttl),
	}
}

func (s *CachingService) GetChannel(ctx context.Context, name string) (*tg.Channel, error) {
	key := fmt.Sprintf("telegram:channel:%s", name)
	if cached, found := s.cache.Get(key); found {
		if channel, ok := cached.(*tg.Channel); ok {
			fmt.Println("Данные получены из кэша.")
			return channel, nil
		}
	}

	channel, err := s.tgs.GetChannel(ctx, name)
	if err != nil {
		return nil, err
	}
	if channel != nil {
		s.cache.Set(key, channel, cache.DefaultExpiration)
		fmt.Println("Данные получены из API и сохранены в кэш.")
	}
	return channel, nil
}

//func (s *CachingService) Init(ctx context.Context) error {
//	return s.tgs.Init(ctx)
//}

func (s *CachingService) ProcessChannel(ctx context.Context, ch config.ConfChannel) error {
	return s.tgs.ProcessChannel(ctx, ch)
}

func (s *CachingService) GetMessages(ctx context.Context, channel *tg.Channel) ([]TelegramMessage, error) {
	return s.tgs.GetMessages(ctx, channel)
}

func (s *CachingService) ForwardMessage(ctx context.Context, msg TelegramMessage, channelFrom *tg.Channel, channelTo *tg.Channel) error {
	return s.tgs.ForwardMessage(ctx, msg, channelFrom, channelTo)
}
