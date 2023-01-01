package schemas

type SetMapTaskDoneArgs struct {
	TaskId            int
	IntermediaryFiles map[int]string
}

type SetMapTaskDoneReply struct {
}
