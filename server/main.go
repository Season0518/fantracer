package main

import (
	"core/driver"
	"log"
	"server/router"
)

func main() {
	err := driver.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	router := router.StartRouter()
	router.Run(":8082")
}
