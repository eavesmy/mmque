package main

import (
	"./lib"
	"flag"
	"fmt"
	"net"
)

var Port string

func init() {

	port := flag.String("p", "8081", "Server start use whitch port.")
	flag.Parse()

	Port = ":" + *port
}

// 先写一套 tcp 的.
func main() {

	// 运行顺序：
	// 查询本地消息文件，如有未完成的消息则加入队列。
	// 启动服务。

	lib.CheckLocalDir()

}