package handler

import (
	"../models"
	"../state"
	"fmt"
	"net"
)

func Push(conn net.Conn, iData interface{}) {

	data := iData.(*models.Task)

	queue := state.CreateQueue(data.Channal)
	result := queue.Push(data)

	res := &models.Res{
		Msg:    result,
		Status: "200",
	}

	models.Send(conn, res)
}

func Pull(conn net.Conn, data *interface{}) {

}
