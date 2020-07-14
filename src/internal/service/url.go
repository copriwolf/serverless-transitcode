package service

import (
	"net/url"
)

// Url 根据 opt 运算符确定对 content 进行编码/解码
func Url(opt int, content string) string {
	switch opt {
	case 1: //todo enum // Encode
		return urlEncode(content)
	case 2: //todo enum // Decode
		return urlDecode(content)
	default:
		return ""
	}
}

// urlEncode use content to url encode
func urlEncode(content string) string {
	if len(content) <= 0 {
		return ""
	}
	return url.QueryEscape(content)
}

// urlDecode use content to url decode
func urlDecode(content string) string {
	ret := ""

	if len(content) <= 0 {
		return ret
	}
	ret, _ = url.QueryUnescape(content)
	return ret

}
