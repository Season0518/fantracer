package douyin

import (
	"fmt"
	"testing"
)

func TestGetPosts(t *testing.T) {
	var user UserInfo
	err := user.GetPosts("MS4wLjABAAAAu-A0s9aIathifzLqcPvwBvMaOIA5XGicTEU8wc1dilk")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(user.Posts.AwemeList[0].Desc)
}

func TestGetProfile(t *testing.T) {
	var user UserInfo
	err := user.GetProfile("MS4wLjABAAAAu-A0s9aIathifzLqcPvwBvMaOIA5XGicTEU8wc1dilk")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(user.Profile.User.Nickname)
}
