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
	var resp models.SinaWeiboResp
	var err error
	maxAttempt := 3
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < maxAttempt; i++ {
		resp, err = weibo.GetLatestBlog(uid, 1)
		if err == nil && resp.Ok == 1 {
			break
		}

		log.Printf("在%d未能成功对%d执行Update操作,重试次数: %d", time.Now().UnixNano(), uid, i)
		time.Sleep(time.Duration(600+r.Intn(600)) * time.Second)
	}
	if err != nil {
		return err
	}

	if resp.Ok != 1 {
		return fmt.Errorf("获取微博更新状态发生错误: 状态码异常,重试到达上限,请重启服务")
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

	// 如果PostInfo中没有记录某UID最后一次更新的时间戳，则将数据库中该UID的最后一条记录作为最后一次更新的时间戳
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

		err = cqhttp.PostMessageSendEvent(660717822, updateMessage)
		if err != nil {
			return err
		}
		return nil
	}
}

func UpdateExceptionHandler(exception error) error {
	failedMessage, _ := BuildFailedMessage(exception)
	err1 := cqhttp.PostMessageSendEvent(865444787, failedMessage)
	if err1 != nil {
		return err1
	}
	return nil
}

func SendUpdateMessage() error {
	platform := []struct {
		FetchEvent func(int64) error
		UIDs       []int64
		Platform   string
	}{
		{FetchLatestWeibo, []int64{6620766403, 6593497650}, "微博"},
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		for _, uf := range platform {
			for _, uid := range uf.UIDs {
				//执行不同平台的更新操作。
				err := uf.FetchEvent(uid)

				if err != nil {
					_ = UpdateExceptionHandler(err)
					return err
				}

				err = UpdatePostInfo(uid, uf.Platform)
				if err != nil {
					_ = UpdateExceptionHandler(err)
					return err
				}

				log.Printf("%s 在%d成功对%d执行操作: Update", uf.Platform, time.Now().UnixNano(), uid)
				time.Sleep(time.Duration(5+r.Intn(15)) * time.Second)
			}
			time.Sleep(time.Duration(35+r.Intn(120)) * time.Second)
		}
	}
}
