package coordinator

import "github.com/brisk-dusk6157/mapReduceExercise/mr_backend/schemas"

func (c *Coordinator) SetReduceTaskDone(args *schemas.SetReduceTaskDoneArgs, reply *schemas.SetReduceTaskDoneReply) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	rTask := c.rTasks[args.TaskId]
	if rTask.state == STATE_IN_PROGRESS && args.Shot == rTask.shot {
		rTask.state = STATE_DONE
		rTask.output = args.File
	}
	return nil
}
