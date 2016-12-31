package main

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	initConfig()
	// Testing twitter API
	// Consumer key, consumer secret
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	// Token, token secret
	token := oauth1.NewToken(token, tokenSecret)

	// httpClient will automatically authorize http.Request's
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
	params := &twitter.StreamSampleParams{
		StallWarnings: twitter.Bool(true),
	}

	stream, err := client.Streams.Sample(params)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	// demultiplexer
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		msg := ""
		if tweet.Coordinates != nil {
			coords := tweet.Coordinates.Coordinates
			msg = fmt.Sprintf("%s Lat: %f Long: %f", msg, coords[0], coords[1])
		}

		if tweet.Place != nil {
			msg = fmt.Sprintf("%s  Country: %s", msg, tweet.Place.Country)
		}

		if msg != "" {
			logger.Info("%s", msg)
		}
	}

	for message := range stream.Messages {
		demux.Handle(message)
	}

}
