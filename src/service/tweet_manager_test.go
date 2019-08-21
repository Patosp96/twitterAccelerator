package service_test

import (
	"github.com/Patosp96/twitterAccelerator/src/domain"
	"github.com/Patosp96/twitterAccelerator/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweet()
	if publishedTweet.User != user &&
		publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, publishedTweet.User, publishedTweet.Text)
	}
	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestWithoutUserIsNotPublished(t *testing.T) {

	var tweet *domain.Tweet
	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err != nil && err.Error() != "user is required" {
		t.Error("Excpected error is user is required")
	}
}

func TestWithoutTextIsNotPublished(t *testing.T) {

	var tweet *domain.Tweet
	user := "User123"
	var text string

	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err != nil && err.Error() != "text is required" {
		t.Error("Excpected error is text is required")
	}

}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {

	var tweet *domain.Tweet
	user := "User123"
	text := "123456789123456789123456789" +
		    "123456789123456789123456789" +
			"123456789123456789123456789" +
			"123456789123456789123456789" +
			"123456789123456789123456789" +
			"123456"

	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err != nil && err.Error() != "text is longest than 140" {
		t.Error("Excpected error is text is longest than 140")
	}
}




