package main

import "strings"

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for i := 0; i < len(messages); i++ {
		tags := tagger(messages[i])
		if tags == nil {
			tags = []string{}
		}
		messages[i].tags = tags
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	lower := strings.ToLower(msg.content)
	if strings.Contains(lower, "urgent") {
		tags = append(tags, "Urgent")
	}

	if strings.Contains(lower, "sale") {
		tags = append(tags, "Promo")
	}
	return tags
}
