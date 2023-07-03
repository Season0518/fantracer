package components

import (
	"core/driver"
	"core/models"
	"core/services"
	"core/services/cqhttp"
	"strconv"
)

func FetchMemberList() error {
	route := "/get_group_member_list"
	fanGroups := [...]int64{700922190, 660717822, 671112420, 763084701, 669599441}

	for _, groupID := range fanGroups {
		rawData, err := cqhttp.FetchHttpData("", route, map[string]string{
			"group_id": strconv.FormatInt(groupID, 10),
			"no_cache": strconv.FormatBool(true),
		})
		if err != nil {
			return err
		}

		var members []models.MemberInfo
		err = cqhttp.SerializeRespData(rawData, &members)
		if err != nil {
			return err
		}

		for _, member := range members {
			err = services.UpsertDB(member, driver.Engine)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
