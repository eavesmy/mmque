package handler

import (
	"../models"
	"../state"
	// "encoding/json"
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

	res.Send(conn)
}

func QueryOne(conn net.Conn, iData interface{}) {

	data := iData.(*models.QueryOne)

	queue := state.Pool[data.Channal]

	res := &models.Res{}

	if queue == nil {

		res.Msg = "No this queue!"
		res.Status = "404"

		models.Send(conn, res)
		return
	}

	task := queue.Pull(data.Version)

	if task == nil {

		res.Msg = "No this task!"
		res.Status = "404"

		models.Send(conn, res)
		return
	}

	res.Msg = task.Msg
	res.Status = "200"

	res.Send(conn)
}

func Pull(conn net.Conn, iData interface{}) {
	data := iData.(*models.RequestPull)

	queue := state.Pool[data.Channal]

	res := &models.Res{}

	if queue == nil {
		res.Msg = "No this queue!"
		res.Status = "404"

		res.Send(conn)
		return
	}

	if len(queue.List) == 0 {
		res.Msg = "Empty queue."
		res.Status = "400"

		res.Send(conn)
		return
	}

	task := queue.List[0]

	res.Msg = fmt.Sprintf("%d", task.Version) + "|" + task.Msg
	res.Status = "200"

	res.Send(conn)
}

func Ack(conn net.Conn, iData interface{}) {

	data := iData.(*models.QueryOne)

	queue := state.Pool[data.Channal]

	res := &models.Res{}

	if queue == nil {

		res.Msg = "No this queue!"
		res.Status = "404"

		models.Send(conn, res)
		return
	}

	done := queue.Ack(data.Version)

	if done {
		res.Msg = "Done"
		res.Status = "200"

		models.Send(conn, res)
		return

	} else {
		res.Msg = "Not found this task."
		res.Status = "404"

		models.Send(conn, res)
		return

	}
}

func NewVersion(conn net.Conn, iData interface{}) {
	data := iData.(*models.RequestVersion)

	queue := state.Pool[data.Channal]

	res := &models.Res{}

	if queue == nil {

		res.Status = "200"
		res.Msg = "0"

		res.Send(conn)

		return
	}

	version := queue.NewVersion()

	res.Status = "200"
	res.Msg = fmt.Sprintf("%d", version)
	res.Send(conn)
}
