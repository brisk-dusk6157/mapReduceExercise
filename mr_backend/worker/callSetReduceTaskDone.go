package worker

import (
	"github.com/brisk-dusk6157/mapReduceExercise/mr_backend/schemas"
	"log"
)

func (w *Worker) callSetReduceTaskDone(taskId int, file string) schemas.SetReduceTaskDoneReply {
	args := schemas.SetReduceTaskDoneArgs{
		TaskId: taskId,
		File:   file,
	}
	reply := schemas.SetReduceTaskDoneReply{}
	err := w.call("Coordinator.SetReduceTaskDone", &args, &reply)
	if err != nil {
		log.Fatal("SetReduceTaskDone call failed: ", err)
	}
	return reply
}
