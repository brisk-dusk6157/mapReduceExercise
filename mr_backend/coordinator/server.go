package coordinator

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func (c *Coordinator) Serve() {
	rpc.Register(c)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", "127.0.0.1:2534")
	if e != nil {
		log.Fatal("Listen error", e)
	}
	go http.Serve(l, nil)
	fmt.Println("Serving at 127.0.0.1:2534")

	for !c.Done() {
		c.debugPrintState()
		time.Sleep(5 * time.Second)
	}
}
