package main

import (
	"collector/components"
	"fmt"
	"log"
)
import "core/driver"

func main() {
	err := driver.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	err = components.FetchMemberListV2()
	if err != nil {
		fmt.Println(err)
	}
}
