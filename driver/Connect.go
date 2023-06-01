package driver

import (
	"fantracer/models"
	"fantracer/utils"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func init() {
    var err error

    account,password,port,err := utils.ReadMySQLConfig()
    if err != nil {
        log.Fatalf("读取MySQL配置失败!")
    }

    connection := fmt.Sprintf("%s:%s@(localhost:%s)/fantracer?charset=utf8mb4",account,password,port)
    fmt.Println(connection)
    Engine, err = xorm.NewEngine("mysql", connection)
    if err != nil {
        log.Fatalf("Fail to create engine: %v\n", err)
    }

    // you can adjust the configurations according to your needs
    // Engine.ShowSQL(true)
    // Engine.SetMaxOpenConns(10)
    // Engine.SetMaxIdleConns(5)

    if err = Engine.Ping(); err != nil {
        log.Fatalf("Fail to ping database: %v\n", err)
    }
    
    // 创建多个表
    modelsToSync := []interface{}{
        new(models.Group),
        new(models.Member),
    }

    for _,model := range modelsToSync {
        err = Engine.Sync2(model)
        if err != nil {
            log.Fatalf("Fail to sync database: %v\n", err)
        }
    }
}