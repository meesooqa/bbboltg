package filter

// github.com/gotd/td/tg.MessageClass
type Stringable interface {
	String() string
}

func FilterStringables(items []Stringable, condition func(string) bool) <-chan Stringable {
	ch := make(chan Stringable)
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
