package mailer

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"
)

type Message struct {
	Text string `json:"text"`
	Blocks []interface{} `json:"blocks"`
}

type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type MailSection struct {
	Type string `json:"type"`
	Text Text `json:"text"`
	Fields []Text `json:"fields"`
}

func ForwardMail(webhookUrl string, message *mail.Message) {
	postMessage(webhookUrl, mailToMessage(message))
}

func mailToMessage(message *mail.Message) *Message {
	text := bodyToText(message.Body)
	return &Message{
		Text: text,
		Blocks: []interface{} {
			MailSection{
				Type: "section",
				Text: Text{
					Text: text,
					Type: "mrkdwn",
				},
				Fields: []Text{
					{"mrkdwn", "*To*"},
					{"mrkdwn", "*Subject*"},
					{"plain_text", message.Header.Get("To")},
					{"plain_text", message.Header.Get("Subject")},
				},
			},
		},
	}
}

func bodyToText(reader io.Reader) string {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(reader); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func postMessage(webhookUrl string, message *Message) {
	content, _ := json.Marshal(message)

	//log.Fatal(bodyToText(bytes.NewReader(content)))

	resp, err := http.Post(webhookUrl, "application/json", bytes.NewBuffer(content))
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Fatalf("Message failed with code [%s] and content:\n%s\n", resp.StatusCode, body)
	}
}
