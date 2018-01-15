package lib

import (
	"os"
)

func CheckLocalDir() {
	if _, err := os.Stat("./msg"); os.IsNotExist(err) {
		os.Mkdir("msg", 0777)
	}
}

func CheckLocalMsg() {
	// 读取msg目录内所有msg
}
