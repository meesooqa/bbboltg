package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterBolt(t *testing.T) {
	// High Voltage Sign
	f := NewPrefixFilter("⚡️")

	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"b Строка начинается с эмодзи", "⚡️Пример строки", true},
		{"b Строка не начинается с эмодзи", "Пример ⚡️ строки", false},
		{"b Пустая строка", "", false},
		{"b Строка с только эмодзи", "⚡️", true},
		{"b Пробел перед эмодзи", " ⚡️Пример", false},
		{"b Настоящая новость", "⚡️Apple выпустила обновление ПО для iPhone и iPad из-за возможной уязвимости iOS", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := f.IsValid(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFilterImportant(t *testing.T) {
	// Heavy Exclamation Mark Symbol
	f := NewPrefixFilter("❗️")

	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"Эмодзи с вариационным селектором", "❗️Это тестовая строка", true},
		//{"Эмодзи без вариационного селектора", "❗Это тестовая строка", true}, // rune
		{"Строка не начинается с ❗", "⚡❗Это тестовая строка", false},
		{"Строка без эмодзи", "Это тестовая строка", false},
		{"Пустая строка", "", false},
		{"Пробельный символ перед эмодзи", " ❗️Это тестовая строка", false},
		{"i Настоящая новость", "❗️Посадочный модуль Blue Ghost покинул орбиту Земли", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := f.IsValid(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
