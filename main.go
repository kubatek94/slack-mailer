package main

import (
	"encoding/json"
	"github.com/kubatek94/slack-mailer/mailer"
	"log"
	"net/mail"
	"os"
)

func main() {
	email, err := mail.ReadMessage(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	message := mailer.ForwardMail(config.WebhookUrl, email)

	if config.Debug {
		if content, err := json.MarshalIndent(message, "", "  "); err != nil {
			log.Fatalf("failed to JSON encode message for debug purposes: %v", err)
		} else {
			log.Printf("POST %s\n%s\n", config.WebhookUrl, content)
		}
	}
}
