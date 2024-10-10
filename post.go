package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

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
	// OAuth1 configuration
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	client := config.Client(oauth1.NoContext, token)

	// Tweet content
	tweetData := map[string]string{"text": "Hello, This is a test API!"}
	tweetJSON, _ := json.Marshal(tweetData)

	// Post the tweet
	resp, err := client.Post("https://api.twitter.com/2/tweets", "application/json", bytes.NewBuffer(tweetJSON))
	if err != nil {
		log.Fatal("Failed to post tweet:", err)
	}
	defer resp.Body.Close()

	// Parse response to get the tweet ID
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	tweetID := result["data"].(map[string]interface{})["id"].(string)
	fmt.Println("Tweet posted, ID:", tweetID)
}
