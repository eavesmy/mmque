package models

import (
	"encoding/json"
	"net"
)

type Res struct {
	ID     string `json:ID`
	Msg    string `json:Msg`
	Status string `json:Status`
}

func (res *Res) Send(conn net.Conn) {
	info, _ := json.Marshal(res)
	conn.Write(info)
}

func Send(conn net.Conn, data *Res) {

	info, _ := json.Marshal(data)
	conn.Write(info)
}
