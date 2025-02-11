package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBolt(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"b Строка начинается с эмодзи", "⚡ Пример строки", true},
		{"b Строка не начинается с эмодзи", "Пример ⚡ строки", false},
		{"b Пустая строка", "", false},
		{"b Строка с только эмодзи", "⚡", true},
		{"b Пробел перед эмодзи", " ⚡ Пример", false},
		{"b Настоящая новость", "⚡️Apple выпустила обновление ПО для iPhone и iPad из-за возможной уязвимости iOS", true},
	}

	v := NewBbbreaking()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := v.IsBolt(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsImportant(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "Эмодзи с вариационным селектором",
			input: "❗️Это тестовая строка", // обычно представляет [U+2757, U+FE0F, ...]
			want:  true,
		},
		{
			name:  "Эмодзи без вариационного селектора",
			input: "❗Это тестовая строка", // только U+2757
			want:  true,
		},
		{
			name:  "Строка не начинается с ❗",
			input: "⚡❗Это тестовая строка",
			want:  false,
		},
		{
			name:  "Строка без эмодзи",
			input: "Это тестовая строка",
			want:  false,
		},
		{
			name:  "Пустая строка",
			input: "",
			want:  false,
		},
		{
			name:  "Пробельный символ перед эмодзи",
			input: " ❗️Это тестовая строка",
			want:  false,
		},
		{
			name:  "i Настоящая новость",
			input: "❗️Посадочный модуль Blue Ghost покинул орбиту Земли",
			want:  true,
		},
	}

	v := NewBbbreaking()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := v.IsImportant(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
