# TwitterCronBot

[![Go](https://github.com/jpstrube/twitter-cron-bot/actions/workflows/go.yml/badge.svg)](https://github.com/jpstrube/twitter-cron-bot/actions/workflows/go.yml)

This is a very simple Twitter bot written in Go that can send tweets once a day at a given time. It takes 3 parameters: hour, minute and text.

First you have to set your app credentials in this environment variables:
* CONSUMER_KEY
* CONSUMER_SECRET
* ACCESS_TOKEN
* ACCESS_TOKEN_SECRET

(Visit https://developer.twitter.com/, apply and create an app with write permissions enabled.)

Then you can start it with `go run . <hour>:<minute> "<text>"` (or use a precompiled executable). It will run until you terminate it.

For Windows you can download the .exe and twitterbot.cmd into the same folder, change the .cmd for your needs and doubleclick it.
