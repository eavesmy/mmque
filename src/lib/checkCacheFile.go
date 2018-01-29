package lib

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var SaveDataPipe = make(chan *Package, 1)

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
		fmt.Println(len(queue.List))
		if len(queue.List) == 0 {
			queue.List = append(queue.List, pack)
			continue
		}

		var index int

		for i, _pack := range queue.List {

			if pack.Version > _pack.Version {
				index = i
			}

			if pack.Version < _pack.Version {
				index = i - 1

				if index < 0 {
					index = 0
				}
			}
		}
	}

	fmt.Println(Pool["test"], 222)
	// 读取msg目录内所有msg
}

func Save2Local() {

	pack := <-SaveDataPipe
	if pack == nil {
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

	Save2Local()
}
