package telegram

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/meesooqa/bbboltg/app/config"
)

type mockService struct {
	calls int
	err   error
}

func (s *mockService) ProcessChannel(ctx context.Context, log *zap.Logger, ch config.ConfChannel) error {
	s.calls++
	return s.err
}

func TestProcessor_Do(t *testing.T) {
	channels := make(map[string]config.ConfChannel)
	channels["name1"] = config.ConfChannel{From: "from1", To: "to1"}
	channels["name2"] = config.ConfChannel{From: "from2", To: "to2"}

	conf := &config.Conf{
		UpdateInterval: 10 * time.Millisecond,
		Channels:       channels,
	}

	mockTgService := &mockService{}
	proc := &Processor{
		Conf:            conf,
		TelegramService: mockTgService,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Millisecond)
	defer cancel()

	log := zap.NewNop()

	err := proc.Do(ctx, log)

	assert.ErrorIs(t, err, context.DeadlineExceeded)
	require.GreaterOrEqual(t, mockTgService.calls, 2, "There must be at least two processChannels calls")
}
