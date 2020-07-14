package service

import "encoding/base64"

// Base64 根据 opt 运算符确定对 content 进行加解密
func Base64(opt int, content string) string {
	switch opt {
	case 1: //todo enum // Encode
		return base64Encode(content)
	case 2: //todo enum // Decode
		return base64Decode(content)
	default:
		return ""
	}
}

// base64Encode use content to base64 encode
func base64Encode(content string) string {
	ret := ""
	if len(content) <= 0 {
		return ret
	}
	b := []byte(content)
	ret = base64.StdEncoding.EncodeToString(b)
	return ret
}

// base64Decode use content to base64 decode
func base64Decode(content string) string {
	ret := ""

	if len(content) <= 0 {
		return ret
	}
	b, _ := base64.StdEncoding.DecodeString(content)
	ret = string(b)
	return ret

}
