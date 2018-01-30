package lib

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var Count int

func init() {
	CheckLocalDir()
	CheckLocalMsg()
	go Save2Local()
}

func CheckLocalDir() {

	if _, err := os.Stat("./msg"); os.IsNotExist(err) {
		os.Mkdir("msg", 0777)
	}
}

// 启动获取本地缓存
func CheckLocalMsg() {

	fs, _ := os.Open("./msg")
	files, _ := fs.Readdir(0)

	for _, v := range files {

		info := strings.Split(v.Name(), "_")

		version, _ := strconv.Atoi(info[1])

		pack := &Package{
			Channal: info[0],
			Version: version,
			Msg:     info[2],
		}

		queue := CreateQueue(pack.Channal)

		if len(queue.List) == 0 {
			queue.List = append(queue.List, pack)
			continue
		}

		index := 1
		for i, _pack := range queue.List {

			if pack.Version > _pack.Version {
				index = i + 1
			}

			if pack.Version < _pack.Version {
				index = i - 1
			}
		}

		// index++

		if index == 0 {

			temp_queue := []*Package{
				pack,
			}

			queue.List = append(temp_queue, queue.List...)
		} else {

			temp_queue := []*Package{}
			temp_queue = append(temp_queue, queue.List[:index]...)
			temp_queue = append(temp_queue, pack)
			queue.List = append(temp_queue, queue.List[index:]...)

			// temp_queue := append(queue.List[:index],pack)
		}
	}

	kCount := 0
	taskCount := 0
	for _, v := range Pool {
		kCount++

		for range v.List {
			taskCount++
		}
	}

	fmt.Println("本地缓存读取完毕,共", kCount, "条队列,", taskCount, "条任务")
}

func Save2Local() {

	//TODO : 连续刷盘 or 定时刷盘?
	// 迭代的话需要有被迭代的对象，需要做个大表，不合适
	// 进一个存一个也不太好，这个就不是连续刷盘了
	// 一个channal 一直取，一直放(指针)

	pack := <-SaveDataPipe

	if pack == nil {
		time.AfterFunc(5*time.Second, Save2Local)
		return
	}
	// 将内容转成buf,存到本地

	var buf bytes.Buffer

	buf.WriteString(pack.Channal)
	buf.WriteString("_")
	buf.WriteString(fmt.Sprintf("%d", pack.Version))
	buf.WriteString("_")
	buf.WriteString(pack.Msg)

	var buf_id []byte
	for _, v := range buf.Bytes() {
		if v != 0 {
			buf_id = append(buf_id, v)
		}
	}
	str_id := "./msg/" + string(buf_id)

	if _, err := os.Stat(str_id); os.IsNotExist(err) { // 存在放弃写入
		fs, _ := os.Create(str_id)
		defer fs.Close()
	}

	SaveDataPipe <- pack

	fmt.Println("刷盘", SaveDataPipe)
	Save2Local()
}
