package utils

import (
	"core/models"
	"strings"
)

func SerializeCQCode(messageChain []models.MessageBody) string {
	var sb strings.Builder
	for _, code := range messageChain {
		sb.WriteString("[CQ:")
		sb.WriteString(code.Type)
		for k, v := range code.Data {
			sb.WriteString(",")
			sb.WriteString(k)
			sb.WriteString("=")
			sb.WriteString(v)
		}
		sb.WriteString("]")
	}
	return sb.String()
}
