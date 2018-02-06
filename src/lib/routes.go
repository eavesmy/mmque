package lib

import (
	"../handler"
)

func init() {
	router := &Router{}

	router.Register(1, handler.Push)
	router.Register(2, handler.QueryOne)
	router.Register(3, handler.Ack)
	router.Register(4, handler.NewVersion)
	router.Register(5, handler.Pull)
}
