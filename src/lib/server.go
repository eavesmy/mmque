package lib

import (
	"fmt"
	"net"
)

// 消息格式 :

//  id   <- 2
//  len  <- 2
//	version <- 2
//	msg  <- 2

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

		go ReceiveBuffer(conn)
	}
}
