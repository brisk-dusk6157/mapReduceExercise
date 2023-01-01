package worker

import (
	"github.com/brisk-dusk6157/mapReduceExercise/mr_backend/schemas"
	"log"
)

func (w *Worker) callSetMapTaskDone(taskId int, files map[int]string) schemas.SetMapTaskDoneReply {
	args := schemas.SetMapTaskDoneArgs{
		TaskId: taskId,
		Files:  files,
	}
	reply := schemas.SetMapTaskDoneReply{}
	err := w.call("Coordinator.SetMapTaskDone", &args, &reply)
	if err != nil {
		log.Fatal("SetMapTaskDone call failed: ", err)
	}
	return reply
}
