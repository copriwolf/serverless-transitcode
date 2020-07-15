package service

import (
	"html"
	"internal/enumeration"
)

// Html 根据 opt 运算符确定对 content 进行编/解码
func Html(opt int, content string) string {
	switch opt {
	case enumeration.Encode:
		return htmlEncode(content)
	case enumeration.Decode:
		return htmlDecode(content)
	default:
		return ""
	}
}

// htmlEncode use content to url encode
func htmlEncode(content string) string {
	if len(content) <= 0 {
		return ""
	}
	return html.EscapeString(content)
}

// htmlDecode use content to url decode
func htmlDecode(content string) string {
	ret := ""

	if len(content) <= 0 {
		return ret
	}
	return html.UnescapeString(content)
}
