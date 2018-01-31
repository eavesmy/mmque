package lib

import (
	"../handler"
)

func init() {
	router := &Router{}

	router.Register(1, handler.Push)
}
