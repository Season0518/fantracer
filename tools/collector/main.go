package main

import (
	"core/models"
	"core/services"
	"fmt"
	"log"
)
import "core/driver"

func main() {
	err := driver.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	//err = components.FetchMemberList()
	//if err != nil {
	//	fmt.Println(err)
	//}

	var q []models.MemberInfo
	err = services.Query(fmt.Sprintf("user_id = %v", 1342367762), &q, driver.Engine)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(q)

}
