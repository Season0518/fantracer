package utils

import (
	"core/models"
	"encoding/json"
	"os"
)

func ReadMsgConfig() (models.WelcomeMessage, error) {
	filePath, err := ConvertToFullPath("MsgConfig.json")
	if err != nil {
		return models.WelcomeMessage{}, err
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return models.WelcomeMessage{}, err
	}

	var welcomeMessage models.WelcomeMessage
	err = json.Unmarshal(data, &welcomeMessage)
	if err != nil {
		return models.WelcomeMessage{}, err
	}

	return welcomeMessage, nil
}

func ReadMediaURL() ([]string, error) {
	welcomeMessage, err := ReadMsgConfig()
	if err != nil {
		return nil, err
	}
	return welcomeMessage.MediaURL, nil
}

func ReadWelcomeText() (string, error) {
	welcomeMessage, err := ReadMsgConfig()
	if err != nil {
		return "", err
	}

	return welcomeMessage.Text, nil
}
