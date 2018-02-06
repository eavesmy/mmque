package models

import (
	"encoding/binary"
)

type RequestVersion struct {
	Channal string
}

func UnpackVersion(_len int, buf []byte) interface{} {
	index := 4
	d := &RequestVersion{}

	channalLen := int(binary.BigEndian.Uint16(buf[index : index+2]))
	index += 2

	d.Channal = string(buf[index : index+channalLen])

	var t interface{}
	t = d

	return t
}
