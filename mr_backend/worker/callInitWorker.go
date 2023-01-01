package worker

import (
	"github.com/brisk-dusk6157/mapReduceExercise/mr_backend/schemas"
	"log"
)

func (w *Worker) callInitWorker() schemas.InitWorkerReply {
	args := schemas.InitWorkerArgs{}
	reply := schemas.InitWorkerReply{}
	err := w.call("Coordinator.InitWorker", &args, &reply)
	if err != nil {
		log.Fatal("InitWorker call failed: ", err)
	}
	return reply
}
