package worker

import (
	"log"
	"net/rpc"
)

func initClient(addr string) *rpc.Client {
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Fatal("initClient failed: ", err)
	}
	return client
}

func (w *Worker) call(rpcName string, args interface{}, reply interface{}) error {
	err := w.client.Call(rpcName, args, reply)
	if err != nil {
		return err
	}
	return nil
}
