package service

import "github.com/Patosp96/twitterAccelerator/src/domain"

var Tweet *domain.Tweet

func PublishTweet(tweet *domain.Tweet) {
	Tweet = tweet
}

func GetTweet() *domain.Tweet {
	return Tweet
}