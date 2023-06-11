package services

import (
	"core/driver"
	"core/models"
)

func InsertGroupInfo(groupInfo models.GroupInfo) error {
	_, err := driver.Engine.Insert(groupInfo)

	return err
}
