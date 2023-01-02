package schemas

type SetMapTaskDoneArgs struct {
	TaskId            int
	Shot              int
	IntermediaryFiles map[int]string
}

type SetMapTaskDoneReply struct {
}
