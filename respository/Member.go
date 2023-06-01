package respository

import (
	"fantracer/driver"
	"fantracer/models"
	"fmt"
	"strings"
)


func InsertMember(member models.Member) error {
	_,err := driver.Engine.Insert(member)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil
		}
		fmt.Println(err,member)
	}

	return err
}

func FindMemberInGroups(memberID int ) ([]models.Member, error) {
	var memberRecord []models.Member

	err := driver.Engine.Where("i_d = ?",memberID).Find(&memberRecord)
	if err != nil {
		return nil,err
	}

	return memberRecord, nil
}