package linkslack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

var msg = ""

var (
	apitoken     = os.Getenv("SLACKAPI_TOKEN")
	slackChannel = "general"
)

func Dummy() {
}
func SendMessage(msg, apitoken, slackChannel string) error {
	api := slack.New(apitoken)

	_, _, err := api.PostMessage(slackChannel, slack.MsgOptionText(msg, false))

	return err
}

func ReciveMessage(apitoken, slackChannel string) string {
	http.HandleFunc("/slack/events", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		eventsAPI, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		switch eventsAPI.Type {
		/*
			Events API利用のためのURL検証
			チャレンジリクエストに対し，リクエストボディ内JSONに含まれるChallengeの値をレスポンスボディに詰めて応答
		*/
		case slackevents.URLVerification:
			var res *slackevents.ChallengeResponse
			if err := json.Unmarshal(body, &res); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/plain")
			if _, err := w.Write([]byte(res.Challenge)); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

		case slackevents.CallbackEvent:
			innerEvent := eventsAPI.InnerEvent
			switch event := innerEvent.Data.(type) {
			case *slackevents.AppMentionEvent:
				message := strings.Split(event.Text, " ")
				if len(message) < 2 {
					w.WriteHeader(http.StatusBadRequest)
					log.Println("Bad Request")
					return
				}
				msg = message[1]
			}

		}
	}) //func（httpヘッダ内）終了
	return msg
}

func testsend() {
	fmt.Println("test")
	message := "Hello from golang"

	if err := SendMessage(message, apitoken, slackChannel); err != nil {
		log.Fatal(err)
	}
}
