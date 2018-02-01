package lib

import (
	"../handler"
)

func init() {
	router := &Router{}

	router.Register(1, handler.Push)
	router.Register(2, handler.Pull)
	router.Register(3, handler.Ack)
}
