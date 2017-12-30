package main

import (
	"fmt"

	"github.com/turnage/graw/reddit"
)

func main() {
	bot, err := reddit.NewBotFromAgentFile("credintials.txt", 0)
	if err != nil {
		panic(err)
	}

	harvest, err := bot.Listing("/r/wallpapers", "")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v: %T", harvest.Posts[0], harvest.Posts[0])
}
