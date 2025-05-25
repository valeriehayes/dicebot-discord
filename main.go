package main

import (
	"context"
	"log"
	"os"
	"regexp"

	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/v3/bot"
	"github.com/diamondburned/arikawa/v3/session"
	"github.com/joho/godotenv"
)

type DiceBot struct {
	*bot.Context
}

var dicePattern = regexp.MustCompile(`(?i)^/roll\s+(\d+)d(\d+)$`)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatalln("BOT_TOKEN not found in environment")
	}

	s := session.New("Bot " + token)
	b, err := bot.New(s, &DiceBot{})
	if err != nil {
		log.Fatalln("Failed to create bot: ", err)
	}

	if err := b.Open(context.Background()); err != nil {
		log.Fatalln("Failed to open bot: ", err)
	}
	defer b.Close()

	select {}
}
