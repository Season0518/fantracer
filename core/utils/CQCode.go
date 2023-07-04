package utils

import (
	"core/models"
	"strings"
)

// EscapeText 将字符串raw中部分字符转义
//
//   - & -> &amp;
//   - [ -> &#91;
//   - ] -> &#93;
func EscapeText(s string) string {
	count := strings.Count(s, "&")
	count += strings.Count(s, "[")
	count += strings.Count(s, "]")
	if count == 0 {
		return s
	}

	// Apply replacements to buffer.
	var b strings.Builder
	b.Grow(len(s) + count*4)
	start := 0
	for i := 0; i < count; i++ {
		j := start
		for index, r := range s[start:] {
			if r == '&' || r == '[' || r == ']' {
				j += index
				break
			}
		}
		b.WriteString(s[start:j])
		switch s[j] {
		case '&':
			b.WriteString("&amp;")
		case '[':
			b.WriteString("&#91;")
		case ']':
			b.WriteString("&#93;")
		}
		start = j + 1
	}
	b.WriteString(s[start:])
	return b.String()
}

// EscapeValue 将字符串value中部分字符转义
//
//   - , -> &#44;
//   - & -> &amp;
//   - [ -> &#91;
//   - ] -> &#93;
func EscapeValue(value string) string {
	ret := EscapeText(value)
	return strings.ReplaceAll(ret, ",", "&#44;")
}

func SerializeCQCode(messageChain []models.MessageBody) string {
	var sb strings.Builder
	for _, code := range messageChain {
		if code.Type == "text" {
			sb.WriteString(EscapeText(code.Data["text"]))
			continue
		}
		sb.WriteString("[CQ:")
		sb.WriteString(code.Type)
		for k, v := range code.Data {
			sb.WriteString(",")
			sb.WriteString(k)
			sb.WriteString("=")
			sb.WriteString(EscapeValue(v))
		}
		sb.WriteString("]")
	}
	return sb.String()
}
