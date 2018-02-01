package handler

import (
	"../models"
	"../state"
	// "encoding/json"
	// "fmt"
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

func Pull(conn net.Conn, iData interface{}) {

	data := iData.(*models.PullRequest)

	queue := state.Pool[data.Channal]

	res := &models.Res{}

	if queue == nil {

		res.Msg = "No this queue!"
		res.Status = "404"

		models.Send(conn, res)
		return
	}

	var task *models.Task

	if data.Version == -1 {
		task = queue.List[0]
	} else {

		for _, t := range queue.List {
			if t.Version == data.Version {
				task = t
			}
		}
	}

	res.Msg = task.Msg
	res.Status = "200"

	models.Send(conn, res)
}
