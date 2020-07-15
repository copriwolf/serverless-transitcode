package service

import (
	"context"
	"encoding/json"
	"github.com/tencentyun/scf-go-lib/events"

	"internal/enumeration"
	"internal/structure"
	"log"
)

// Encode format content to encode string
func Encode(content string)  (resp structure.OutputContent){
	req := structure.InputContent{
		Opt: enumeration.Encode,
		Content: content,
	}
	return handler(req)
}

// Decode format content to decode string
func Decode(content string) (resp structure.OutputContent) {
	req := structure.InputContent{
		Opt: enumeration.Decode,
		Content: content,
	}
	return handler(req)
}

// Handler 加工入参、出参，并调用核心处理函数
func Handler(ctx context.Context, reqOrigin events.APIGatewayRequest) (resp events.APIGatewayResponse){

	// phase body json into struct
	req := structure.InputContent{}
	err := json.Unmarshal([]byte(reqOrigin.Body), &req)
	if err != nil {
		return handleErr(reqOrigin, err.Error())
	}

	// encode / decode handler
	respTransitCode := handler(req)

	// phase content into json
	body, err := json.Marshal(respTransitCode)
	if err != nil {
		return handleErr(reqOrigin, err.Error())
	}

	// structure the resp struct
	resp = events.APIGatewayResponse{
		IsBase64Encoded: false,
		StatusCode:      200,
		Headers:        map[string]string{"ContentType": "application/json"},
		Body:            string(body),
	}

	// log request & response
	log.Printf("[处理完毕] \n=> ctx \n%v \n=> req \n%v \n=> resp \n%v \n========== \n", ctx, req, resp)
	return
}

// handler 根据 opt 确定加/解密，并传入 content 字符串进行对应的编/解码
func handler(req structure.InputContent) structure.OutputContent {
	return structure.OutputContent{
		Origin: req.Content,
		Base64: Base64(req.Opt, req.Content),
		Html: Html(req.Opt, req.Content),
		Unicode: Unicode(req.Opt, req.Content),
		Url: Url(req.Opt, req.Content),
		Utf8ToGbk: Utf8ToGbk(req.Opt, req.Content),
	}
}

// handleErr 处理错误
func handleErr(reqOrigin events.APIGatewayRequest, errString string) (resp events.APIGatewayResponse) {
	// log
	log.Printf("[出现错误] \n//err %v \n//req %v \n========== \n", errString, reqOrigin)

	// handle error
	j, _ := json.Marshal(errString)
	resp = events.APIGatewayResponse{
		IsBase64Encoded: false,
		StatusCode:      500,
		Headers:        map[string]string{"ContentType": "application/json"},
		Body:           string(j),
	}
	return
}
