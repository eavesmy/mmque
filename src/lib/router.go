package lib

import (
	"net"
)

type Router struct {
}

var Routes = make(map[int]func(net.Conn, interface{}))

func (r *Router) Register(id int, handler func(net.Conn, interface{})) {
	Routes[id] = handler
}
