package filter

import "github.com/meesooqa/bbboltg/app/reader"

func FilterStringables(items []reader.Stringable, condition func(string) bool) <-chan reader.Stringable {
	ch := make(chan reader.Stringable)
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
