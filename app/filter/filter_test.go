package filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Мок-структура для имитации интерфейса Message
type MockMessage string

// Реализация метода String() для MockMessage
func (m MockMessage) String() string {
	return string(m)
}

// Функция для создания мок-сообщений
func createMockMessages(contents ...string) []Stringable {
	messages := make([]Stringable, len(contents))
	for i, content := range contents {
		messages[i] = MockMessage(content)
	}
	return messages
}

// Функция для чтения всех значений из канала
func readAllFromChannel(ch <-chan Stringable) []Stringable {
	var result []Stringable
	for msg := range ch {
		result = append(result, msg)
	}
	return result
}

func TestFilterStringables(t *testing.T) {
	tests := []struct {
		name      string
		input     []Stringable
		condition func(string) bool
		expected  []string
	}{
		{
			name:      "Empty input slice",
			input:     createMockMessages(),
			condition: func(s string) bool { return len(s) > 3 },
			expected:  []string{},
		},
		{
			name:      "No matches",
			input:     createMockMessages("a", "ab", "abc"),
			condition: func(s string) bool { return len(s) > 5 },
			expected:  []string{},
		},
		{
			name:      "Partial matches",
			input:     createMockMessages("short", "longer_message", "tiny", "huge_message"),
			condition: func(s string) bool { return len(s) > 5 },
			expected:  []string{"longer_message", "huge_message"},
		},
		{
			name:      "All matches",
			input:     createMockMessages("abcdef", "ghijklmn", "opqrstu"),
			condition: func(s string) bool { return len(s) > 3 },
			expected:  []string{"abcdef", "ghijklmn", "opqrstu"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := FilterStringables(tt.input, tt.condition)
			result := readAllFromChannel(ch)

			// convert to []string
			var resultStrings []string
			for _, msg := range result {
				resultStrings = append(resultStrings, msg.String())
			}

			assert.Equal(t, len(tt.expected), len(resultStrings))
			for i := range tt.expected {
				assert.Equal(t, tt.expected[i], resultStrings[i])
			}
		})
	}
}
