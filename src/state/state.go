package state

import (
	"../models"
)

const MAX_QUEUE_LEN = 100
const MAX_QUEUE_COUNT = 50

type Queue struct {
	List []*models.Task
}

var Pool = make(map[string]*Queue)
var SaveDataPipe = make(chan *models.Task, 1000)

func CreateQueue(name string) *Queue {

	if Pool[name] != nil {
		return Pool[name]
	}

	queue := &Queue{}
	Pool[name] = queue

	return queue
}

func (q *Queue) Push(data *models.Task) string {

	// 先查重,从末端查，判断版本号大小
	for _, pack := range q.List {
		if pack.Version == data.Version {
			return "0"
		}
	}

	q.List = append(q.List, data)
	SaveDataPipe <- data

	return "1"
}

func (q *Queue) Count() int {
	return len(q.List)
}

func (q *Queue) Pull(version int) *models.Task {

	if version == -1 {
		return q.List[0]
	}

	for _, task := range q.List {
		if task.Version == version {
			return task
		}
	}

	return nil
}

func (q *Queue) Ack(version int) bool {

	for i, task := range q.List {

		if task.Version == version {
			q.List = append(q.List[:i], q.List[i+1:]...)

			return DelFromLocal(task)
		}

	}

	return false
}

func (q *Queue) NewVersion() int {

	index := len(q.List) - 1
	version := 0

	if index < 0 {
		version = 0
	} else {
		version = q.List[index].Version + 1
	}

	return version
}
