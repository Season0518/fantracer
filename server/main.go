package main

import (
	"server/router"
)

func main() {
	router := router.StartRouter()
	router.Run(":8082")
}
