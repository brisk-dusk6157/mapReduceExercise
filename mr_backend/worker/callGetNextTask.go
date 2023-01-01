package worker

import (
	"github.com/brisk-dusk6157/mapReduceExercise/mr_backend/schemas"
	"log"
)

func (w *Worker) callGetNextTask() schemas.GetNextTaskReply {
	args := schemas.GetNextTaskArgs{}
	reply := schemas.GetNextTaskReply{}
	err := w.call("Coordinator.GetNextTask", &args, &reply)
	if err != nil {
		log.Fatal("GetNextTask call failed: ", err)
	}
	return reply
}
