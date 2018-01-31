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

func (q *Queue) Pull() *models.Task {
	return q.List[0]
}

func (q *Queue) Check() []*models.Task {
	return q.List
}

func DeleteQueue() {

}

func FreshData() {
}
