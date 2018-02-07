package models

import (
	"encoding/binary"
)

type RequestPull struct {
	Channal string
}

func UnpackRequestPull(buf []byte) interface{} {
	index := 4
	d := &RequestPull{}

	channalLen := int(binary.BigEndian.Uint16(buf[index : index+2]))
	index += 2

	d.Channal = string(buf[index : index+channalLen])

	var t interface{}
	t = d

	return t
}
