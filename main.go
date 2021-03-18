package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/robfig/cron"
)

func main() {
	when := strings.Split(os.Args[1], ":")
	text := os.Args[2]
	fmt.Printf("Planning tweet at %sh %sm: %s\n", when[0], when[1], text)

	creds := credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	client, err := getClient(&creds)
	if err != nil {
		fmt.Printf("Error getting Twitter client: %s", err)
	}

	c := cron.New()
	c.AddFunc("0 "+when[1]+" "+when[0]+" * *", func() {
		tweet, resp, err := client.Statuses.Update(text, nil)

		if err != nil {
			fmt.Printf("Error tweeting: %s\n", err)
		}

		fmt.Printf("Tweeted: %+v\n", tweet.Text)
		fmt.Printf("Status: %+v\n", resp.Status)
	})
	c.Start()

	for {
		time.Sleep(5 * time.Minute)
	}
}

type credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func getClient(creds *credentials) (*twitter.Client, error) {
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// we can retrieve the user and verify if the credentials
	// we have used successfully allow us to log in!
	_, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	return client, nil
}
