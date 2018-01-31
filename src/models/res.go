package models

import (
	"encoding/json"
	"net"
)

type Res struct {
	Msg    string `json:Msg`
	Status string `json:Status`
}

func Send(conn net.Conn, data *Res) {

	info, _ := json.Marshal(data)
	conn.Write(info)
}
