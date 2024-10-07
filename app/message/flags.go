package message

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
	"github.com/sfomuseum/go-messenger"
)

var agent_uris multi.MultiString
var to multi.MultiString
var subject string
var from string

func DefaultFlagSet() *flag.FlagSet {

	agent_schemes := messenger.AgentSchemes()
	str_schemes := strings.Join(agent_schemes, ", ")

	agent_desc := fmt.Sprintf("One or more known sfomuseum/go-messenger.DeliveryAgent URIs. Valid options are: %s", str_schemes)

	fs := flagset.NewFlagSet("message")
	fs.Var(&agent_uris, "agent-uri", agent_desc)
	fs.StringVar(&from, "from", "", "The name or address of the person or process sending the message.")
	fs.StringVar(&subject, "subject", "", "The subject of the message.")
	fs.Var(&to, "to", "One or more addresses where messages should be delivered.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Command-line tool for delivering messages using or more delivery agents.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options] message\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		fs.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nIf the only message input is \"-\" then data will be read from STDIN.\n\n")
	}

	return fs
}
