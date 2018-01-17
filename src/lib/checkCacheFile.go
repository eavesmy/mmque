package lib

import (
	"fmt"
	"os"
)

func init() {
	CheckLocalDir()
	CheckLocalMsg()
}

func CheckLocalDir() {

	if _, err := os.Stat("./msg"); os.IsNotExist(err) {
		os.Mkdir("msg", 0777)
	}
}

func CheckLocalMsg() {

	fs, _ := os.Open("./msg")
	files, _ := fs.Readdir(0)

	// 读取msg目录内所有msg
}
