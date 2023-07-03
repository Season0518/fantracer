package main

import (
	"collector/components"
	"core/driver"
	"fmt"
	"log"
)

func main() {
	err := driver.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	err = components.FetchMemberList()
	err = components.FetchGroupInfo()
	if err != nil {
		fmt.Println(err)
	}

	//var q []models.MemberInfo
	//err = services.QueryDB(fmt.Sprintf("user_id = %v", 1342367762), &q, driver.Engine)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(q)

}
