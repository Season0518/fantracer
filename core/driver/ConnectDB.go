package driver

import (
	"core/models"
	"core/utils"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func InitDB() error {
	var err error

	account, password, port, err := utils.ReadMySQLConfig()
	if err != nil {
		return err
	}

	connection := fmt.Sprintf("%s:%s@(localhost:%s)/fantracer?charset=utf8mb4", account, password, port)
	fmt.Println(connection)
	Engine, err = xorm.NewEngine("mysql", connection)
	if err != nil {
		//log.Fatalf("Fail to create engine: %v\n", err)
		return err
	}

	// you can adjust the configurations according to your needs
	// Engine.ShowSQL(true)
	// Engine.SetMaxOpenConns(10)
	// Engine.SetMaxIdleConns(5)

	if err = Engine.Ping(); err != nil {
		//log.Fatalf("Fail to ping database: %v\n", err)
		return err
	}

	// 创建多个表
	modelsToSync := []interface{}{
		new(models.MemberInfo),
		new(models.GroupInfo),
		new(models.PostRecord),
		new(models.PostInfo),
	}

	for _, model := range modelsToSync {
		exist, err := Engine.IsTableExist(model)
		if err != nil {
			//log.Printf("Error while checking table existence: %v\n", err)
			return err
		}

		if !exist {
			log.Printf("表单不存在，正在同步数据.")
			err = Engine.Sync2(model)
			if err != nil {
				//log.Fatalf("Fail to sync database: %v\n", err)
				return err
			}
		}
	}

	return nil
}
