package models

import (
	"encoding/binary"
)

type QueryOne struct {
	Channal string
	Version int
}

func UnpackQueryOne(_len int, buf []byte) interface{} {
	index := 4
	d := &QueryOne{}

	channalLen := int(binary.BigEndian.Uint16(buf[index : index+2]))
	index += 2

	d.Channal = string(buf[index : index+channalLen])
	index += channalLen

	d.Version = int(binary.BigEndian.Uint16(buf[index : index+2]))

	var t interface{}
	t = d

	return t
}
