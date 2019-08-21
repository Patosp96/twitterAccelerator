package service

import (
	"fmt"
	"github.com/Patosp96/twitterAccelerator/src/domain"
)

var Tweet *domain.Tweet

func PublishTweet(tweet *domain.Tweet) error {
	Tweet = tweet

	if tweet.User == "" {
		 return fmt.Errorf("user is required")
	}

	if tweet.Text == "" {
		return fmt.Errorf("text is required")
	}

	if len(tweet.Text) > 140 {
		return fmt.Errorf("text is longest than 140")
	}

	return nil
}

func GetTweet() *domain.Tweet {
	return Tweet
}