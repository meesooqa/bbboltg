package proc

import (
	"context"
	"testing"
	"time"

	"github.com/gotd/td/tg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/meesooqa/bbboltg/app/config"
	"github.com/meesooqa/bbboltg/app/telegram"
)

type mockService struct {
	calls int
	err   error
}

func (s *mockService) ProcessChannel(ctx context.Context, ch config.ConfChannel) error {
	s.calls++
	return s.err
}

func (s *mockService) GetChannel(ctx context.Context, name string) (*tg.Channel, error) {
	return nil, nil
}

func (s *mockService) GetMessages(ctx context.Context, channel *tg.Channel) ([]telegram.TelegramMessage, error) {
	return nil, nil
}

func (s *mockService) ForwardMessage(ctx context.Context, msg telegram.TelegramMessage, channelFrom *tg.Channel, channelTo *tg.Channel) error {
	return nil
}

type mockReader struct {
	calls int
	err   error
}

func (r *mockReader) Execute(ctx context.Context, channel *tg.Channel) ([]telegram.TelegramMessage, error) {
	r.calls++
	return nil, r.err
}

type mockWriter struct{}

func (w *mockWriter) Execute(ctx context.Context, messages []telegram.TelegramMessage, src *tg.Channel, dest *tg.Channel) error {
	return nil
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
	mr := &mockReader{}
	mw := &mockWriter{}
	proc := &Processor{
		Conf:            conf,
		log:             zap.NewNop(),
		telegramService: mockTgService,
		reader:          mr,
		writer:          mw,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Millisecond)
	defer cancel()

	err := proc.Do(ctx)

	assert.ErrorIs(t, err, context.DeadlineExceeded)
	require.GreaterOrEqual(t, mockTgService.calls, 2, "There must be at least two Reader.Execute calls")
}
