package main

import (
	"fmt"
	"internal/service"
)

func main() {
	// test encode
	fmt.Println(service.Encode("加密测试"))
	// test decode
	fmt.Println(service.Decode("\u52a0\u5bc6\u6d4b\u8bd5"))
}


