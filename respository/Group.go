package respository

import (
	"fantracer/driver"
	"fantracer/models"
)

func InsertGroup(group models.Group) error {
	_, err := driver.Engine.Insert(group)

	return err
}

