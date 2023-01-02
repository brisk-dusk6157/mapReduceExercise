package coordinator

import (
	"github.com/brisk-dusk6157/mapReduceExercise/mr_backend/schemas"
)

func (c *Coordinator) SetMapTaskDone(args *schemas.SetMapTaskDoneArgs, reply *schemas.SetMapTaskDoneReply) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	mTask := c.mTasks[args.TaskId]
	if mTask.state == STATE_IDLE || args.Shot != mTask.shot {
		return nil
	}

	mTask.state = STATE_DONE
	mTask.outputs = args.IntermediaryFiles
	return nil
}
