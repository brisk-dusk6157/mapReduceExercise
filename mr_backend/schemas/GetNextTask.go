package schemas

const TASK_MAP = "map"
const TASK_REDUCE = "reduce"
const TASK_WAIT = "wait"
const TASK_STOP = "stop"

type GetNextTaskArgs struct {
}

type GetNextTaskReply struct {
	Task       string
	TaskId     int
	MapFile    string
	ReducePart int
}
