package main

import (
	"fmt"
	"strings"

	"github.com/turnage/graw/reddit"
)

func main() {
	// Initialize the bot, using the credintials from a local file
	bot, err := reddit.NewBotFromAgentFile("credintials.txt", 0)
	if err != nil {
		panic(err)
	}

	// Retrieve a harvest from a specific reddit
	harvest, err := bot.Listing("/r/wallpapers", "")
	if err != nil {
		panic(err)
	}

	var post *reddit.Post
	for x := range harvest.Posts {
		if x <= 0 {
			post = harvest.Posts[x]
		} else {
			if harvest.Posts[x].Ups >= post.Ups {
				post = harvest.Posts[x]
			} else {
				continue
			}
		}
	}

	if strings.HasPrefix(post.URL, "https://imgur.com/StsRgCA") {

	}
	fmt.Println(post.URL)
	// Set the wallpaper as the URL aquired from the sorted posts
	// err = wallpaper.SetFromURL(post.URL)
	// if err != nil {
	// 	panic(err)
	// }

}
