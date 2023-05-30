package main

import "fantracer/console"

func main(){
	// console.
	// test.FetchFlow()
	console.FetchMemberList()
	defer console.Conrs.Stop()
	// return
}