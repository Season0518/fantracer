package main

import (
	"fantracer/console"
)

func main(){
	console.FetchMemberList()
	defer console.Conrs.Stop()
}