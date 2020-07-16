package main

import (
	"./nuture"
	"os"
)
import "./linkslack"

var msg = "test"

var (
	apitoken     = os.Getenv("SLACKAPI_TOKEN")
	slackChannel = "general"
)

func main() {
	nuture.Curltest()
	linkslack.SendMessage(msg, apitoken, slackChannel)

}
