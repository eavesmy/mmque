package lib

import (
	// "bufio"
	// "bytes"
	"encoding/binary"
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

	tempBuf := make([]byte, 256)
	bufLen, _ := conn.Read(tempBuf)

	realBuf := tempBuf[0:bufLen]

	if len(BufferPool) != 0 {

		// 检查BufferPool是否有数据。有则先拼接，清空pool,再继续
		// 拆len，如果len不够，扔回pool。

		realBuf = append(realBuf, BufferPool...)
		BufferPool = BufferPool[:0:0]
	}

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
