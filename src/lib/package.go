package lib

import (
	// "bufio"
	// "bytes"
	"../models"
	"encoding/binary"
	"net"
	// "strings"
)

const BUF_INDEX = 2

var BufferPool = make([]byte, 0)
var count = 0

func ReceiveBuffer(conn net.Conn) {

	for {
		index := 0

		tempBuf := make([]byte, 256)
		bufLen, connectErr := conn.Read(tempBuf)

		if connectErr != nil {

			if connectErr.Error() != "EOF" {
				conn.Close() // 连接被 client 端关闭
				break
			}

			continue
		}

		realBuf := tempBuf[0:bufLen]

		if len(BufferPool) != 0 {

			// 检查BufferPool是否有数据。有则先拼接，清空pool,再继续
			// 拆len，如果len不够，扔回pool。

			realBuf = append(realBuf, BufferPool...)
			BufferPool = BufferPool[:0:0]
		}

		// fmt.Println(bufLen, "receive buf len", string(realBuf))

		if bufLen < 4 {
			BufferPool = append(BufferPool, realBuf...)
			continue
		} // 储存 buffer 碎片

		id := int(binary.BigEndian.Uint16(realBuf[index : index+2]))
		index += 2

		// fmt.Println("Receive package id =>", id)

		_len := binary.BigEndian.Uint16(realBuf[index : index+2])
		index += 2

		if bufLen < (int(_len) + index) {
			BufferPool = append(BufferPool, realBuf...)
			continue
		}

		handler := Routes[id]

		if handler == nil {
			continue
		}

		data := Parse(id, realBuf)

		go handler(conn, data)
	}
}

func Parse(id int, buf []byte) interface{} {
	var data interface{}

	switch id {
	case 1:
		data = models.UnpackPush(buf)
	case 2:
		data = models.UnpackQueryOne(buf)
	case 3:
		data = models.UnpackQueryOne(buf)
	case 4:
		data = models.UnpackVersion(buf)
	case 5:
		data = models.UnpackRequestPull(buf)
	}

	return data
}
