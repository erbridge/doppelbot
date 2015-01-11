package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/ChimeraCoder/anaconda"
	"github.com/erbridge/gotwit/bot"
	"github.com/erbridge/gotwit/callback"
	"github.com/erbridge/gotwit/twitter"
)

func createRepeatCallback(b *bot.Bot) func(anaconda.Tweet) {
	return func(t anaconda.Tweet) {
		sender := t.User.ScreenName
		botName := b.ScreenName()

		if sender == botName {
			return
		}

		raw := t.Text
		fmt.Println(sender + " sent: " + raw)

		text := regexp.MustCompile("@"+botName).ReplaceAllString(raw, "@"+sender)

		b.Reply(t, text, t.PossiblySensitive)
		fmt.Println(" - replied: " + text)
	}
}

func main() {
	var (
		con twitter.ConsumerConfig
		acc twitter.AccessConfig
	)

	f := "secrets.json"
	if _, err := os.Stat(f); err == nil {
		con, acc, _ = twitter.LoadConfigFile(f)
	} else {
		con, acc, _ = twitter.LoadConfigEnv()
	}

	b := bot.New("doppelbot", con, acc)

	repeatTweet := createRepeatCallback(&b)
	b.RegisterCallback(callback.Reply, repeatTweet)
	b.RegisterCallback(callback.Mention, repeatTweet)

	b.Start()

	// Stop the program from exiting.
	var input string
	fmt.Scanln(&input)

	b.Stop()
}
