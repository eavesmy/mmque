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
	Id      int
	Version int
	Len     int
	Msg     string
}

var keyDictionary = []string{
	"Id",
	"Version",
	"Len",
	"Msg",
}

var BufferPool = make([]byte, 0)

func ReceiveBuffer(conn net.Conn) {

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

	id := binary.BigEndian.Uint16(realBuf[0:2])
	len := binary.BigEndian.Uint16(realBuf[2:4])

	if bufLen < len+4 {
		BufferPool = append(BufferPool, realBuf...)
		return
	}

	version := binary.BigEndian.Uint16(realBuf[4:6])
	msg := string(realBuf[6:])

	fmt.Println(id, version, len, msg) // -> 0

	/*
		_package := &Package{
			Id:      int16(realBuf[0:2]),
			Version: realBuf[2:4],
			Len:     realBuf[4:6],
			Msg:     realBuf[6:],
		}
	*/

	// fmt.Println(_package)

	// byteId := realBuf[0:BUF_INDEX]
	// fmt.Println(len(BufferPool), status, err)
}
