package main

import (
	"github.com/kubatek94/slack-mailer/mailer"
	"log"
	"net/mail"
	"net/url"
	"os"
)

func main() {
	message, err := mail.ReadMessage(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	mailer.ForwardMail(getWebhookUrl(), message)
}

func getWebhookUrl() string {
	webhookUrl := os.Getenv("WEBHOOK_URL")
	if webhookUrl == "" {
		log.Fatal("WEBHOOK_URL environment variable must be provided")
	}

	parsedWebhookUrl, err := url.Parse(webhookUrl)
	if err != nil {
		log.Fatal(err)
	}

	if parsedWebhookUrl.Host == "" {
		log.Fatal("WEBHOOK_URL environment variable must be a full URL, including host")
	}

	return webhookUrl
}
