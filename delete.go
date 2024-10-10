package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dghubble/oauth1"
)

// API credentials
const (
	consumerKey    = "ZHHPRBd3NB1772nQNjxGUQxZw"
	consumerSecret = "BxGMNsQrnb2OZVTynmZWHw1V1HkcQmFKi328MONVsr7nxFXan2"
	accessToken    = "1844361911883730954-A4u92kbu0rmpuwbxSmohHUzpky50sI"
	accessSecret   = "X0yT7uNXL2Uzr0Nh2MLBhHyxeEWKHQrrnOXr9bxBWUzWb"
)

func main() {
	// Tweet ID you want to delete
	tweetID := "1844468089498272198"

	// Deleting the tweet
	err := deleteTweet(tweetID)
	if err != nil {
		log.Fatalf("Tweet successfully deleted, ID: %s\n", tweetID)
	}
	//fmt.Printf("Tweet successfully deleted, ID: %s\n", tweetID)
}

// function for deleting tweet by its ID
func deleteTweet(tweetID string) error {
	// OAuth1 authentication
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// Create HTTP client with OAuth1
	ctx := context.Background()
	httpClient := config.Client(ctx, token)

	// Create a new DELETE request
	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://api.twitter.com/2/tweets/%s", tweetID), nil)
	if err != nil {
		return fmt.Errorf("failed to create delete request: %w", err)
	}

	// Execute the DELETE request
	response, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute delete request: %w", err)
	}
	defer response.Body.Close()

	// Check the response status
	if response.StatusCode != http.StatusNoContent {
		// Read the response body for additional error info
		body, _ := ioutil.ReadAll(response.Body)
		return fmt.Errorf("failed to delete tweet, status: %s, body: %s", response.Status, string(body))
	}

	return nil
}
