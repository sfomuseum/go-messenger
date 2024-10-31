package message

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-messenger"
)

type RunOptions struct {
	DeliveryAgent messenger.DeliveryAgent
	From          string
	To            []string
	Subject       string
	Body          string
}

func DeriveOptionsFromFlagSet(ctx context.Context, fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	agent, err := messenger.NewMultiDeliveryAgentWithURIs(ctx, agent_uris...)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive agents, %w", err)
	}

	var body string

	args := fs.Args()

	if len(args) == 1 && args[0] == "-" {

		v, err := io.ReadAll(os.Stdin)

		if err != nil {
			return nil, fmt.Errorf("Failed to read from STDIN")
		}

		body = string(v)

	} else {

		body = strings.Join(args, " ")
	}

	opts := &RunOptions{
		DeliveryAgent: agent,
		From:          from,
		To:            to,
		Subject:       subject,
		Body:          body,
	}

	return opts, nil
}
