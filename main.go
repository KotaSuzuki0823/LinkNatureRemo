package main

import (
	"./linkslack"
	"log"
	"os"
)

var msg = "test"

var (
	apitoken     = os.Getenv("SLACKAPI_TOKEN")
	slackChannel = "general"
)

func main() {
	//nuture.Curltest()
	err := linkslack.SendMessage(msg, apitoken, slackChannel)
	if err != nil {
		log.Fatal(err)
	}

}
