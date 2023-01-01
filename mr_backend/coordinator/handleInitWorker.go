package coordinator

import "github.com/brisk-dusk6157/mapReduceExercise/mr_backend/schemas"

func (c *Coordinator) InitWorker(args *schemas.InitWorkerArgs, reply *schemas.InitWorkerReply) error {
	reply.ImplPath = c.implPath
	reply.NParts = c.nParts
	return nil
}
