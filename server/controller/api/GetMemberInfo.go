package api

import (
	"core/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMemberInfo(c *gin.Context) {
	memberID, err := strconv.ParseInt(c.Query("memberID"), 10, 64)

	if err != nil {
		fmt.Println("QQ号异常")
	}
	records, err := services.FindMemberInGroups(memberID)

	if err != nil {
		fmt.Println("在查询时发生异常")
	}

	for _, record := range records {
		fmt.Println(record)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "errormsg",
		"data": records,
	})
}
