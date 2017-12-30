package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/reujab/wallpaper"
	"github.com/turnage/graw/reddit"
)

func randomPost(posts []*reddit.Post) (post *reddit.Post) {
	rand.Seed(int64(time.Now().Second()))
	return posts[rand.Intn(len(posts))]
}

func sortedPost(posts []*reddit.Post) (post *reddit.Post) {
	post = posts[0]
	for x := range posts {
		if posts[x].Ups >= post.Ups {
			post = posts[x]
		}
	}
	return post
}

func setWallpaper(post *reddit.Post) error {
	url := post.URL
	if strings.HasPrefix(url, "https://imgur.com/") {
		url = strings.TrimPrefix(url, "https://imgur.com/")
		url = "https://i.imgur.com/" + url + ".jpg"
	}
	err := wallpaper.SetFromURL(url)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// Initialize the bot, using the credintials from a local file
	bot, err := reddit.NewBotFromAgentFile("credintials.txt", 0)
	if err != nil {
		panic(err)
	}

	// Retrieve a harvest from the wallpaper reddit
	for {
		fmt.Println("Polling r/wallpapers")
		harvest, err := bot.Listing("/r/wallpapers", "")
		if err != nil {
			panic(err)
		}
		posts := harvest.Posts

		post := randomPost(posts)

		err = setWallpaper(post)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(15 * time.Second)
	}
}
