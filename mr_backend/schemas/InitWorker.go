package schemas

type InitWorkerArgs struct {
}

type InitWorkerReply struct {
	ImplPath string
	NParts   int
}
