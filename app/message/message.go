package message

// echo "TESTING" | ./bin/message -from aaron -to aaron -agent-uri beeep:// -agent-uri stdout:// -subject testing -

import (
	"context"
	"flag"
	"log/slog"
	"sync"

	"github.com/sfomuseum/go-messenger"
)

func Run(ctx context.Context) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	opts, err := DeriveOptionsFromFlagSet(ctx, fs)

	if err != nil {
		return err
	}

	return RunWithOptions(ctx, opts)
}

func RunWithOptions(ctx context.Context, opts *RunOptions) error {

	if len(opts.To) == 0 {

		msg := &messenger.Message{
			From:    opts.From,
			Subject: opts.Subject,
			Body:    opts.Body,
		}

		err := opts.DeliveryAgent.DeliverMessage(ctx, msg)

		if err != nil {
			slog.Error("Failed to deliver message", "error", err)
			return err
		}

		slog.Debug("Message delivered")
		return nil
	}

	wg := new(sync.WaitGroup)

	for _, to := range opts.To {

		wg.Add(1)

		go func(to string) {

			defer wg.Done()

			msg := &messenger.Message{
				To:      to,
				From:    opts.From,
				Subject: opts.Subject,
				Body:    opts.Body,
			}

			err := opts.DeliveryAgent.DeliverMessage(ctx, msg)

			if err != nil {
				slog.Error("Failed to deliver message", "to", to, "error", err)
			}

			slog.Debug("Message delivered", "to", to)
		}(to)
	}

	wg.Wait()
	return nil
}
