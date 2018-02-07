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

var Pool = make([]net.Conn, 0)

func init() {
	go KeepAlive()
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
			break
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

			if len(Pool) == 1 {
				Pool = make([]net.Conn, 0)
			} else {
				Pool = append(Pool[:i], Pool[i+1:]...)
			}
			conn = nil
		}
	}

	time.AfterFunc(time.Second*10, KeepAlive)
}
