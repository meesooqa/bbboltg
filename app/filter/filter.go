package filter

import "github.com/meesooqa/bbboltg/app/reader"

func FilterTelegramMessages(items []reader.TelegramMessage, condition func(string) bool) <-chan reader.TelegramMessage {
	ch := make(chan reader.TelegramMessage)
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
