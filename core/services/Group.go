package services

import (
	"core/driver"
	"core/models"
	"log"
	"strings"
)

func InsertGroup(group models.Group) error {
	_, err := driver.Engine.Insert(group)

	return err
}

func UpsertGroup(group models.Group) error {
	_, err := driver.Engine.Insert(group)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			// 如果插入失败并且错误信息中包含 "Duplicate entry"，则尝试更新
			_, err = driver.Engine.Where("i_d = ?", group.ID).AllCols().Update(&group)
			if err != nil {
				log.Fatalf("2 Failed to update group: %v\n", err)
			}
		} else {
			// 如果插入失败并且错误信息中不包含 "Duplicate entry"，则记录错误
			log.Fatalf("1 Failed to insert group: %v\n", err)
		}
	}
	return err
}
