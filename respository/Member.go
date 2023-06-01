package respository

import (
	"fantracer/driver"
	"fantracer/models"
)


func InsertMember(member models.Member) error {
	_,err := driver.Engine.Insert(member)
	return err
}