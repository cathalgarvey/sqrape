package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alecthomas/kingpin"
	"github.com/cathalgarvey/sqrape"
)

var (
	profile = kingpin.Arg("profile", "Profile to fetch").Required().String()
)

// Tweet data to parse.
type Tweet struct {
	Author  string `csss:"div.original-tweet;attr=data-screen-name"`
	TweetID int64  `csss:"div.original-tweet;attr=data-tweet-id"`
	Content string `csss:"p.js-tweet-text;text"`
}

// TwitterProfile to parse.
type TwitterProfile struct {
	Tweets []Tweet `csss:"li.js-stream-item;obj"`
}

func main() {
	kingpin.Parse()
	resp, err := http.Get("https://twitter.com/" + *profile)
	if err != nil {
		log.Fatal(err)
	}
	tp := new(TwitterProfile)
	err = sqrape.ExtractHTMLReader(resp.Body, tp)
	if err != nil {
		log.Fatal(err)
	}
	for _, tweet := range tp.Tweets {
		fmt.Printf("@%s: %s\n", tweet.Author, tweet.Content)
	}
}
