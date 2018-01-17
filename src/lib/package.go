package lib

import (
	// "bufio"
	// "bytes"
	"encoding/binary"
	"fmt"
	"net"
)

const BUF_INDEX = 2

type Package struct {
	Version int
	Len     int
	Msg     string
}

var BufferPool = make([]byte, 0)

func ReceiveBuffer(conn net.Conn) {

	index := 0

	if len(BufferPool) == 0 {

		// 检查BufferPool是否有数据。有则先拼接，再继续
	}

	tempBuf := make([]byte, 256)
	bufLen, _ := conn.Read(tempBuf)

	realBuf := tempBuf[0:bufLen]

	if bufLen < 4 {
		BufferPool = append(BufferPool, realBuf...)
		return
	} // 储存 buffer 碎片

	version := binary.BigEndian.Uint16(realBuf[index : index+2])
	index += 2
	len := binary.BigEndian.Uint16(realBuf[index : index+2])
	index += 2

	if bufLen < (int(len) + index) {
		BufferPool = append(BufferPool, realBuf...)
		return
	}

	msg := string(realBuf[index:])
}
