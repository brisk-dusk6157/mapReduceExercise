package coordinator

import (
	"fmt"
	"sync"
)

const STATE_IDLE = 0
const STATE_IN_PROGRESS = 1
const STATE_DONE = 2

type MapTask struct {
	id      int
	file    string
	state   int
	outputs map[int]string
}

type ReduceTask struct {
	id     int
	part   int
	state  int
	output string
}

type Coordinator struct {
	implPath string
	nParts   int

	mu sync.RWMutex

	mTasks map[int]*MapTask
	rTasks map[int]*ReduceTask
}

func New(implPath string, files []string, nParts int) Coordinator {
	c := Coordinator{
		implPath: implPath,
		nParts:   nParts,
		mTasks:   make(map[int]*MapTask),
		rTasks:   make(map[int]*ReduceTask),
	}
	for i, file := range files {
		c.mTasks[i] = &MapTask{
			id:    i,
			file:  file,
			state: STATE_IDLE,
		}
	}
	for i := 0; i < nParts; i++ {
		c.rTasks[i] = &ReduceTask{
			id:    i,
			part:  i,
			state: STATE_IDLE,
		}
	}
	return c
}

func (c *Coordinator) Done() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.debugPrintState()
	for _, mTask := range c.mTasks {
		if mTask.state != STATE_DONE {
			return false
		}
	}
	for _, rTask := range c.rTasks {
		if rTask.state != STATE_DONE {
			return false
		}
	}
	return true
}

func (c *Coordinator) debugPrintState() {
	c.mu.RLock()
	defer c.mu.RUnlock()
	fmt.Println("mTasks:")
	for _, mTask := range c.mTasks {
		fmt.Printf("- id=%d, state=%d\n", mTask.id, mTask.state)
	}
	fmt.Println("rTasks:")
	for _, rTask := range c.rTasks {
		fmt.Printf("- id=%d, state=%d\n", rTask.id, rTask.state)
	}
	fmt.Print("===========================\n\n")
}
