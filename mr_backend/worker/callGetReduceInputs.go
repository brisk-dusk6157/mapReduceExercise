package worker

import (
	"github.com/brisk-dusk6157/mapReduceExercise/mr_backend/schemas"
	"log"
)

func (w *Worker) callGetReduceInputs(part int) schemas.GetReduceInputsReply {
	args := schemas.GetReduceInputsArgs{
		Part: part,
	}
	reply := schemas.GetReduceInputsReply{}
	err := w.call("Coordinator.GetReduceInputs", &args, &reply)
	if err != nil {
		log.Fatal("GetReduceInputs call failed: ", err)
	}
	return reply
}
