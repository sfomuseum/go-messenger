package messenger

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"strings"
)

type SlogDeliveryAgent struct {
	DeliveryAgent
	level slog.Level
}

func init() {
	ctx := context.Background()
	RegisterDeliveryAgent(ctx, "slog", NewSlogDeliveryAgent)
}

func NewSlogDeliveryAgent(ctx context.Context, uri string) (DeliveryAgent, error) {

	level := slog.LevelInfo

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	q := u.Query()

	if q.Has("level") {

		switch strings.ToLower(q.Get("level")) {
		case "debug":
			level = slog.LevelDebug
		case "info":
			level = slog.LevelInfo
		case "warn":
			level = slog.LevelWarn
		case "error":
			level = slog.LevelError
		default:
			return nil, fmt.Errorf("Invalid or unsupported log level")
		}
	}

	a := &SlogDeliveryAgent{
		level: level,
	}

	return a, nil
}

func (a *SlogDeliveryAgent) DeliverMessage(ctx context.Context, msg *Message) error {

	args := make([]any, 0)

	if msg.Subject != "" {
		args = append(args, "subject", msg.Subject)
	}

	if msg.To != "" {
		args = append(args, "to", msg.To)
	}

	if msg.From != "" {
		args = append(args, "from", msg.From)
	}

	logger := slog.Default()
	logger.Log(ctx, a.level, strings.TrimSpace(msg.Body), args...)
	return nil
}
