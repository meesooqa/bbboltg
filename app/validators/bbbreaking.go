// Package validators has all validators for Telegram Channel "Bbbreaking"
package validators

type Bbbreaking struct{}

func NewBbbreaking() *Bbbreaking {
	return &Bbbreaking{}
}

func (v Bbbreaking) IsBolt(s string) bool {
	runes := []rune(s)
	if len(runes) == 0 {
		return false
	}
	// https://unicodeplus.com/U+26A1 ⚡️High Voltage Sign
	return runes[0] == '\u26A1'
}

func (v Bbbreaking) IsImportant(s string) bool {
	runes := []rune(s)
	if len(runes) == 0 {
		return false
	}
	// https://unicodeplus.com/U+2757 ❗Heavy Exclamation Mark Symbol
	return runes[0] == '\u2757'
}
