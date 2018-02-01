package models

import (
	"encoding/binary"
)

type PullRequest struct {
	Channal string
	Version int
}

func UnpackPull(_len int, buf []byte) interface{} {
	index := 4
	d := &PullRequest{}

	channalLen := int(binary.BigEndian.Uint16(buf[index : index+2]))
	index += 2

	d.Channal = string(buf[index : index+channalLen])
	index += channalLen

	d.Version = int(binary.BigEndian.Uint16(buf[index : index+2]))

	var t interface{}
	t = d

	return t
}
