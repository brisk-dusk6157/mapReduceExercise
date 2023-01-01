package worker

import (
	"github.com/brisk-dusk6157/mapReduceExercise/mr_backend/schemas"
	"net/rpc"
	"time"
)

type Worker struct {
	coordinatorAddr string
	nParts          int
	impl            Impl
	client          *rpc.Client
}

func Run(coordinatorAddr string) {
	w := Worker{
		coordinatorAddr: coordinatorAddr,
		client:          initClient(coordinatorAddr),
	}
	defer w.client.Close()

	initReply := w.callInitWorker()
	w.impl = loadImpl(initReply.ImplPath)
	w.nParts = initReply.NParts

	running := true
	for running {
		task := w.callGetNextTask()
		switch task.Task {
		case schemas.TASK_MAP:
			intermediaryFiles := w.execMap(task.TaskId, task.MapFile)
			w.callSetMapTaskDone(task.TaskId, intermediaryFiles)
		case schemas.TASK_REDUCE:
			intermediaryFiles := w.waitReduceInputs(task.ReducePart)
			result := w.execReduce(task.ReducePart, intermediaryFiles)
			w.callSetReduceTaskDone(task.TaskId, result)
		case schemas.TASK_WAIT:
			time.Sleep(1 * time.Second)
		case schemas.TASK_STOP:
			running = false
		}
	}
}

func (w *Worker) waitReduceInputs(part int) []string {
	for {
		reply := w.callGetReduceInputs(part)
		if reply.Ready {
			return reply.IntermediaryFiles
		}
		time.Sleep(1 * time.Second)
	}
}
