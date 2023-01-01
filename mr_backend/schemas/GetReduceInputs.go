package schemas

type GetReduceInputsArgs struct {
	Part int
}

type GetReduceInputsReply struct {
	Ready             bool
	IntermediaryFiles []string
}
