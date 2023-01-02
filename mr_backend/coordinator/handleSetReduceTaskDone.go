package coordinator

import "github.com/brisk-dusk6157/mapReduceExercise/mr_backend/schemas"

func (c *Coordinator) SetReduceTaskDone(args *schemas.SetReduceTaskDoneArgs, reply *schemas.SetReduceTaskDoneReply) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	rTask := c.rTasks[args.TaskId]
	if rTask.state == stateInProgress && args.Shot == rTask.shot {
		rTask.state = stateDone
		rTask.output = args.File
	}
	return nil
}
