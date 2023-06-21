package components

import (
	"collector/middleware"
	"core/driver"
	"core/services"
)

func FetchMemberList() error {
	fanGroups := [...]int64{700922190, 660717822, 671112420, 763084701, 669599441}

	for _, groupID := range fanGroups {
		members, err := middleware.FetchGroupMembers("", groupID)
		if err != nil {
			return err
		}

		for _, member := range members {
			err = services.Upsert(member, driver.Engine)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
