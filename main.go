package main

import (
	"fantracer/console"
	"fantracer/respository"
	"fmt"
)

func main(){
	console.FetchMemberList()
	memberRecord,err := respository.FindMemberInGroups(290975396)
	if err != nil { 
		fmt.Println(err)
	} else {
		fmt.Println(memberRecord)
	}
	defer console.Conrs.Stop()
}