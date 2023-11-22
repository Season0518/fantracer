package main

import (
	"collector/components"
	"core/driver"
	"fmt"
	"log"
)

var err error

func main() {
	err = driver.InitCfg()
	if err != nil {
		log.Fatal(err)
	}

	err = driver.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	err = driver.InitWS()
	if err != nil {
		log.Fatal(err)
	}

	err = components.FetchMemberList()
	err = components.FetchGroupInfo()
	if err != nil {
		fmt.Println(err)
	}
}
