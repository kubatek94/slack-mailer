package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

var config struct {
	Path       string
	Debug      bool   `json:"debug"`
	WebhookUrl string `json:"webhook_url"`
}

func init() {
	var content []byte
	content, config.Path = readConfigFile()

	if err := json.Unmarshal(content, &config); err != nil {
		panic(fmt.Sprintf("in \"%s\": %s", config.Path, err))
	}

	if err := validateWebhookUrl(); err != nil {
		panic(fmt.Sprintf("in \"%s\": %s", config.Path, err))
	}
}

func readConfigFile() ([]byte, string) {
	basename := "slack-mailer.json"
	dirs := []string{}

	if homeDir, err := os.UserHomeDir(); err == nil {
		dirs = append(dirs, homeDir+"/"+basename)
	}

	if configDir, err := os.UserConfigDir(); err == nil {
		dirs = append(dirs, configDir+"/"+basename)
	}

	dirs = append(dirs, basename, "/etc/"+basename)

	for _, filename := range dirs {
		if content, err := ioutil.ReadFile(filename); err == nil {
			return content, filename
		}
	}

	panic(fmt.Sprintf("No config found in (\"%s\")", strings.Join(dirs, "\", \"")))
}

func validateWebhookUrl() error {
	if config.WebhookUrl == "" {
		return fmt.Errorf("webhook_url config value must be set")
	}

	parsedWebhookUrl, err := url.Parse(config.WebhookUrl)
	if err != nil {
		return err
	}

	if parsedWebhookUrl.Host == "" {
		return fmt.Errorf("webhook_url config value must be a full URL, including host")
	}

	return nil
}
