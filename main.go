package main

import (
	"fmt"
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
	con, acc, _ := twitter.LoadConfig("secrets.json")
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
