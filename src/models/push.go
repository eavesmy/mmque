package models

import (
	"encoding/binary"
)

type Task struct {
	Channal string
	Version int
	Msg     string
}

func UnpackPush(buf []byte) interface{} {

	index := 4
	d := &Task{}

	d.Version = int(binary.BigEndian.Uint16(buf[index : index+2]))
	index += 2

	channalLen := int(binary.BigEndian.Uint16(buf[index : index+2]))
	index += 2

	d.Channal = string(buf[index : index+channalLen])
	index += channalLen

	msgLen := int(binary.BigEndian.Uint16(buf[index : index+2]))
	index += 2

	d.Msg = string(buf[index : index+msgLen])

	var t interface{}
	t = d

	return t
}
