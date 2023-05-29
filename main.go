package main

import (
	"encoding/json"
	"fantracer/models"
	"fmt"
)

func main(){
	fmt.Print("hello World")
	jsonStr := `{
        "id": 3574118153,
        "memberName": "八尾傲世",
        "specialTitle": "",
        "permission": "MEMBER",
        "joinTimestamp": 1684841547,
        "lastSpeakTimestamp": 1684841547,
        "muteTimeRemaining": 0,
        "group": {
            "id": 671112420,
            "name": "冰灵的土匪老窝",
            "permission": "MEMBER"
        }
    }`

    var member models.Member
    err := json.Unmarshal([]byte(jsonStr), &member)
    if err != nil {
        fmt.Println("JSON unmarshaling failed:", err)
        return
    }

    fmt.Println("Member ID:", member.ID)
    fmt.Println("Member Name:", member.MemberName)
    fmt.Println("Group ID:", member.Group.ID)
    fmt.Println("Group Name:", member.Group.Name)
}
