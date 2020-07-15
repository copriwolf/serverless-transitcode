package service

import (
	"internal/enumeration"
	"internal/eventStruct"
	"log"
)

// Encode format content to encode string
func Encode(content string)  (resp eventStruct.RespContent){
	req := eventStruct.RequestContent{
		Opt: enumeration.Encode,
		Content: content,
	}
	return Handler(req)
}

// Decode format content to decode string
func Decode(content string) (resp eventStruct.RespContent) {
	req := eventStruct.RequestContent{
		Opt: enumeration.Decode,
		Content: content,
	}
	return Handler(req)
}

// handler 根据 opt 确定加/解密，并传入 content 字符串进行对应的编/解码
func Handler(req eventStruct.RequestContent) (resp eventStruct.RespContent){
	resp = eventStruct.RespContent{
		Origin: req.Content,
		Base64: Base64(req.Opt, req.Content),
		Html: Html(req.Opt, req.Content),
		Unicode: Unicode(req.Opt, req.Content),
		Url: Url(req.Opt, req.Content),
		Utf8ToGbk: Utf8ToGbk(req.Opt, req.Content),
	}
	log.Println(req, resp)
	return
}

