package main

import (
	"github.com/brisk-dusk6157/mapReduceExercise/mr_backend/coordinator"
	"github.com/brisk-dusk6157/mapReduceExercise/mr_backend/worker"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Usage: `<prog> coordinator pathToPlugin.so N f1 f2...` or `<prog> worker <coordinatorAddr>`")
	}
	role := os.Args[1]
	switch role {
	case "coordinator":
		implPath := os.Args[2]
		nParts, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatal("Expected integer: ", err)
		}
		files := os.Args[4:]
		c := coordinator.New(implPath, files, nParts)
		c.Serve()
	case "worker":
		coordinatorAddr := os.Args[2]
		worker.Run(coordinatorAddr)
	default:
		log.Fatal("Unknown args %q", os.Args[1:])
	}
}
