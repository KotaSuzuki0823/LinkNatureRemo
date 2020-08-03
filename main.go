package main

import (
	"./linkslack"
	"./nuture"
	"encoding/json"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	apitoken       = os.Getenv("SLACKAPI_TOKEN")
	SIGNING_SECRET = os.Getenv("SIGNING_SECRET")
	ACCESS_TOKEN   = os.Getenv("ACCESS_TOKEN")
	slackChannel   = "general"
)

func service() {
	result := nuture.Curltest()
	//nuture.Dummy()
	err := linkslack.SendMessage(string(result), apitoken, slackChannel)
	if err != nil {
		log.Fatal(err)
	}
}

//https://qiita.com/frozenbonito/items/cf75dadce12ef9a048e9
func main() {
	api := slack.New(ACCESS_TOKEN)

	http.HandleFunc("slack/events", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())

	})
}
