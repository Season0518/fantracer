package console

import (
	"fantracer/middleware"
	"fantracer/utils"
	"log"

	"fantracer/respository"

	"github.com/robfig/cron/v3"
)

func FetchMemberList() {
	fanGroups := [...] int64 { 700922190,660717822,671112420,763084701,669599441 }

	// 获取SessionKey
	verifyKey,_ := utils.ReadVerifyKey()
	botAccount,err := utils.ReadBotAccount()

	if err != nil {
		log.Printf("读取BotQQ时发生错误: %v",err)
		return
	}

	sessionKey,err := middleware.VerifySession(verifyKey,botAccount)
	if err != nil {
		log.Printf("在认证时发生错误: %v",err)
		return
	}
	log.Printf("sessionKey: %v\n",sessionKey)

	// 对SessionKey进行绑定
	_,err = middleware.BindSession(sessionKey,botAccount)
	if err != nil {
		log.Printf("在绑定时发生错误: %v",err)
		return
	}

	// 获取群聊成员列表
	for _,group := range fanGroups {
		members,err := middleware.FetchGroupMembers(sessionKey,group,0)
		if err != nil {
			log.Printf("在查询时发生错误: %v",err)
		}
		if len(members) != 0 {
			respository.InsertGroup(members[0].Group)
			for _,member := range members {
				member.GroupID = member.Group.ID
				respository.InsertMember(member)
			}
		} else {
			log.Printf("群聊 %d 可能不存在\n",group)
		}
	}

	// 释放SessionKey，防止盗用
	_,err = middleware.ReleaseSession(sessionKey,botAccount)
	if err != nil {
		log.Printf("释放时发生错误: %v",err)
		return

	}
}

var Conrs *cron.Cron

func init(){
	Conrs = cron.New() // 定时任务
	Conrs.Start()
	// 定时获取群聊人数
	_, err := Conrs.AddFunc("@every 1m", FetchMemberList)
	if err != nil {
		return
	}
}