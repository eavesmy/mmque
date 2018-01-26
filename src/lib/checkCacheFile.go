package lib

import (
	"bytes"
	"fmt"
	"os"
	// "strings"
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

	fmt.Println(files)
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

	id := "./msg/" + buf.String()
	// TODO : 字符串拼接无效

	// fi, err := os.Stat("./msg/" + "aaa")

	// if err != nil {
	// fmt.Println(err.Error())
	// }

	// fmt.Println(fi, err, id)

	if _, err := os.Stat("./msg/" + "aaa"); os.IsNotExist(err) { // 存在放弃写入
		fmt.Println("Write done")
		os.Create(id)
		// defer fs.Close()
	}

	Save2Local()
}
