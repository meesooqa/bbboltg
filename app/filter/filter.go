package filter

import (
	"errors"

	"github.com/meesooqa/bbboltg/app/config"
	"github.com/meesooqa/bbboltg/app/telegram"
)

func NewFilter(ccFilter *config.ConfChannelFilter) (*PrefixFilter, error) {
	if ccFilter.Prefix != "" {
		return newPrefixFilter(ccFilter.Prefix), nil
	}
	return nil, errors.New("prefix is empty")
}

func FilterTelegramMessages(messages []telegram.TelegramMessage, condition func(string) bool) []telegram.TelegramMessage {
	var result []telegram.TelegramMessage
	for _, message := range messages {
		if condition(message.String()) {
			result = append(result, message)
		}
	}
	return result
}

/*
func FilterTelegramMessages(items []telegram.TelegramMessage, condition func(string) bool) <-chan telegram.TelegramMessage {
	ch := make(chan telegram.TelegramMessage)
	go func() {
		defer close(ch)
		for _, item := range items {
			if condition(item.String()) {
				ch <- item
			}
		}
	}()
	return ch
}
*/
