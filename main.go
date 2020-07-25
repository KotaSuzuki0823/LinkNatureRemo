package main

import (
	"./linkslack"
	"./nuture"
	"log"
	"os"
)

var (
	apitoken     = os.Getenv("SLACKAPI_TOKEN")
	slackChannel = "general"
)

func main() {
	result := nuture.Curltest()
	//nuture.Dummy()
	err := linkslack.SendMessage(string(result), apitoken, slackChannel)
	if err != nil {
		log.Fatal(err)
	}

}
