package models

type WelcomeMessage struct {
	Text     string   `json:"text"`
	MediaURL []string `json:"media_url"`
}
