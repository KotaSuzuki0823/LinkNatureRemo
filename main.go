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
	//nuture.Curltest()
	nuture.Dummy()
	err := linkslack.SendMessage("test", apitoken, slackChannel)
	if err != nil {
		log.Fatal(err)
	}

}
