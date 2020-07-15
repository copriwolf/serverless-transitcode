package service

import (
	"internal/enumeration"
	"strconv"
)

// Url 根据 opt 运算符确定对 content 进行编码/解码
func Unicode(opt int, content string) string {
	switch opt {
	case enumeration.Encode:
		return unicodeEncode(content)
	case enumeration.Decode:
		return unicodeDecode(content)
	default:
		return ""
	}
}

// unicodeEncode use content to url encode
func unicodeEncode(content string) string {
	if len(content) <= 0 {
		return ""
	}
	ret := ""
	rc := []rune(content)
	for _,r := range rc {
		rint := int(r)
		if rint < 128 {
			ret += string(r)
		} else {
			ret += "\\u"+strconv.FormatInt(int64(rint), 16)
		}
	}
	return ret
}

// unicodeDecode use content to url decode
func unicodeDecode(content string) string {
	if len(content) <= 0 {
		return ""
	}
	return content
}
