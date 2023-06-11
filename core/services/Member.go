package services

import (
	"core/driver"
	"core/models"
	"log"
	"strings"
)

func UpsertMember(member models.Member) error {
	_, err := driver.Engine.Insert(member)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			// 如果插入失败并且错误信息中包含 "Duplicate entry"，则尝试更新
			// _, err = driver.Engine.AllCols().Update(&member, &member)
			_, err = driver.Engine.Where("i_d = ? and group_i_d = ?", member.ID, member.GroupID).AllCols().Update(&member)
			if err != nil {
				log.Fatalf("2 Failed to update member: %v\n", err)
			}
		} else {
			// 如果插入失败并且错误信息中不包含 "Duplicate entry"，则记录错误
			log.Fatalf("1 Failed to insert member: %v\n", err)
		}
	}
	return err
}

func FindMemberInGroups(memberID int64) ([]models.Member, error) {
	var memberRecord []models.Member

	err := driver.Engine.Where("i_d = ?", memberID).Find(&memberRecord)
	if err != nil {
		return nil, err
	}

	return memberRecord, nil
}
