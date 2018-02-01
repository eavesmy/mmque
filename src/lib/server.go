package lib

import (
	"fmt"
	"net"
	"time"
)

// 消息格式 :

//  id   <- 2
//  len  <- 2
//	version <- 2
//	msg  <- 2

var Pool []net.Conn

func init() {
	KeepAlive()
}

func Server(port string) {
	tcpServer, tcpCreateErr := net.Listen("tcp", port)

	if tcpCreateErr != nil {
		fmt.Println("!Server start err:")
		panic(tcpCreateErr)
	}

	fmt.Println("Mmque server started on port ", port)

	for {
		conn, err := tcpServer.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		Pool = append(Pool, conn)

		go ReceiveBuffer(conn)
	}
}

func KeepAlive() {

	for i, conn := range Pool {
		_, err := conn.Write([]byte{1})

		if err != nil {
			conn.Close()
			Pool = append(Pool[:i], Pool[i+1:]...)
		}
	}

	time.AfterFunc(time.Second*10, KeepAlive)
}
