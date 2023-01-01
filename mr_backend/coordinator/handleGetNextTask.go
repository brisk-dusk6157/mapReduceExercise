package coordinator

import "github.com/brisk-dusk6157/mapReduceExercise/mr_backend/schemas"

func (c *Coordinator) GetNextTask(args *schemas.GetNextTaskArgs, reply *schemas.GetNextTaskReply) error {
	// TODO: handle tasks timeouts
	//   - where to save startTime
	//   - how to check for timeout
	//   - what to do with Done reports from old tasks
	//   - limiting retries
	c.mu.Lock()
	for mTaskId, mTask := range c.mTasks {
		if mTask.state == STATE_IDLE {
			mTask.state = STATE_IN_PROGRESS
			reply.Task = schemas.TASK_MAP
			reply.TaskId = mTaskId
			reply.MapFile = mTask.file
			c.mu.Unlock()
			return nil
		}
	}
	for rTaskId, rTask := range c.rTasks {
		if rTask.state == STATE_IDLE {
			rTask.state = STATE_IN_PROGRESS
			reply.Task = schemas.TASK_REDUCE
			reply.TaskId = rTaskId
			reply.ReducePart = rTask.part
			c.mu.Unlock()
			return nil
		}
	}
	c.mu.Unlock()

	if c.Done() {
		reply.Task = schemas.TASK_STOP
		return nil
	}
	reply.Task = schemas.TASK_WAIT
	return nil
}
