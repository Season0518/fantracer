package components

import (
	"core/driver"
	"core/services"
	"core/services/cqhttp"
)

func FetchGroupInfo() error {
	fanGroups := [...]int64{700922190, 660717822, 671112420, 763084701, 669599441}

	for _, groupID := range fanGroups {
		groupInfo, err := cqhttp.GetGroupInfo(groupID, true)
		if err != nil {
			return err
		}

		err = services.UpsertDB(groupInfo, driver.Engine)
		if err != nil {
			return err
		}
	}

	return nil
}

func FetchMemberList() error {
	fanGroups := [...]int64{700922190, 660717822, 671112420, 763084701, 669599441}

	for _, groupID := range fanGroups {
		groupMembers, err := cqhttp.GetMemberList(groupID, true)
		if err != nil {
			return err
		}

		for _, member := range groupMembers {
			err = services.UpsertDB(member, driver.Engine)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
