package controllers
/*
curl -X POST --data-urlencode 'payload={"channel": "#bottest", "username": "webhookbot", "text": "This is posted to #bottest and comes from a bot named webhookbot.", "icon_emoji": ":ghost:"}' https://hooks.slack.com/services/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
*/

import (
	"fmt"
	"encoding/json"
	"net/http"
	"net/url"
)

// FIXME
var incomingHookURL string = "https://hooks.slack.com/services/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"

type payload struct {
	Channel string `json:"channel"`
	Botname string `json:"username"`
	Msg string `json:"text"`
	Icon string `json:"icon_emoji"`
}

func CreateIncomingRequestJSON(channel string, botname string, msg string, icon string) []byte {
	data := &payload{
				Channel:	channel,
				Botname:	botname,
				Msg:		msg,
				Icon:		icon,
			}
	jsonBytes,err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
	}
	return jsonBytes
}

func TestPost() error {
	jsonBytes := CreateIncomingRequestJSON("#bottest", "gotaro", "test from golang @wakame", ":go:")
	res, err := http.PostForm(incomingHookURL, url.Values{
		"payload": { string(jsonBytes) },
	})
	fmt.Println(res)
	return err
}
