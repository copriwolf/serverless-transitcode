package service

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"internal/enumeration"
	"io/ioutil"
)

// Utf8ToGbk 根据 opt 运算符确定是中文转 UTF8，还是 UTF8 转中文
func Utf8ToGbk(opt int, content string) string {
	switch opt {
	case enumeration.Encode:  // Encode - UTF8 => GBK
		return encode(content)
	case enumeration.Decode:  // Decode - GBK => UTF8
		return decode(content)
	default:
		return ""
	}
}

// encode format UTF-8 => GBK
func encode(content string) string {
	ret := ""
	if len(content) <= 0 {
		return ret
	}
	readers := transform.NewReader(bytes.NewReader([]byte(content)), simplifiedchinese.GBK.NewEncoder())
	b, err := ioutil.ReadAll(readers)
	if err != nil {
		return ret
	}
	ret = string(b)
	return ret
}

// decode format GBK => UTF-8
func decode(content string) string {
	ret := ""
	if len(content) <= 0 {
		return ret
	}
	readers := transform.NewReader(bytes.NewReader([]byte(content)), simplifiedchinese.GBK.NewDecoder())
	b, err := ioutil.ReadAll(readers)
	if err != nil {
		return ret
	}
	ret = string(b)
	return ret
}
