package components

import (
	"collector/middleware"
	"core/driver"
	"core/models"
	"core/services"
	"core/utils"
	"log"
	"time"
)

func FetchMemberList() {
	fanGroups := [...]int64{700922190, 660717822, 671112420, 763084701, 669599441}

	// 获取SessionKey
	verifyKey, _ := utils.ReadVerifyKey()
	botAccount, err := utils.ReadBotAccount()

	if err != nil {
		log.Printf("读取BotQQ时发生错误: %v", err)
		return
	}

	sessionKey, err := middleware.VerifySession(verifyKey, botAccount)
	if err != nil {
		log.Printf("在认证时发生错误: %v", err)
		return
	}
	log.Printf("sessionKey: %v\n", sessionKey)

	// 对SessionKey进行绑定
	_, err = middleware.BindSession(sessionKey, botAccount)
	if err != nil {
		log.Printf("在绑定时发生错误: %v", err)
		return
	}

	// 获取群聊成员列表
	for _, group := range fanGroups {
		members, err := middleware.FetchGroupMembers(sessionKey, group, 0)
		if err != nil {
			log.Printf("在查询时发生错误: %v", err)
		}

		var groupInfo models.GroupInfo

		groupInfo.GroupID = group
		groupInfo.MemberCount = len(members)
		groupInfo.TimeStamp = time.Now().Unix()
		err = services.InsertGroupInfo(groupInfo)
		if err != nil {
			log.Println("插入群聊信息时出错！")
		}

		if len(members) != 0 {
			services.UpsertGroup(members[0].Group)
			for _, member := range members {
				member.GroupID = member.Group.ID
				services.UpsertMember(member)
			}
		} else {
			log.Printf("群聊 %d 可能不存在\n", group)
		}
	}

	// 释放SessionKey，防止盗用
	_, err = middleware.ReleaseSession(sessionKey, botAccount)
	if err != nil {
		log.Printf("释放时发生错误: %v", err)
		return
	}
}

func FetchMemberListV2() error {
	fanGroups := [...]int64{700922190, 660717822, 671112420, 763084701, 669599441}

	for _, groupID := range fanGroups {
		members, err := middleware.FetchGroupMembersV2("", groupID)
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
