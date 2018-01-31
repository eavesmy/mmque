package lib

import (
	// "bufio"
	// "bytes"
	"../models"
	"encoding/binary"
	"fmt"
	"net"
	// "strings"
)

const BUF_INDEX = 2

var BufferPool = make([]byte, 0)
var count = 0

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

	id := int(binary.BigEndian.Uint16(realBuf[index : index+2]))
	index += 2

	fmt.Println("Receive package id =>", id)

	_len := binary.BigEndian.Uint16(realBuf[index : index+2])
	index += 2

	if bufLen < (int(_len) + index) {
		BufferPool = append(BufferPool, realBuf...)
		return
	}

	handler := Routes[id]

	if handler == nil {
		return
	}

	data := Parse(id, int(_len), realBuf)
	handler(conn, data)
}

func Parse(id int, length int, buf []byte) interface{} {
	var data interface{}

	switch id {
	case 1:
		data = models.UnpackPush(length, buf)
	case 2:
		data = models.UnpackPull(length, buf)
	}

	return data
}
