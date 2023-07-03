package components

import (
	"core/driver"
	"core/models"
	"core/services"
	"core/services/cqhttp"
	"strconv"
	"time"
)

func FetchGroupInfo() error {
	route := "/get_group_info"
	currentTime := time.Now().Unix()
	fanGroups := [...]int64{700922190, 660717822, 671112420, 763084701, 669599441}

	for _, groupID := range fanGroups {
		rawData, err := cqhttp.FetchHttpData("", route, map[string]string{
			"group_id": strconv.FormatInt(groupID, 10),
			"no_cache": strconv.FormatBool(true),
		})
		if err != nil {
			return err
		}

		var groupInfo models.GroupInfo
		err = cqhttp.SerializeRespData(rawData, &groupInfo)
		if err != nil {
			return err
		}

		groupInfo.InfoRetrievedAt = currentTime
		err = services.UpsertDB(groupInfo, driver.Engine)
		if err != nil {
			return err
		}
	}

	return nil
}
