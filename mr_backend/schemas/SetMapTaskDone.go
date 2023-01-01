package schemas

type SetMapTaskDoneArgs struct {
	TaskId int
	Files  map[int]string
}

type SetMapTaskDoneReply struct {
}
