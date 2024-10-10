## INTRODUCTION

The purpose of the assignment is to:

- create a new tweet using the Twitter API's POST statuses/update endpoint.
- delete a tweet using the POST statuses/destroy endpoint.

The another purpose of the assignment is to gain hands-on experience using Golang and GitHub.


## SETUP INSTRUCTIONS

1. Setting up a Twitter Developer account

- Login using your Twitter account or create a new one for testing.
- Visit https://developer.twitter.com/ and apply for a developer account. It may take few minutes to get approved.
- Once approved, create a new project and app.

2. Generating API Keys

- Go to project > app > keys and tokens. 
- Under Consumer Keys, re-generate the API Key and Secret (if you are not able to see full API and secret key).
- Under Authentication Tokens, generate Bearer Token, Access Token and Secret.

3. Run a program

- Create a API program using any programming language of your choice (Python, Go, Node.js, etc.
- Now, install any dependencies needed (e.g., oauth1, go-twitter, etc).
- enter the command to run your program (e.g., for Go program; go run filename.go).

## PROGRAM DETAILS

1. How a program posts a new tweet?

The program sends a POST request to the Twitter API's statuses/update endpoint. This request includes the tweet content as a parameter. The program uses OAuth authentication to ensure that the request is authorized.

2. How a program deletes an existing tweet?

The program sends a POST request to the statuses/destroy/:id endpoint, where :id is the ID of the tweet to be deleted. Similar to posting a tweet, OAuth authentication is used for authorization.

## ERROR HANDLING

- These errors occur when the provided API keys, access tokens, or any other credentials are invalid or missing. HTTP Status Code: 401 Unauthorized
- Exceeding limits will result in a 429 Too Many Requests status code.
- This occurs when the input provided to the API is incorrect or invalid, such as using an invalid tweet ID for deletion. HTTP Status Code: 404 Not Found.

