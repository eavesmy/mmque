package lib

import (
	"fmt"
	"net"
)

var BufferPool []byte

func ReceiveBuffer(conn net.Conn) {
	fmt.Println(len(BufferPool))
}
