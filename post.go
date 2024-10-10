package main

import (
    "fmt"
    "log"
    "os"

)

func main() {

    //Setting up the credentials
    consumerKey := "WWpTZGYyeG1vVjdrMjM1MjZ4RHQ6MTpjaQ"
    consumerSecret := "5VTkfGZMEaaLtsS_8Z2_ykpILOv8hauO1pIGLHilmxZCsJZ7qb"
    accessToken := "1844361911883730954-baCVOYXcbMVEGQpulmvMN0SYA6Tjnm"
    accessSecret := "2gZiDeLpuQde1xM2yJAP4fVcWSeWEbVoaNkpT9hqw7p0T"

    // OAuth1 authentication
    config := oauth1.NewConfig(consumerKey, consumerSecret)
    token := oauth1.NewToken(accessToken, accessSecret)
    httpClient := config.Client(oauth1.NoContext, token)

    // Creating a Client
    client := twitter.NewClient(httpClient)

    // Creating a Tweet
    tweet, _, err := client.Statuses.Update("Hello from Twitter API!", nil)
    if err != nil {
        log.Fatalf("Failed to post tweet: %v", err)
    }

    // Printing the Tweet
    fmt.Printf("Tweet posted: %s\n", tweet.Text)
}
