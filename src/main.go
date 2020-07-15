package main

import (
	"fmt"
	"internal/enumeration"
	"internal/service"
)

func main() {
	// test encode
	fmt.Println(Encode("加密测试"))
	// test decode
	fmt.Println(Decode("\u52a0\u5bc6\u6d4b\u8bd5"))
}

// Encode format content to encode string
func Encode(content string)  (resp RespContent){
	req := RequestContent{
		Opt: enumeration.Encode,
		Content: content,
	}
	return handler(req)
}

// Decode format content to decode string
func Decode(content string) (resp RespContent) {
	req := RequestContent{
		Opt: enumeration.Decode,
		Content: content,
	}
	return handler(req)
}

type RequestContent struct {
	Opt 		int `json:"opt"`
	Content 	string `json:"content"`
}

type RespContent struct {
	Origin     	string `json:"origin"`
	Base64     	string `json:"base64"`
	Html       	string `json:"html"`
	Unicode    	string `json:"unicode"`
	Url  		string `json:"url"`
	Utf8ToGbk   string `json:"utf8ToGbk"`
}

// handler 根据 opt 确定加/解密，并传入 content 字符串进行对应的编/解码
func handler(req RequestContent) (resp RespContent){
	return RespContent{
		Origin: req.Content,
		Base64: service.Base64(req.Opt, req.Content),
		Html: service.Html(req.Opt, req.Content),
		Unicode: service.Unicode(req.Opt, req.Content),
		Url: service.Url(req.Opt, req.Content),
		Utf8ToGbk: service.Utf8ToGbk(req.Opt, req.Content),
	}
}

