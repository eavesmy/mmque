package lib

import (
	"fmt"
	"net"
)

// 消息格式 :
//  id
//	msg
//	version

func Server(port string) {
	tcpServer, tcpCreateErr := net.Listen("tcp", port)

	if tcpCreateErr != nil {
		fmt.Println("!Server start err:")
		panic(tcpCreateErr)
	}

	fmt.Println("Mmque server started on port ", port)

	defer tcpServer.Close()

	for {
		conn, err := tcpServer.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		go ReceiveBuffer(conn)
	}

}
