package main

import (
	"collector/components"
	"core/driver"
	"fmt"
	"log"
)

func main() {
	var err error
	err = driver.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	//err = driver.InitWS()
	//if err != nil {
	//	log.Fatal(err)
	//}

	err = components.FetchMemberList()
	err = components.FetchGroupInfo()
	if err != nil {
		fmt.Println(err)
	}
}
