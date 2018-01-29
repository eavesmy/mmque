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

	for _i, v := range files {
		info := strings.Split(v.Name(), "_")

		version, _ := strconv.Atoi(info[1])

		pack := &Package{
			Channal: info[0],
			Version: version,
			Msg:     info[2],
		}

		queue := CreateQueue(pack.Channal)

		_len := len(queue.List)

		fmt.Println("gogogo", _len, _i, queue.List, pack)

		if _len == 0 {
			queue.List = append(queue.List, pack)
			continue
		}

		fmt.Println(queue.List, 345, _len)
		for i, _pack := range queue.List {

			fmt.Println(_pack.Version, pack.Version)

			if pack.Version > _pack.Version {

				queue.List = append(queue.List, pack)

			} else {

				temp_list := []*Package{}

				temp_queue := append(temp_list, queue.List[i:]...)
				queue.List = append(queue.List[:i], pack)
				queue.List = append(queue.List, temp_queue...)
			}
		}
	}

	fmt.Println(Pool, 222)
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
