package coordinator

import "github.com/brisk-dusk6157/mapReduceExercise/mr_backend/schemas"

func (c *Coordinator) GetReduceInputs(args *schemas.GetReduceInputsArgs, reply *schemas.GetReduceInputsReply) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, mTask := range c.mTasks {
		if mTask.state != STATE_DONE {
			reply.Ready = false
			return nil
		}
	}

	reply.Ready = true
	for _, mTask := range c.mTasks {
		reply.IntermediaryFiles = append(reply.IntermediaryFiles, mTask.outputs[args.Part])
	}
	return nil
}