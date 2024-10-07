package main

import (
	"context"
	"log"

	"github.com/sfomuseum/go-messenger/app/message"
)

func main() {

	ctx := context.Background()
	err := message.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run message app, %v", err)
	}
}
