package mailer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/mail"
	"strings"
)

func ForwardMail(webhookUrl string, message *mail.Message) {
	postMessage(webhookUrl, mailToMessage(message))
}

func mailToMessage(m *mail.Message) Message {
	sections := make([]interface{}, 0, 7)

	sections = append(sections,
		section(
			mrkdwn(fmt.Sprintf(":bangbang:  %s", m.Header.Get("Subject")))),
		divider(),
		section(
			mrkdwn(fmt.Sprintf(":outbox_tray: *From*\n%s", m.Header.Get("From"))),
			mrkdwn(fmt.Sprintf(":inbox_tray: *To*\n%s", m.Header.Get("To")))),
		divider(),
	)

	if contextSection, numElements := headersToContext(m.Header); numElements > 0 {
		sections = append(sections,
			contextSection,
			divider(),
		)
	}

	text := bodyToText(m.Body)

	sections = append(sections,
		section(
			mrkdwn(text)),
	)

	return message(text, sections...)
}

func headersToContext(headers mail.Header) (map[string]interface{}, uint) {
	results := make([]interface{}, 0, math.Max(0, float64(len(headers)-3)))
	for k, items := range headers {
		switch k {
		case "To", "From", "Subject":
			continue
		default:
			results = append(results,
				mrkdwn(fmt.Sprintf("*%s* %s", k, strings.Join(items, ", "))))
		}
	}
	return context(results...), uint(len(results))
}

func bodyToText(reader io.Reader) string {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(reader); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func postMessage(webhookUrl string, message Message) {
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
		log.Fatalf("Message failed with status [%s] and content: %s\n", resp.Status, body)
	}
}
