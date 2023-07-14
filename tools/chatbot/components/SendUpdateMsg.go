package components

import (
	"core/driver"
	"core/models"
	"core/services"
	"core/services/cqhttp"
	"core/services/weibo"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func FetchLatestWeibo(uid int64) error {
	resp, err := weibo.GetLatestBlog(uid, 1)
	if err != nil {
		return err
	}

	if resp.Ok != 1 {
		return fmt.Errorf("不能获取微博更新状态,响应:%s", resp.Data)
	}

	for _, blog := range resp.Data.Cards {
		err = services.UpsertDB(models.PostRecord{
			UserID:    uid,
			TimeStamp: blog.Timestamp,
			Refer:     blog.Scheme,
			Text:      blog.Mblog.Text,
		}, driver.Engine)

		if err != nil {
			return err
		}
	}

	return nil
}

func UpdatePostInfo(userID int64, platform string) error {
	var err error
	var updateRecord models.PostRecord //更新的数据
	var updateInfo models.PostInfo     //最后一次更新的记录

	has, err := driver.Engine.Where("user_id = ?", userID).Get(&updateInfo)
	if err != nil {
		return err
	}

	has, err = driver.Engine.Where("user_id = ?", userID).Desc("time_stamp").Limit(1).Get(&updateRecord)
	if err != nil || !has {
		return err
	}

	// 如果不存在更新记录，则最后一次发布记录为更新记录。
	if updateInfo.UserID == 0 {
		err = services.InsertDB(models.PostInfo{
			UserID:    userID,
			TimeStamp: updateRecord.TimeStamp,
			Platform:  platform,
		}, driver.Engine)
		if err != nil {
			return err
		}
		has, err = driver.Engine.Where("user_id = ?", userID).Get(&updateInfo)
		if err != nil {
			return err
		}
	}

	if updateInfo.TimeStamp == updateRecord.TimeStamp {
		return nil
	} else {
		updateInfo.TimeStamp = updateRecord.TimeStamp

		err = services.UpdateDB(updateInfo, driver.Engine)
		if err != nil {
			return err
		}

		updateMessage, err := BuildUpdateMessage(updateInfo, updateRecord)
		if err != nil {
			return err
		}

		err = cqhttp.PostMessageSendEvent(865444787, updateMessage)
		if err != nil {
			return err
		}
		return nil
	}
}

func SendUpdateMessage() error {
	funcs := []struct {
		Func     func(int64) error
		UIDs     []int64
		Platform string
	}{
		{FetchLatestWeibo, []int64{6620766403, 6593497650}, "微博"},
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		for _, uf := range funcs {
			// 所有错误预期打印到测试群中，暂未实现。
			for _, uid := range uf.UIDs {
				err := uf.Func(uid)
				if err != nil {
					log.Printf("Error updating for UIDs %v: %v\n", uf.UIDs, err)
					return err
				}
				log.Println("Updated: ", uid)

				err = UpdatePostInfo(uid, uf.Platform)
				if err != nil {
					return err
				}
				time.Sleep(time.Duration(5+r.Intn(5)) * time.Second)

			}
			time.Sleep(time.Duration(10+r.Intn(30)) * time.Second)
		}
	}
}
