package mailer

import "strings"

type text map[string]string
type Message map[string]interface{}

var escaper = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
)

func mrkdwn(content string) text {
	return text{
		"type": "mrkdwn",
		"text": escaper.Replace(content),
	}
}

func divider() map[string]string {
	return map[string]string{
		"type": "divider",
	}
}

func textSection(text text) map[string]interface{} {
	return map[string]interface{}{
		"type": "section",
		"text": text,
	}
}

func fieldsSection(fields ...text) map[string]interface{} {
	return map[string]interface{}{
		"type":   "section",
		"fields": fields,
	}
}

func section(items ...text) map[string]interface{} {
	if len(items) == 1 {
		return textSection(items[0])
	} else {
		return fieldsSection(items...)
	}
}

func context(elements ...interface{}) map[string]interface{} {
	return map[string]interface{}{
		"type":     "context",
		"elements": elements,
	}
}

func message(text string, blocks ...interface{}) Message {
	return map[string]interface{}{
		"text":   text,
		"blocks": blocks,
	}
}
