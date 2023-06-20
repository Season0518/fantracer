package main

import (
	"collector/components"
	"log"
)
import "core/driver"

func main() {
	err := driver.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	components.FetchMemberList()
}
