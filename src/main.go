package main

import (
	"context"
	"fmt"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"github.com/tencentyun/scf-go-lib/events"
	"log"

	"internal/service"
)


// main display a demo with serverless
// main 展示的是一个使用了 serverless 的转码程序
func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(server)
}

func server(ctx context.Context, req events.APIGatewayRequest) (resp events.APIGatewayResponse, err error) {
	log.Printf("[请求接收] \n=> ctx \n%v \n=> req \n%v \n========== \n", ctx, req)
	return service.Handler(ctx, req), nil
}

// =====================================================================
// mainOrigin display a demo without serverless
// mainOrigin 展示的是一个不使用 serverless 的转码程序
func mainOrigin() {
	// test encode
	fmt.Println(service.Encode("加密测试"))
	// test decode
	fmt.Println(service.Decode("\u52a0\u5bc6\u6d4b\u8bd5"))
}
