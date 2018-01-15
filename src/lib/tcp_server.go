package lib

import (
	"net"
)

// 消息格式 :
//  id
//	msg
//	version

func Server(port string) {
	tcpServer, tcpCreateErr := net.Listen("tcp", Port)

	if tcpCreateErr != nil {
		fmt.Println("!Server start err:")
		panic(tcpCreateErr)
	}

	fmt.Println("Mmque server started on port ", Port)

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
